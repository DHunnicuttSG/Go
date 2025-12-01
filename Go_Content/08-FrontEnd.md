# ✅ Module 8 — Frontend Integration with Go (HTMX + HTML)

## 🎯 Learning Objectives

By the end of this module, students will:

* Connect a frontend to a Go REST API
* Use HTMX to make dynamic updates (no JS frameworks)
* Create a login form using JWT
* Build a contact dashboard
* Perform CRUD from the browser
* Handle authentication in the frontend

---

## 8.1 What is HTMX?

HTMX allows you to:

✅ Use normal HTML
✅ Avoid React / Vue
✅ Avoid writing tons of JavaScript
✅ Still get a SPA-like feel

Just add:

```html
<script src="https://unpkg.com/htmx.org@1.9.10"></script>
```

---

## 8.2 Folder Layout

```
contacts-app/
│
├── main.go
├── go.mod
├── templates/
│   ├── login.html
│   └── dashboard.html
└── static/
    └── styles.css
```

---

## 8.3 Serve HTML from Go

Add this to `main.go`:

```go
http.Handle("/static/",
	http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static"))))

http.HandleFunc("/", serveLogin)
http.HandleFunc("/dashboard", serveDashboard)
```

```go
func serveLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/login.html")
}

func serveDashboard(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/dashboard.html")
}
```

---

## 8.4 Login Page (login.html)

```html
<!DOCTYPE html>
<html>
<head>
  <title>Login</title>
  <script src="https://unpkg.com/htmx.org@1.9.10"></script>
  <link rel="stylesheet" href="/static/styles.css">
</head>
<body>

<h1>🔐 Login</h1>

<form 
  hx-post="/login" 
  hx-trigger="submit" 
  hx-swap="none"
  hx-on="htmx:afterRequest: window.location = '/dashboard'">

    <input type="text" name="username" placeholder="Username" required><br><br>
    <input type="password" name="password" placeholder="Password" required><br><br>

    <button>Login</button>
</form>

</body>
</html>
```

✅ After login → moves to `/dashboard`

---

## 8.5 Dashboard (dashboard.html)

```html
<!DOCTYPE html>
<html>
<head>
  <title>Contacts</title>
  <script src="https://unpkg.com/htmx.org@1.9.10"></script>
  <link rel="stylesheet" href="/static/styles.css">
</head>
<body>

<h1>📇 Contact Dashboard</h1>

<button 
   hx-get="/contacts" 
   hx-target="#contacts"
   hx-trigger="load">
   Load Contacts
</button>

<div id="contacts"></div>

<hr>

<h2>Add Contact</h2>

<form hx-post="/contacts" hx-target="#contacts" hx-swap="innerHTML">
  <input name="first_name" placeholder="First name">
  <input name="last_name" placeholder="Last name">
  <input name="company" placeholder="Company">
  <input name="phone" placeholder="Phone">
  <input name="email" placeholder="Email">
  <button>Add</button>
</form>

</body>
</html>
```

---

## 8.6 Make your API return partial HTML (HTMX format)

Change `getContacts()` to return HTML instead of JSON:

```go
func getContacts(w http.ResponseWriter) {
	rows, _ := db.Query("SELECT id, first_name, last_name, company, phone, email FROM contacts")
	defer rows.Close()

	w.Write([]byte("<table border='1'><tr><th>Name</th><th>Phone</th><th>Email</th></tr>"))

	for rows.Next() {
		var c Contact
		rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.Company, &c.Phone, &c.Email)

		w.Write([]byte("<tr>"))
		w.Write([]byte("<td>" + c.FirstName + " " + c.LastName + "</td>"))
		w.Write([]byte("<td>" + c.Phone + "</td>"))
		w.Write([]byte("<td>" + c.Email + "</td>"))
		w.Write([]byte("</tr>"))
	}

	w.Write([]byte("</table>"))
}
```

✅ HTMX loads it right into `<div id="contacts"></div>`

---

## 8.7 AJAX Delete Button

Inside the row:

```go
"<td><button hx-delete='/contacts/" + strconv.Itoa(c.ID) + 
"' hx-target='#contacts' hx-trigger='click'>Delete</button></td>"
```

Very powerful. No JavaScript required.

---

## 8.8 Add Basic CSS (styles.css)

```css
body { font-family: Arial; margin: 40px; }
input { margin: 4px; padding: 5px; }
button { padding: 6px 10px; cursor: pointer; }
table { border-collapse: collapse; margin-top: 15px; }
td, th { padding: 6px 10px; border: 1px solid #ccc; }
```

---

## ✅ Final Mini Project

Students now have:

✅ Go backend API
✅ JWT authentication
✅ MySQL database
✅ HTMX frontend
✅ Full contact system

**Give them this challenge:**

* Add edit button
* Add search bar
* Add pagination
* Add logout
* Add role-based UI

---
