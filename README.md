# golearn

Repository to learn golang from software development point of view.

Here’s a comprehensive list of topics and items one should learn to fully grasp **Golang (Go)** — from beginner to advanced levels. The content is structured to go from no knowledge to mastery:

---

## ✅ **1. Prerequisites (Optional)**

If you're new to programming:

* Basics of programming concepts (variables, loops, functions, data structures)
* Familiarity with another language like Python, Java, or C can help

---

## 🟢 **2. Go Basics**

* Setting up Go: Installation and environment (`GOROOT`, `GOPATH`)
* `go run`, `go build`, `go mod init`, `go get`, `go install`
* Basic syntax and formatting (`gofmt`)

### Key Concepts:

* **Variables and constants**
* **Data types** (int, float, string, bool, rune, byte)
* **Collections**
  * Arrays, slices, maps
  * Slice manipulation and capacity
  * Built-in functions (`append`, `copy`, `delete`, etc.)
* **Operators**
* **Control structures** (if, switch, for, range)
* **Functions** and **multiple return values**
* **Error handling** (`error` type)

---

## 🟡 **3. Intermediate Concepts**

### Structs and Interfaces

* Structs
* Methods
* Interfaces and interface composition
* Type assertions and type switches

### Advanced Functions

* First-class functions
* Closures
* Defer, Panic, and Recover
* Recursion

---

## 🔵 **4. Concurrency in Go (Core Strength)**

* Goroutines
* Channels (buffered and unbuffered)
* `select` statement
* Mutexes and `sync` package
* WaitGroups
* Worker pools
* Contexts (`context` package) for cancellation and timeouts

---

## 🟣 **5. Packages and Modules**

* Creating and importing packages
* `go mod` system (dependency management)
* Internal packages
* Versioning and `replace` directive

---

## 🟤 **6. Testing and Benchmarking**

* Writing tests using `testing` package
* Table-driven tests
* Mocking techniques
* Test coverage
* Benchmarking (`go test -bench`)

---

## 🟠 **7. Standard Library Essentials**

* `fmt`, `os`, `io`, `bufio`, `log`
* `net/http` for HTTP servers/clients
* `encoding/json`, `encoding/xml`
* `time`, `math`, `strings`, `strconv`
* File handling (`os` and `io/ioutil`)
* Reflection (`reflect` package)

---

## 🔴 **8. Advanced Go**

* Generics (since Go 1.18)
* Build tags and conditional compilation
* CGo (calling C code from Go)
* Embedding files using `//go:embed`
* Writing CLI tools
* Advanced interface usage (empty interfaces, dynamic types)

---

## ⚙️ **9. Web Development with Go**

* Creating HTTP servers
* Routing (using standard lib, famous `gin framework`, `gorilla/mux`, `chi`)
* Middleware
* JSON APIs
* Templating (`html/template`)
* Working with WebSockets
* Authentication (JWT, sessions)

---

## 🧰 **10. Tools and Ecosystem**

* `gofmt`, `go vet`, `golint`, `staticcheck`
* Profiling: `pprof`
* Dependency tools: `go mod`, `go get`
* IDEs: GoLand, VS Code with Go plugin

---

## 🌐 **11. Real-world Projects**

* REST API server
* CLI tool
* Web scraper
* Concurrency-based task manager
* Chat app using goroutines and channels

---

## 📚 **12. Resources**

* [golang.org](https://golang.org) official docs
* [Go by Example](https://gobyexample.com)
* [Effective Go](https://golang.org/doc/effective_go.html)
* [Go Wiki](https://github.com/golang/go/wiki)
* Books:

  * *The Go Programming Language* by Donovan & Kernighan
  * *Go in Action*
  * *Concurrency in Go* by Katherine Cox-Buday

---

