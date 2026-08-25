# Day 96 — GoLang Interview Notes

## 1. File Handling

Go uses the `os` package for basic file operations.

### Reading a file

```go
data, err := os.ReadFile("message.txt")

if err != nil {
    return
}

fmt.Println(string(data))
```

`os.ReadFile()` returns:

```text
[]byte + error
```

### Writing a file

```go
data := []byte("Cloud Platform Kit")

err := os.WriteFile("output.txt", data, 0644)

if err != nil {
    return
}
```

`os.WriteFile()` takes:

```text
filename + []byte data + permissions
```

and returns an `error`.

---

## 2. JSON

Go uses the `encoding/json` package.

### Marshal — Go → JSON

```go
data, err := json.Marshal(service)

if err != nil {
    return
}

fmt.Println(string(data))
```

Mental model:

```text
Go struct → json.Marshal() → []byte JSON
```

### Unmarshal — JSON → Go

```go
var service Service

err := json.Unmarshal(data, &service)

if err != nil {
    return
}
```

Mental model:

```text
JSON → json.Unmarshal() → Go struct
```

The struct is the **destination** where the decoded JSON is stored.

### JSON Struct Tags

```go
type ScanRequest struct {
    Repository string `json:"repository"`
    Branch     string `json:"branch"`
}
```

Struct tags tell Go's JSON package which JSON field names to use.

---

## 3. Environment Variables

Environment variables allow configuration to be provided outside the application.

```go
appName := os.Getenv("APP_NAME")
```

If the variable doesn't exist, `os.Getenv()` returns an empty string.

```go
if appName == "" {
    fmt.Println("APP_NAME is not configured")
    return
}
```

This is useful for keeping configuration out of the source code and changing configuration between environments.

Example:

```text
Development → APP_NAME=scanner
Production  → APP_NAME=scanner-prod
```

---

## 4. Packages

A package groups related Go code.

```go
package config
```

The important distinction is:

```text
Same package       → can access unexported identifiers
Different package  → only exported identifiers are accessible
```

### Exported vs Unexported

```go
func GetPort() int {
    return 8080
}
```

`GetPort()` is exported because it starts with a capital letter.

```go
func getPort() int {
    return 8080
}
```

`getPort()` is unexported because it starts with a lowercase letter.

Mental model:

```text
Capital letter → Exported
lowercase      → Unexported
```

This is based on **package visibility**, not whether code is in the same file.

---

## 5. `go.mod`

`go.mod` identifies the Go module/project and declares its dependencies and Go version.

Created using:

```bash
go mod init cloud-platform
```

A module might contain:

```text
module cloud-platform

go 1.XX
```

Dependencies are also declared in `go.mod`.

`go.sum` contains checksums used to verify dependency contents.

### Important distinction

```text
go.mod → module identity + dependencies + Go version
go.sum → dependency checksums
```

`go mod tidy` cleans up and synchronizes module dependencies.

---

# Day 96 — Core Mental Model

```text
Files
  ↓
Read / Write external data

JSON
  ↓
Convert between JSON and Go structs

Environment Variables
  ↓
External application configuration

Packages
  ↓
Organize and reuse Go code

go.mod
  ↓
Manage the Go module and dependencies
```

## Key Interview Questions

### What is `json.Marshal()`?

It converts a Go value, such as a struct, into JSON represented as `[]byte`.

### What is `json.Unmarshal()`?

It converts JSON data into a Go value, usually by populating a struct.

### Why do we pass `&service` to `json.Unmarshal()`?

Because `Unmarshal()` needs access to the actual variable so it can populate its fields.

### What does `os.Getenv()` return if the variable doesn't exist?

An empty string.

### What is the difference between exported and unexported identifiers?

Identifiers beginning with an uppercase letter are exported and accessible from other packages. Lowercase identifiers are unexported and accessible only within the same package.

### What is `go.mod`?

It defines the Go module and tracks its Go version and dependencies.

### What is `go.sum`?

It stores checksums for dependencies so their contents can be verified.