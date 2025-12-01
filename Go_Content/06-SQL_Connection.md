# ✅ **Module 6 — Databases and REST APIs in Go**

## 🎯 Learning Objectives

By the end of this module, students will be able to:

* Connect to a MySQL database from Go
* Perform CRUD operations (Create, Read, Update, Delete)
* Build a REST API using `net/http`
* Structure a basic API project
* Send and receive JSON
* Test endpoints using Postman or curl

---

## 6.1 Go + MySQL Basics

### Required package:

```bash
go get github.com/go-sql-driver/mysql
```

### Import in code:

```go
import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
```

The underscore means:

> Import the package **only for its side effects** (registering the MySQL driver), not for direct use.

---

## 6.2 Database Connection Example

```go
var db *sql.DB

func initDB() {
	var err error
	dsn := "root:password@tcp(localhost:3306)/contacts_db"
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	println("✅ Database connected")
}
```

> ✅ `Ping()` confirms the DB is actually reachable.

---

## 6.3 Contact Struct

```go
type Contact struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Company   string `json:"company"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}
```

---

## 6.4 Simple REST API with net/http

### Basic server

```go
package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	initDB()

	http.HandleFunc("/contacts", contactsHandler)
	http.HandleFunc("/contacts/", contactByIDHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

## 6.5 GET: All Contacts

```go
func contactsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getContacts(w)
	case http.MethodPost:
		addContact(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getContacts(w http.ResponseWriter) {
	rows, err := db.Query("SELECT id, first_name, last_name, company, phone, email FROM contacts")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var contacts []Contact

	for rows.Next() {
		var c Contact
		rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Company, &c.Phone, &c.Email)
		contacts = append(contacts, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}
```

---

## 6.6 GET: One Contact by ID

```go
func contactByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/contacts/"):]

	if r.Method == http.MethodGet {
		var c Contact

		err := db.QueryRow("SELECT id, first_name, last_name, company, phone, email FROM contacts WHERE id = ?", id).
			Scan(&c.ID, &c.FirstName, &c.LastName, &c.Company, &c.Phone, &c.Email)

		if err != nil {
			http.Error(w, "Contact not found", 404)
			return
		}

		json.NewEncoder(w).Encode(c)
	}
}
```

Test:

```
GET http://localhost:8080/contacts/1
```

---

## 6.7 POST: Add a Contact

```go
func addContact(w http.ResponseWriter, r *http.Request) {
	var c Contact

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	result, err := db.Exec(`INSERT INTO contacts 
		(first_name, last_name, company, phone, email)
		VALUES (?, ?, ?, ?, ?)`,
		c.FirstName, c.LastName, c.Company, c.Phone, c.Email)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	id, _ := result.LastInsertId()
	c.ID = int(id)

	json.NewEncoder(w).Encode(c)
}
```

Test with Postman or curl:

```bash
curl -X POST http://localhost:8080/contacts \
-H "Content-Type: application/json" \
-d '{
  "first_name":"John",
  "last_name":"Doe",
  "company":"Acme",
  "phone":"123-456-7890",
  "email":"john@acme.com"
}'
```

---

## 6.8 PUT: Update Contact

```go
func updateContact(id string, c Contact) error {
	_, err := db.Exec(`
		UPDATE contacts 
		SET first_name=?, last_name=?, company=?, phone=?, email=?
		WHERE id=?`,
		c.FirstName, c.LastName, c.Company, c.Phone, c.Email, id)

	return err
}
```

---

## 6.9 DELETE Contact

```go
func deleteContact(id string) error {
	_, err := db.Exec("DELETE FROM contacts WHERE id=?", id)
	return err
}
```

You can extend your `contactByIDHandler` to handle `PUT` and `DELETE`.

---

# ✅ Suggested Folder Structure

```
contacts-api/
│
├── go.mod
├── main.go
└── handlers.go
```

✅ Very simple
✅ Minimal imports
✅ Perfect for teaching or demos

---

# ✅ Mini Project (Student Exercise)

**Build a Contact API that:**

* Stores contact info in MySQL
* Supports:

  * GET /contacts
  * GET /contacts/{id}
  * POST /contacts
  * PUT /contacts/{id}
  * DELETE /contacts/{id}

Extra credit:
✅ Add search by last name
✅ Add validation
✅ Add paging

---
