# 📕 Module 5 — Building REST APIs in Go

> Goal: Create real HTTP APIs using Go’s built-in tools (no heavy frameworks).

By the end of this module, students will:
✅ Understand REST concepts
✅ Build an API using `net/http`
✅ Return JSON data
✅ Handle GET, POST, PUT, DELETE
✅ Build a mini Contact API

---

## 1️⃣ What is a REST API?

**REST = Representational State Transfer**

A REST API:

* Uses HTTP methods (GET, POST, PUT, DELETE)
* Sends & receives **JSON**
* Works over URLs (endpoints)

Example endpoints:

| Method | URL           | Action                 |
| ------ | ------------- | ---------------------- |
| GET    | `/contacts`   | Get all contacts       |
| GET    | `/contacts/1` | Get a specific contact |
| POST   | `/contacts`   | Create a new contact   |
| PUT    | `/contacts/1` | Update a contact       |
| DELETE | `/contacts/1` | Delete a contact       |

---

## 2️⃣ Go’s Web Server (net/http)

Go has a **built-in web server**:

```go
import "net/http"

func main() {
    http.ListenAndServe(":8080", nil)
}
```

Create a route:

```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello API")
}

func main() {
    http.HandleFunc("/", helloHandler)
    http.ListenAndServe(":8080", nil)
}
```

Test in browser:

```
http://localhost:8080
```

---

## 3️⃣ Returning JSON (encoding/json)

Instead of HTML, return JSON:

```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    data := map[string]string{"message": "Hello, API"}
    json.NewEncoder(w).Encode(data)
}
```

---

## 4️⃣ Create Contact Struct

```go
type Contact struct {
    ID        int    `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Phone     string `json:"phone"`
    Email     string `json:"email"`
}
```

Temporary storage:

```go
var contacts = []Contact{
    {1, "David", "H", "123-456", "d@email.com"},
    {2, "Sarah", "J", "222-333", "s@email.com"},
}
```

---

## 5️⃣ GET – All Contacts

```go
func getContacts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(contacts)
}
```

Route it:

```go
http.HandleFunc("/contacts", getContacts)
```

Test:

```
GET http://localhost:8080/contacts
```

✅ Returns JSON list!

---

## 6️⃣ GET – Contact by ID

We’ll extract the ID from the URL:

```
/contacts/2
```

```go
func getContact(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/contacts/")
    id, _ := strconv.Atoi(idStr)

    for _, c := range contacts {
        if c.ID == id {
            json.NewEncoder(w).Encode(c)
            return
        }
    }

    http.NotFound(w, r)
}
```

Route:

```go
http.HandleFunc("/contacts/", getContact)
```

---

## 7️⃣ POST – Create Contact

```go
func createContact(w http.ResponseWriter, r *http.Request) {
    var newContact Contact
    json.NewDecoder(r.Body).Decode(&newContact)

    newContact.ID = len(contacts) + 1
    contacts = append(contacts, newContact)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newContact)
}
```

Test with Postman or curl:

```bash
curl -X POST http://localhost:8080/contacts \
-H "Content-Type: application/json" \
-d '{"first_name":"Mark","last_name":"Lee","phone":"444","email":"m@email.com"}'
```

---

## 8️⃣ PUT – Update Contact

```go
func updateContact(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/contacts/")
    id, _ := strconv.Atoi(idStr)

    for i, c := range contacts {
        if c.ID == id {
            json.NewDecoder(r.Body).Decode(&contacts[i])
            contacts[i].ID = id
            json.NewEncoder(w).Encode(contacts[i])
            return
        }
    }

    http.NotFound(w, r)
}
```

---

## 9️⃣ DELETE – Remove Contact

```go
func deleteContact(w http.ResponseWriter, r *http.Request) {
    idStr := strings.TrimPrefix(r.URL.Path, "/contacts/")
    id, _ := strconv.Atoi(idStr)

    for i, c := range contacts {
        if c.ID == id {
            contacts = append(contacts[:i], contacts[i+1:]...)
            w.WriteHeader(http.StatusNoContent)
            return
        }
    }

    http.NotFound(w, r)
}
```

---

## 🔀 Routing All Methods Properly

Use condition on `r.Method`:

```go
func contactsHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getContacts(w, r)
    case http.MethodPost:
        createContact(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
```

And for one contact:

```go
func contactHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getContact(w, r)
    case http.MethodPut:
        updateContact(w, r)
    case http.MethodDelete:
        deleteContact(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
```

Routes:

```go
http.HandleFunc("/contacts", contactsHandler)
http.HandleFunc("/contacts/", contactHandler)
```

---

# 🧪 LAB — Full Mini Project

**Contact REST API**

Students must:
✅ Create struct
✅ Serve on port 8080
✅ Implement:

* GET /contacts
* GET /contacts/{id}
* POST /contacts
* PUT /contacts/{id}
* DELETE /contacts/{id}

BONUS:

✅ Add file persistence with JSON
✅ Add goroutines for logging
✅ Add interfaces for repository

---

## 🧩 Architecture Visual (Text)

```
Client (Postman/Browser)
        ↓
   Go HTTP Server
        ↓
   Handler Functions
        ↓
    Contacts Slice / (DB next)
```

---
