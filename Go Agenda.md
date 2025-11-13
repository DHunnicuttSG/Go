## 🧭 GO LEARNING ROADMAP (Self-Study Outline)

---

### 🗓 **Phase 1: Foundations (Week 1–2)**

**Goal:** Understand Go’s syntax, tooling, and workflow.

#### 📚 Topics

* What is Go and why use it (concurrency, performance, simplicity)
* Installing Go & setting up environment (`go env`, `GOPATH`, `GOROOT`)
* Writing and running a simple Go program
* Go workspace layout and modules (`go mod init`, `go mod tidy`)
* Variables, constants, data types
* Functions (parameters, return values, named returns)
* Control flow: `if`, `for`, `switch`, `defer`
* Basic I/O and `fmt` package

#### 🧠 Exercises

* Write a “Hello, Go” program
* Build a small calculator CLI tool
* Print multiplication tables using loops

#### 🔗 Resources

* [Tour of Go](https://go.dev/tour/)
* [Effective Go](https://go.dev/doc/effective_go)

---

### 🗓 **Phase 2: Working with Data (Week 2–3)**

**Goal:** Get comfortable with Go’s data structures.

#### 📚 Topics

* Arrays and slices
* Maps
* Structs and methods
* Pointers (reference types)
* JSON encoding/decoding (`encoding/json`)
* Working with files (`os` and `io` packages)

#### 🧠 Exercises

* Build a contact book using structs and slices
* Parse and write JSON files
* Implement a word counter that reads a text file

#### 🔗 Resources

* [Go by Example: Structs, Maps, Slices](https://gobyexample.com/)

---

### 🗓 **Phase 3: Functions, Interfaces & Error Handling (Week 3–4)**

**Goal:** Write reusable and idiomatic Go code.

#### 📚 Topics

* Error handling (`error` type, `errors.New`, `fmt.Errorf`)
* Custom error types
* Interfaces (duck typing, polymorphism)
* Type assertions and type switches
* Anonymous functions and closures
* Packages and modular design

#### 🧠 Exercises

* Build a logging package
* Implement a shape interface with multiple structs (Circle, Square)
* Create a CLI app that validates user input with custom error types

---

### 🗓 **Phase 4: Concurrency & Channels (Week 4–5)**

**Goal:** Learn Go’s key strength — concurrency.

#### 📚 Topics

* Goroutines
* Channels (unbuffered and buffered)
* `select` statement
* WaitGroups (`sync` package)
* Mutexes (`sync.Mutex`)
* Context for cancellation and timeouts

#### 🧠 Exercises

* Build a concurrent web scraper
* Implement a worker pool pattern
* Create a timer or progress bar using goroutines

#### 🔗 Resources

* [Concurrency in Go by Katherine Cox-Buday](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/)
* [Go blog: Share Memory by Communicating](https://go.dev/blog/codelab-share)

---

### 🗓 **Phase 5: Web Development & APIs (Week 5–6)**

**Goal:** Build real-world web services.

#### 📚 Topics

* HTTP server (`net/http`)
* Routing with `http.ServeMux` or external router (`gorilla/mux`, `chi`)
* RESTful API design
* Request parsing and JSON responses
* Middleware (logging, authentication)
* Environment variables and config files

#### 🧠 Exercises

* Build a simple CRUD REST API (e.g., notes or tasks)
* Connect API to a SQLite/PostgreSQL database using `database/sql`
* Add middleware for logging and basic auth

#### 🔗 Resources

* [Go Web Examples](https://gowebexamples.com/)
* [Chi Router](https://github.com/go-chi/chi)

---

### 🗓 **Phase 6: Testing, Tooling & Deployment (Week 7–8)**

**Goal:** Write production-ready code.

#### 📚 Topics

* Unit testing (`testing` package)
* Table-driven tests
* Benchmarks and code coverage
* Using `go vet`, `golint`, `gofmt`
* Building and deploying binaries (`go build`, `go install`)
* Basic Docker setup for Go apps

#### 🧠 Exercises

* Write unit tests for your API
* Benchmark your functions
* Containerize your app with Docker

#### 🔗 Resources

* [Testing in Go](https://go.dev/doc/tutorial/add-a-test)
* [Uber Go Style Guide](https://github.com/uber-go/guide)

---

### 🏁 **Final Project (Optional but Highly Recommended)**

Build something that combines everything:

* Example ideas:

  * Task manager API with concurrency (goroutines for background cleanup)
  * Chat server using WebSockets
  * URL shortener service
  * CLI tool to monitor system stats

---

### 🧩 Optional Advanced Topics (After mastering the above)

* Generics
* Reflection
* Go Modules in monorepos
* gRPC with Protocol Buffers
* Working with microservices
* Advanced concurrency patterns (fan-in/fan-out, pipelines)

---
