# 📚 Module 1 – Introduction to Go (Golang)

## ✅ What is Go and why use it?

**Go (Golang)** is a statically-typed compiled programming language created by Google in 2007 and released in 2009. It was designed to be:

* **Simple** like Python
* **Fast** like C++
* **Concurrent** like Erlang
* **Safe** like Java

### 🔹 Why companies use Go

| Feature                                | Why it matters                                                   |
| -------------------------------------- | ---------------------------------------------------------------- |
| **Concurrency** (goroutines, channels) | Efficient multi-tasking (web servers, APIs, distributed systems) |
| **Performance**                        | Compiles to machine code ⇒ very fast                             |
| **Simplicity**                         | Small language, easy to learn                                    |
| **Built-in tooling**                   | Testing, formatting, modules, etc                                |
| **Cloud-native**                       | Kubernetes, Docker, Terraform are written in Go                  |

### Example use cases

* APIs & microservices
* Networking tools
* CLI applications
* Cloud platforms
* Real-time systems

---

# 💻 Module 2 – Installing Go & Environment Setup

Check installation:

```bash
go version
```

Important environment values:

```bash
go env
```

### Key terms

| Term     | Meaning                                                   |
| -------- | --------------------------------------------------------- |
| `GOROOT` | Where Go is installed                                     |
| `GOPATH` | Your workspace for projects (less important with modules) |
| `GOMOD`  | Path to your module                                       |
| `GOBIN`  | Where Go binaries install                                 |

VS Code Extension:

```
Go (by Google)
```

---

# 📝 Module 3 – Writing & Running Your First Go Program

Create: `main.go`

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

Run it:

```bash
go run main.go
```

Build it:

```bash
go build main.go
```
* Creates a main.exe file

Run it:
```bash
./main
```

### Important concepts

* Every executable must use: `package main`
* Program starts in: `func main()`
* Go requires **explicit imports**

---

# 📦 Module 4 – Go Workspace & Modules

Go uses **modules** instead of the old GOPATH system.

Create a project:

```bash
mkdir contact-app
cd contact-app
go mod init contact-app
```

You now have:

```
contact-app/
  go.mod
  main.go
```

Install missing packages:

```bash
go mod tidy
```

This automatically adds packages to `go.mod`.

---

# 🧮 Module 5 – Variables, Constants & Data Types

### Variables

```go
var name string = "David"
age := 30        // short form
salary := 45.5
active := true
```

### Constants

```go
const pi = 3.14
```

### Common data types

| Type           | Example |
| -------------- | ------- |
| string         | "Hello" |
| int            | 42      |
| float64        | 3.14    |
| bool           | true    |
| []int          | slices  |
| map[string]int | maps    |

---

# 🔧 Module 6 – Functions

### Basic function

```go
func add(a int, b int) int {
    return a + b
}
```

### Multiple return values

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

### Named return

```go
func getName() (name string) {
    name = "David"
    return
}
```

---

# 🔁 Module 7 – Control Flow

### If statement

```go
if age > 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

### For loops

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### While-style loop

```go
count := 0
for count < 5 {
    fmt.Println(count)
    count++
}
```

### Switch

```go
day := 3

switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
default:
    fmt.Println("Other day")
}
```

### Defer (runs last)

```go
func main() {
    defer fmt.Println("End")
    fmt.Println("Start")
}
```

Output:

```
Start
End
```

---

# 📤 Module 8 – Basic I/O & fmt Package

### Output

```go
fmt.Println("Hello")
fmt.Printf("Age: %d\n", 30)
```

### Input

```go
var name string
fmt.Print("Enter name: ")
fmt.Scanln(&name)
fmt.Println("Hello", name)
```

### Common format verbs

| Verb | Meaning |
| ---- | ------- |
| %s   | string  |
| %d   | integer |
| %f   | float   |
| %t   | boolean |

---

# 📝 Practice Exercises

1. Write a program that:

   * Asks for your name
   * Asks for your age
   * Prints: “Hello <name>, you’ll be <age+10> in 10 years”

2. Write a loop that prints numbers from 1 to 100

3. Write a function that:

   * Accepts 2 numbers
   * Returns the larger one

---
