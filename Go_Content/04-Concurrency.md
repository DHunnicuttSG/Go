# 📙 Module 4 — Concurrency in Go (Goroutines & Channels)

> “Don’t communicate by sharing memory; share memory by communicating.”
> — Go Proverb

Go treats **concurrency as a first-class feature** of the language.

---

## 1️⃣ What is Concurrency?

**Concurrency**: handling multiple tasks at the same time
**Parallelism**: running on multiple CPU cores at the same time

Go supports **massive concurrency** easily.

In other languages:

* Threads are heavy
* Hard to manage

In Go:

* Goroutines are **lightweight threads**
* Thousands can run easily

---

## 2️⃣ Goroutines

A **goroutine** is started with the `go` keyword.

```go
func sayHello() {
    fmt.Println("Hello")
}

func main() {
    go sayHello()
    time.Sleep(time.Second)
}
```

This runs concurrently!

### ⚠️ Without sleep:

Program ends before goroutine runs.

You will soon fix this with **channels or WaitGroups**.

### Multiple goroutines

```go
for i := 1; i <= 5; i++ {
    go fmt.Println(i)
}
time.Sleep(time.Second)
```

---

## 3️⃣ Channels (Communication Between Goroutines)

A **channel** is a typed conduit for communication.

```go
ch := make(chan string)
```

Send:

```go
ch <- "Hello"
```

Receive:

```go
msg := <- ch
```

### Real example:

```go
func greet(ch chan string) {
    ch <- "Hello from goroutine"
}

func main() {
    ch := make(chan string)
    go greet(ch)

    message := <- ch
    fmt.Println(message)
}
```

This is **synchronous communication**.

---

## 4️⃣ Buffered vs Unbuffered Channels

### Unbuffered (default)

Blocks until receiver is ready.

```go
ch := make(chan int)
```

### Buffered

Allows N items before blocking.

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3  // still OK
```

---

## 5️⃣ WaitGroup (sync package)

Lets main wait for goroutines to finish.

```go
var wg sync.WaitGroup

wg.Add(2)

go func() {
    defer wg.Done()
    fmt.Println("Task 1")
}()

go func() {
    defer wg.Done()
    fmt.Println("Task 2")
}()

wg.Wait()
```

✅ Proper concurrency management

---

## 6️⃣ Select statement

Used with multiple channels.

```go
select {
case msg1 := <- ch1:
    fmt.Println(msg1)
case msg2 := <- ch2:
    fmt.Println(msg2)
default:
    fmt.Println("No message yet")
}
```

Think of `select` like `switch` for channels.

---

## 7️⃣ Mutex (Race Condition Protection)

Used for **shared data**.

```go
var mu sync.Mutex
var count = 0

func increment() {
    mu.Lock()
    count++
    mu.Unlock()
}
```

Without mutex → ❌ unpredictable behavior

---

## 8️⃣ Go Concurrency Pattern – Worker Pool

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        time.Sleep(time.Second)
        results <- job * 2
    }
}
```

Main:

```go
jobs := make(chan int, 5)
results := make(chan int, 5)

for w := 1; w <= 3; w++ {
    go worker(w, jobs, results)
}

for j := 1; j <= 5; j++ {
    jobs <- j
}
close(jobs)

for a := 1; a <= 5; a++ {
    fmt.Println(<-results)
}
```

This is powerful for APIs and background work.

---

## 9️⃣ Common Mistakes

| Error                     | Fix                     |
| ------------------------- | ----------------------- |
| Main exits too early      | Use wg.Wait()           |
| Deadlock                  | Don’t wait if no sender |
| Data race                 | Use mutex               |
| Sending to closed channel | Only sender closes      |
| Forgetting to close       | Close when done         |

Detect race conditions:

```bash
go run -race main.go
```

---

# 🧪 LAB — Concurrent Contact Processor

Goal: Process contacts in parallel.

```go
func process(c Contact, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Processing:", c.Name)
}
```

Main:

```go
var wg sync.WaitGroup

for _, c := range contacts {
    wg.Add(1)
    go process(c, &wg)
}

wg.Wait()
```

Extend:
✅ Add channel for logging
✅ Save to file
✅ Add timer

---
