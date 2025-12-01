# 📚 Module 2 – Data Structures & Core Features in Go

## 1️⃣ Arrays and Slices

### ✅ Arrays

Arrays are **fixed-size** lists of a single data type.

```go
var nums [5]int
nums[0] = 10

scores := [3]int{90, 85, 88}
```

Characteristics:

* Fixed length (cannot grow/shrink)
* Rarely used in practice
* Passed by value (copied!)

### ✅ Slices (MOST USED)

Slices are **dynamic, flexible**, and built on top of arrays.

```go
names := []string{"Alice", "Bob", "Carl"}
names = append(names, "David")
```

**Length vs Capacity**

```go
fmt.Println(len(names), cap(names))
```

Simple Diagram:

```
Array:  [ a ][ b ][ c ][ d ]
Slice:        [ b ][ c ]
```

Slice from array:

```go
arr := [5]int{10, 20, 30, 40, 50}
part := arr[1:4]    // [20 30 40]
```

### 🔥 Common Mistakes

| Mistake                             | Fix                                        |
| ----------------------------------- | ------------------------------------------ |
| Changing slice changes array        | This is NORMAL: slices share backing array |
| Forgetting append returns new slice | Always reassign → s = append(s, x)         |

### ✅ Practice

Write a program that:

* Stores 5 names in a slice
* Appends 2 more
* Prints total length

---

## 2️⃣ Maps (Key-Value Storage)

Maps are like dictionaries / objects / hash tables.

```go
studentGrades := make(map[string]int)

studentGrades["David"] = 95
studentGrades["Sarah"] = 89
```

Or:

```go
employee := map[string]string{
    "name": "John",
    "role": "Engineer",
}
```

Retrieve:

```go
grade := studentGrades["David"]
```

Check existence:

```go
value, exists := studentGrades["Mike"]
if !exists {
    fmt.Println("Not found")
}
```

Delete:

```go
delete(studentGrades, "David")
```

---

## 3️⃣ Structs and Methods

Structs define **custom data types** (like classes).

```go
type Contact struct {
    ID        int
    FirstName string
    LastName  string
    Phone     string
    Email     string
}
```

Create struct:

```go
c := Contact{
    ID: 1,
    FirstName: "David",
    LastName: "H",
    Phone: "555-1111",
    Email: "d@email.com",
}
```

### ✅ Methods on structs

```go
func (c Contact) fullName() string {
    return c.FirstName + " " + c.LastName
}
```

Call method:

```go
fmt.Println(c.fullName())
```

Want changes to persist? Use pointer receiver:

```go
func (c *Contact) updatePhone(phone string) {
    c.Phone = phone
}
```

---

## 4️⃣ Pointers (Reference Types)

A pointer stores the **memory address** of a variable.

```go
x := 10
ptr := &x     // memory address
```

Access value with:

```go
fmt.Println(*ptr)
```

Modify via pointer:

```go
*ptr = 20
fmt.Println(x)  // Now 20
```

### Why pointers matter:

✅ Avoid memory copies
✅ Modify original data
✅ Efficient for large structs

### Simple mental model:

```
x     = 10
ptr   = &x
*ptr  = 10
```

---

## 5️⃣ JSON Encoding/Decoding (encoding/json)

Used in **APIs and REST services**.

### Struct + JSON

```go
type Contact struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

### Encode (Go ⇾ JSON)

```go
contact := Contact{1, "David", "d@email.com"}

jsonData, _ := json.Marshal(contact)
fmt.Println(string(jsonData))
```

Output:

```json
{"id":1,"name":"David","email":"d@email.com"}
```

### Decode (JSON ⇾ Go)

```go
jsonString := `{"id":2,"name":"Sarah","email":"s@email.com"}`

var c Contact
json.Unmarshal([]byte(jsonString), &c)

fmt.Println(c.Name)
```

---

## 6️⃣ Working with Files (os & io)

### Writing a file

```go
data := []byte("Hello Go")

err := os.WriteFile("data.txt", data, 0644)
if err != nil {
    fmt.Println(err)
}
```

### Reading a file

```go
data, err := os.ReadFile("data.txt")
if err != nil {
    fmt.Println(err)
}

fmt.Println(string(data))
```

### Appending to file

```go
file, _ := os.OpenFile("log.txt",
    os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

defer file.Close()

file.WriteString("New Entry\n")
```

---

# 💻 Mini-LAB: Contact Manager (No database version)

Goal: Use everything you learned.

```go
type Contact struct {
    ID    int
    Name  string
    Phone string
}

var contacts []Contact

func addContact(c Contact) {
    contacts = append(contacts, c)
}

func findContact(id int) *Contact {
    for _, c := range contacts {
        if c.ID == id {
            return &c
        }
    }
    return nil
}
```

Call in main:

```go
addContact(Contact{1,"David","123-456"})
addContact(Contact{2,"Sarah","654-321"})

fmt.Println(contacts)
```

---
