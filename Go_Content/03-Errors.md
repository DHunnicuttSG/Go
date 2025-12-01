# 📘 Module 3 — Errors, Interfaces & Design Patterns in Go

## 1️⃣ Error Handling in Go

Go does NOT use exceptions like Java/Python.
It uses **explicit error returns**.

### ✅ The error type

The `error` interface looks like this:

```go
type error interface {
    Error() string
}
```

Functions return errors like this:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

Usage:

```go
result, err := divide(10, 0)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(result)
}
```

---

## 2️⃣ Creating Errors

### ✔️ errors.New()

```go
import "errors"

err := errors.New("file not found")
```

### ✔️ fmt.Errorf()

Adds formatting:

```go
err := fmt.Errorf("user %d not found", id)
```

---

## 3️⃣ Custom Error Types

Why? → Add more context + behavior

```go
type NotFoundError struct {
    ID int
}

func (e NotFoundError) Error() string {
    return fmt.Sprintf("Record not found: %d", e.ID)
}
```

Use:

```go
func findUser(id int) (string, error) {
    if id == 0 {
        return "", NotFoundError{ID: id}
    }
    return "David", nil
}
```

Check type:

```go
err := findUser(0)
if _, ok := err.(NotFoundError); ok {
    fmt.Println("Custom error type detected!")
}
```

---

## 4️⃣ Interfaces (Duck Typing & Polymorphism)

> “If it walks like a duck and quacks like a duck…”

Interfaces in Go define **behavior**:

```go
type Speaker interface {
    Speak() string
}
```

Implementations:

```go
type Human struct{}
func (h Human) Speak() string {
    return "Hello"
}

type Dog struct{}
func (d Dog) Speak() string {
    return "Woof"
}
```

Use polymorphism:

```go
func speak(s Speaker) {
    fmt.Println(s.Speak())
}

speak(Human{})
speak(Dog{})
```

### ✅ No "implements" keyword!

Go checks automatically.

---

## 5️⃣ Type Assertions & Type Switches

### Type assertion

```go
var x interface{} = "Hello"

value, ok := x.(string)
if ok {
    fmt.Println(value)
}
```

### Type switch

```go
switch v := x.(type) {
case int:
    fmt.Println("Integer", v)
case string:
    fmt.Println("String", v)
default:
    fmt.Println("Unknown")
}
```

Common use: handling generic data

---

## 6️⃣ Anonymous Functions

Functions without a name:

```go
func() {
    fmt.Println("Hello")
}()
```

Assigned:

```go
add := func(a, b int) int {
    return a + b
}

fmt.Println(add(3, 5))
```

---

## 7️⃣ Closures

Closures “capture” variables from outer scope.

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
```

These are **very powerful** for stateful logic.

---

## 8️⃣ Packages and Modular Design

Create custom package:

```
/contact
   contact.go
main.go
```

### contact/contact.go

```go
package contact

type Contact struct {
    Name string
    Email string
}
```

### main.go

```go
package main

import (
    "fmt"
    "myapp/contact"
)

func main() {
    c := contact.Contact{"David", "d@email.com"}
    fmt.Println(c)
}
```

### Exported names must be Capitalized

| Exported | Not Exported |
| -------- | ------------ |
| Contact  | contact      |
| Save()   | save()       |

---

# 🧪 LAB — Error & Interface Driven Contact System

Use interfaces + custom errors + closures

```go
type ContactService interface {
    Add(c Contact) error
    Get(id int) (Contact, error)
}
```

Implement it and return:

```go
return Contact{}, NotFoundError{ID: id}
```

Use closure for ID counter:

```go
func createIdGenerator() func() int {
    id := 0
    return func() int {
        id++
        return id
    }
}
```

---

# ❌ Common Mistakes

| Mistake                           | Fix                          |
| --------------------------------- | ---------------------------- |
| Ignoring errors                   | Always check `if err != nil` |
| Misusing interfaces               | Use for behavior, not data   |
| Overusing pointers                | Use when mutating or large   |
| Forgetting capital letters export | Must be capitalized          |

---
