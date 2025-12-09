package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

type Contact struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Company   *string   `json:"company,omitempty"`
	Email     string    `json:"email"`
	Phone     *string   `json:"phone,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ContactInput struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Company   *string `json:"company"`
	Email     string  `json:"email"`
	Phone     *string `json:"phone"`
}

type PartialContact struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Company   *string `json:"company"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
}

type errorResponse struct {
	Error string `json:"error"`
}

var (
	db         *sql.DB
	emailRegex = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
)

func main() {
	// Environment-driven configuration
	// Example: MYSQL_DSN="user:pass@tcp(127.0.0.1:3306)/contactsdb?parseTime=true&charset=utf8mb4"
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		// Safe default for local dev (adjust user/pass/db as needed)
		dsn = "root:RootRoot@tcp(127.0.0.1:3306)/contactsdb?parseTime=true&charset=utf8mb4"
	}

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	// Connection pool settings (tune as needed)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	// Ping to ensure connectivity
	if err := db.Ping(); err != nil {
		log.Fatalf("db ping: %v", err)
	}

	if err := migrate(); err != nil {
		log.Fatalf("db migrate: %v", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/contacts", func(r chi.Router) {
		r.Get("/", listContacts)
		r.Post("/", createContact)
		r.Get("/{id}", getContact)
		r.Put("/{id}", updateContact)
		r.Patch("/{id}", patchContact)
		r.Delete("/{id}", deleteContact)
	})

	addr := ":8080"
	log.Printf("Contacts API listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func migrate() error {
	ddl := `
CREATE TABLE IF NOT EXISTS contacts (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  first_name VARCHAR(100) NOT NULL,
  last_name  VARCHAR(100) NOT NULL,
  company    VARCHAR(255),
  email      VARCHAR(255) NOT NULL UNIQUE,
  phone      VARCHAR(50),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX idx_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
`
	_, err := db.Exec(ddl)
	return err
}

func listContacts(w http.ResponseWriter, r *http.Request) {
	page := parseIntDefault(r.URL.Query().Get("page"), 1)
	pageSize := parseIntDefault(r.URL.Query().Get("pageSize"), 50)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 200 {
		pageSize = 50
	}
	offset := (page - 1) * pageSize

	rows, err := db.Query(`
SELECT id, first_name, last_name, company, email, phone, created_at, updated_at
FROM contacts
ORDER BY id
LIMIT ? OFFSET ?`, pageSize, offset)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	var items []Contact
	for rows.Next() {
		var c Contact
		var company, phone sql.NullString
		var created, updated time.Time
		if err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &company, &c.Email, &phone, &created, &updated); err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}
		if company.Valid {
			c.Company = &company.String
		}
		if phone.Valid {
			c.Phone = &phone.String
		}
		c.CreatedAt = created
		c.UpdatedAt = updated
		items = append(items, c)
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"page":     page,
		"pageSize": pageSize,
		"items":    items,
	})
}

