# Day 95 — GoLang Interview Notes

## 1. Why Go for DevOps / Cloud Engineering?

Go is popular in DevOps and cloud tooling because it provides:

- Simple syntax and fast development
- Fast compilation and execution
- Excellent concurrency support
- Strong standard library, especially for networking and HTTP
- Easy distribution as a single compiled binary
- Cross-platform compilation
- Strong cloud-native ecosystem

---

## 2. Variables and `:=`

Go uses `var` for explicit declarations and `:=` for short variable declarations.

```go
var port int = 8080
port := 8080
```

`:=` declares and initializes a new variable.

```go
port = 9090
```

`=` assigns a new value to an existing variable.

---

## 3. Slices and Maps

### Slice

A slice is a dynamically sized collection.

```go
services := []string{"api", "worker"}
services = append(services, "frontend")
```

### Map

A map stores key-value pairs.

```go
ports := map[string]int{
    "api":    8080,
    "worker": 8081,
}

ports["api"] = 9090
```

---

## 4. Loops

Go uses `for` for all loop types.

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

Condition-style loop:

```go
i := 1

for i <= 5 {
    fmt.Println(i)
    i++
}
```

Range loop:

```go
for index, service := range services {
    fmt.Println(index, service)
}
```

Use `_` when a returned value isn't needed:

```go
for _, service := range services {
    fmt.Println(service)
}
```

---

## 5. Functions

Basic function:

```go
func sayHello(name string) {
    fmt.Println("Hello", name)
}
```

Function with a return value:

```go
func add(a int, b int) int {
    return a + b
}
```

Multiple return values:

```go
func calculate(a int, b int) (int, int) {
    return a + b, a - b
}
```

Go commonly uses multiple return values for returning a result and an error.

---

## 6. Error Handling

Go normally handles errors as explicit return values rather than exceptions.

```go
func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }

    return a / b, nil
}
```

Caller:

```go
result, err := divide(10, 2)

if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

Important concepts:

- `error` → error type
- `err` → variable containing the error
- `nil` → no error
- `err != nil` → an error occurred
- `return err` → propagate the error
- `return nil` → successful completion

### Error Propagation

A lower-level function can return an error to its caller.

```go
func startService() error {
    result, err := getPort()

    if err != nil {
        return err
    }

    fmt.Println(result)
    return nil
}
```

This allows the appropriate higher-level layer to decide how to handle the error.

---

## 7. Structs

A struct groups related data into a custom type.

```go
type Service struct {
    Name   string
    Port   int
    Status string
}
```

Create an instance:

```go
service := Service{
    Name:   "api",
    Port:   8080,
    Status: "running",
}
```

Access fields:

```go
fmt.Println(service.Status)
```

Modify fields:

```go
service.Status = "stopped"
```

---

## 8. Methods

A method is a function associated with a type.

```go
func (s Service) Print() {
    fmt.Println(s.Name)
}
```

Call it:

```go
service.Print()
```

The `(s Service)` part is called the receiver.

---

## 9. Pointers

`&` gets the address of a value.

```go
servicePtr := &service
```

A pointer receiver allows a method to modify the original struct.

```go
func (s *Service) Stop() {
    s.Status = "stopped"
}
```

Then:

```go
service.Stop()
```

modifies the original `service`.

---

## 10. Interfaces

An interface defines a behavior contract.

```go
type Runner interface {
    Run()
}
```

It means:

> Anything with a `Run()` method satisfies the `Runner` interface.

Go uses implicit interface implementation. There is no `implements` keyword.

```go
type Server struct{}

func (s Server) Run() {
    fmt.Println("Server running")
}
```

`Server` satisfies `Runner` because it provides the required `Run()` method.

### Mental Model

```text
Struct     → What data does it have?
Interface  → What can it do?
Method     → How does it do it?
```

Interfaces are useful for decoupling implementations and making components easier to substitute and test.

---

# Day 95 — Core Mental Model

```text
Variables / Types
       ↓
Collections
       ↓
Conditions / Loops
       ↓
Functions
       ↓
Structs
       ↓
Methods
       ↓
Pointers
       ↓
Interfaces
       ↓
Errors
```

The goal isn't to memorize every piece of syntax.

The goal is to understand these building blocks well enough to combine them when building real Go applications.