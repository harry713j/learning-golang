# 🧠 Golang Full Concepts Roadmap

---

## 🧱 1. Basics

| Concept           | Description                             |
|-------------------|-----------------------------------------|
| Hello World       | Entry point with `main()` and `package main` |
| Variables         | `var`, `:=` syntax for declaration      |
| Constants         | `const` keyword                         |
| Data Types        | `int`, `float64`, `bool`, `string`, `byte`, `rune`, etc. |
| Type Inference    | Automatic detection using `:=`          |
| Comments          | `//` single-line, `/* */` multi-line    |
| Input / Output    | `fmt.Print`, `fmt.Println`, `fmt.Scanf`, `fmt.Sprintf` |

---

## 🔁 2. Control Structures

| Concept           | Description                             |
|-------------------|-----------------------------------------|
| If / Else         | Basic conditionals                      |
| Switch            | Multiple condition handling             |
| For Loops         | Only one loop type in Go (`for`)        |
| Break / Continue  | Flow control inside loops               |
| Labels            | For labeled break/continue              |

---

## 📦 3. Functions

| Concept                   | Description                                 |
|---------------------------|---------------------------------------------|
| Function Declaration      | `func name(params) returnType`              |
| Multiple Return Values    | `return val1, val2`                         |
| Named Return Values       | Return values declared in the signature     |
| Variadic Functions        | `func sum(nums ...int)`                     |
| First-class Functions     | Assign functions to variables               |
| Anonymous / Lambda Funcs  | `func(x int) int { return x * x }`          |

---

## 🧩 4. Data Structures

| Structure     | Description                      |
|---------------|----------------------------------|
| Arrays        | Fixed-size collections           |
| Slices        | Dynamic-sized views over arrays  |
| Maps          | Key-value pairs                  |
| Structs       | Custom data types (like objects) |
| Pointers      | Memory references: `*`, `&`      |

---

## 🛠️ 5. Advanced Struct Usage

| Feature           | Description                            |
|-------------------|----------------------------------------|
| Struct Embedding  | Like inheritance                      |
| Tags              | Used for JSON, DB mapping, etc.       |
| Methods           | Functions attached to structs          |
| Receivers         | `func (p Person)` vs `func (p *Person)`|

---

## 🧰 6. Error Handling

| Concept           | Description                         |
|-------------------|-------------------------------------|
| `error` interface | Standard way to handle errors       |
| `errors.New()`    | Create new errors                   |
| `fmt.Errorf()`    | Format errors                       |
| Custom Errors     | Implement the `Error()` method      |
| `panic` & `recover` | Exception-like flow (for critical) |

---

## 📦 7. Packages & Modules

| Concept          | Description                                 |
|------------------|---------------------------------------------|
| Packages         | Code organization (`package xyz`)           |
| Importing        | `import "fmt"`                              |
| `go.mod`         | Go module definition                        |
| `go.sum`         | Dependency checksums                        |
| `go get`         | Install external packages                   |
| Custom Packages  | Organize code into multiple files/folders   |

---

## 🔀 8. Interfaces & Polymorphism

| Concept               | Description                           |
|------------------------|---------------------------------------|
| Interfaces            | Type that specifies behavior          |
| Implicit Implementation | No `implements` keyword needed       |
| Type Assertion        | Access underlying value of interface  |
| Type Switch           | Switch based on dynamic type          |
| Empty Interface       | `interface{}` holds any type          |

---

## ⛓️ 9. Concurrency (Go's Superpower)

| Concept           | Description                                 |
|-------------------|---------------------------------------------|
| Goroutines        | Lightweight threads using `go` keyword      |
| Channels          | Communication between goroutines           |
| Buffered Channels | Channels with capacity                      |
| Select Statement  | Multiplex channel operations                |
| WaitGroups        | Wait for a group of goroutines              |
| Mutex / Sync      | Prevent race conditions                     |

---

## 🌐 10. Standard Library & Utilities

| Category       | Packages                                |
|----------------|------------------------------------------|
| I/O            | `fmt`, `os`, `bufio`, `io`              |
| File Handling  | `os`, `ioutil`, `filepath`              |
| Networking     | `net`, `net/http`                       |
| Strings & Time | `strings`, `strconv`, `time`            |
| JSON           | `encoding/json`                         |
| HTTP Clients   | `net/http`, `http.NewRequest()`         |
| Testing        | `testing`, `go test`, table-driven tests|
| Regex          | `regexp`                                |

---

## 🧪 11. Testing

| Concept         | Description                                |
|------------------|--------------------------------------------|
| Unit Tests       | Use `*_test.go` files                      |
| Table-Driven     | Iterate over test cases                    |
| Benchmarking     | `BenchmarkXxx(b *testing.B)`               |
| Mocking          | Use interfaces for mocking behavior        |

---

## 🧰 12. Tooling & Commands

| Tool/Command    | Purpose                           |
|------------------|-----------------------------------|
| `go build`       | Compile code                      |
| `go run`         | Run code directly                 |
| `go mod init`    | Initialize module                 |
| `go get`         | Add a dependency                  |
| `go fmt`         | Format code                       |
| `go vet`         | Analyze suspicious code           |
| `go test`        | Run unit tests                    |
| `go install`     | Build and install binary          |
| `go doc`         | View documentation                |
| `go generate`    | Run code generators               |

---

## 🧱 13. Project Structure (Best Practice)

myapp/
├── go.mod
├── main.go
├── handler/
│ └── user.go
├── models/
│ └── user.go
├── utils/
│ └── helper.go
├── internal/
│ └── somepkg/
├── cmd/
│ └── cli-tool/

## 🔐 14. Miscellaneous Concepts

| Concept             | Description                                 |
|----------------------|---------------------------------------------|
| Reflection           | Use `reflect` package                       |
| Embedding            | Inheritance-like with struct composition    |
| Defer                | Run a function after the current one ends   |
| Closures             | Functions capturing outer scope             |
| File/Env Handling    | `os.Open()`, `os.Getenv()`                  |
| JSON Marshal/Unmarshal | `json.Marshal()`, `json.Unmarshal()`     |

---

## 🏁 Ready to Build?

Once you master these, you're ready to build:

- ✅ CLI tools  
- ✅ Web APIs using `net/http`  
- ✅ Goroutine-based servers  
- ✅ Microservices  
- ✅ Dev tools / automation scripts  