func getContact(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	var c Contact
	var company, phone sql.NullString
	var created, updated time.Time
	err = db.QueryRow(`
SELECT id, first_name, last_name, company, email, phone, created_at, updated_at
FROM contacts WHERE id = ?`, id).
		Scan(&c.ID, &c.FirstName, &c.LastName, &company, &c.Email, &phone, &created, &updated)
	if errors.Is(err, sql.ErrNoRows) {
		writeError(w, http.StatusNotFound, fmt.Errorf("contact %d not found", id))
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	if company.Valid {
		c.Company = &company.String
	}
	if phone.Valid {
		c.Phone = &phone.String
	}
	c.CreatedAt = created
	c.UpdatedAt = updated

	writeJSON(w, http.StatusOK, c)
}

func createContact(w http.ResponseWriter, r *http.Request) {
	var in ContactInput
	if err := decodeJSON(r, &in); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if err := validateInput(in); err != nil {
		writeError(w, http.StatusUnprocessableEntity, err)
		return
	}

	// created_at and updated_at are handled by MySQL defaults, but we can set explicitly if desired
	now := time.Now().UTC().Truncate(time.Second)

	res, err := db.Exec(`
INSERT INTO contacts (first_name, last_name, company, email, phone, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?, ?)`,
		in.FirstName, in.LastName, nullable(in.Company), in.Email, nullable(in.Phone), now, now)
	if err != nil {
		if isUniqueEmailErr(err) {
			writeError(w, http.StatusConflict, fmt.Errorf("email already exists"))
			return
		}
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	id, _ := res.LastInsertId()

	contact := Contact{
		ID:        id,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Company:   in.Company,
		Email:     in.Email,
		Phone:     in.Phone,
		CreatedAt: now,
		UpdatedAt: now,
	}
	writeJSON(w, http.StatusCreated, contact)
}

func updateContact(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	var in ContactInput
	if err := decodeJSON(r, &in); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if err := validateInput(in); err != nil {
		writeError(w, http.StatusUnprocessableEntity, err)
		return
	}
	now := time.Now().UTC().Truncate(time.Second)

	res, err := db.Exec(`
UPDATE contacts
SET first_name = ?, last_name = ?, company = ?, email = ?, phone = ?, updated_at = ?
WHERE id = ?`,
		in.FirstName, in.LastName, nullable(in.Company), in.Email, nullable(in.Phone), now, id)
	if err != nil {
		if isUniqueEmailErr(err) {
			writeError(w, http.StatusConflict, fmt.Errorf("email already exists"))
			return
		}
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		writeError(w, http.StatusNotFound, fmt.Errorf("contact %d not found", id))
		return
	}
	getContact(w, r) // return the updated resource
}

func patchContact(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	var in PartialContact
	if err := decodeJSON(r, &in); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	fields := make([]string, 0, 6)
	args := make([]any, 0, 7)
	if in.FirstName != nil {
		if strings.TrimSpace(*in.FirstName) == "" {
			writeError(w, http.StatusUnprocessableEntity, fmt.Errorf("firstName is required"))
			return
		}
		fields = append(fields, "first_name = ?")
		args = append(args, *in.FirstName)
	}
	if in.LastName != nil {
		if strings.TrimSpace(*in.LastName) == "" {
			writeError(w, http.StatusUnprocessableEntity, fmt.Errorf("lastName is required"))
			return
		}
		fields = append(fields, "last_name = ?")
		args = append(args, *in.LastName)
	}
	if in.Company != nil {
		fields = append(fields, "company = ?")
		args = append(args, nullable(in.Company))
	}
	if in.Email != nil {
		if !emailRegex.MatchString(*in.Email) {
			writeError(w, http.StatusUnprocessableEntity, fmt.Errorf("invalid email"))
			return
		}
		fields = append(fields, "email = ?")
		args = append(args, *in.Email)
	}
	if in.Phone != nil {
		fields = append(fields, "phone = ?")
		args = append(args, nullable(in.Phone))
	}
	fields = append(fields, "updated_at = ?")
	now := time.Now().UTC().Truncate(time.Second)
	args = append(args, now)
	args = append(args, id)

	if len(args) == 2 { // only updated_at + id
		writeError(w, http.StatusBadRequest, fmt.Errorf("no updatable fields provided"))
		return
	}

	q := fmt.Sprintf("UPDATE contacts SET %s WHERE id = ?", strings.Join(fields, ", "))
	res, err := db.Exec(q, args...)
	if err != nil {
		if isUniqueEmailErr(err) {
			writeError(w, http.StatusConflict, fmt.Errorf("email already exists"))
			return
		}
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		writeError(w, http.StatusNotFound, fmt.Errorf("contact %d not found", id))
		return
	}
	getContact(w, r)
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(chi.URLParam(r, "id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	res, err := db.Exec(`DELETE FROM contacts WHERE id = ?`, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		writeError(w, http.StatusNotFound, fmt.Errorf("contact %d not found", id))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// Helpers

func decodeJSON(r *http.Request, v any) error {
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(v)
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, errorResponse{Error: err.Error()})
}

func parseIDParam(s string) (int64, error) {
	id, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	if err != nil || id <= 0 {
		return 0, fmt.Errorf("invalid id: %q", s)
	}
	return id, nil
}

func parseIntDefault(s string, def int) int {
	if s == "" {
		return def
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return n
}

func nullable(p *string) any {
	if p == nil {
		return nil
	}
	return *p
}

func validateInput(in ContactInput) error {
	if strings.TrimSpace(in.FirstName) == "" {
		return fmt.Errorf("firstName is required")
	}
	if strings.TrimSpace(in.LastName) == "" {
		return fmt.Errorf("lastName is required")
	}
	if !emailRegex.MatchString(in.Email) {
		return fmt.Errorf("invalid email")
	}
	return nil
}

func isUniqueEmailErr(err error) bool {
	s := strings.ToLower(err.Error())
	// MySQL constraint messages can vary by server version/SQL mode.
	return strings.Contains(s, "duplicate") && strings.Contains(s, "email")
}
