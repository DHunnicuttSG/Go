# ✅ Module 7 — Authentication & Security in Go (JWT + bcrypt)

## 🎯 Learning Objectives

By the end of this module, students will be able to:

* Hash and verify passwords using `bcrypt`
* Register and authenticate users
* Generate and validate JWT tokens
* Protect API endpoints with middleware
* Understand basic API security best practices

---

## 7.1 Required Packages

Install these:

```bash
go get golang.org/x/crypto/bcrypt
go get github.com/golang-jwt/jwt/v5
```

Imports used:

```go
import (
	"time"
	"net/http"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)
```

---

## 7.2 User Table

```sql
CREATE TABLE users (
	id INT AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(50) UNIQUE,
	password VARCHAR(255)
);
```

---

## 7.3 User Struct

```go
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}
```

---

## 7.4 Hashing Passwords with bcrypt

```go
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
```

### Compare password during login

```go
func checkPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
```

---

## 7.5 Register User Endpoint

```go
func registerUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	hashed, err := hashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", 500)
		return
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)",
		user.Username, hashed)

	if err != nil {
		http.Error(w, "Username already exists", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("✅ User registered"))
}
```

Test in Postman:

```json
{
  "username": "admin",
  "password": "secret123"
}
```

---

## 7.6 Creating JWT Tokens

```go
var jwtKey = []byte("super-secret-key")

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	return token.SignedString(jwtKey)
}
```

---

## 7.7 Login Endpoint

```go
func loginUser(w http.ResponseWriter, r *http.Request) {
	var creds User
	json.NewDecoder(r.Body).Decode(&creds)

	var storedHash string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", creds.Username).
		Scan(&storedHash)

	if err != nil || !checkPassword(storedHash, creds.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := generateToken(creds.Username)
	if err != nil {
		http.Error(w, "Token error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
```

Test:

```bash
POST /login
```

Response:

```json
{ "token": "eyJhbGciOiJIUzI1NiIsInR5..." }
```

---

## 7.8 JWT Middleware (Protect Routes)

```go
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
```

Now protect endpoints like this:

```go
http.HandleFunc("/contacts", authMiddleware(contactsHandler))
http.HandleFunc("/contacts/", authMiddleware(contactByIDHandler))
```

Now Postman **must include a header**:

```
Authorization: <paste token here>
```

✅ No token = denied
✅ Valid token = access

---

## 7.9 Secure Best Practices

| Rule                 | Why                     |
| -------------------- | ----------------------- |
| Hash passwords       | Never store plain text  |
| Use env variables    | No hard-coded secrets   |
| Use HTTPS            | Protect data in transit |
| Use token expiration | Limits hijacking        |
| Limit login attempts | Prevent brute force     |
| Validate input       | Prevent injection       |

Example using env variable:

```go
jwtKey := []byte(os.Getenv("JWT_SECRET"))
```

---

## ✅ Mini Project (Hands-On Lab)

**Secure your Contact API**:

1. Add `/register` and `/login`
2. Protect:

   * `/contacts`
   * `/contacts/{id}`
3. Only authenticated users can:

   * Add/edit/delete contacts

**Bonus:**
✅ Only admins can DELETE
✅ Add refresh tokens
✅ Add logout blacklist

---
