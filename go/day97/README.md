# Day 97 — HTTP Client/Server, REST APIs, Routing, and Middleware

## 1. HTTP Server

An HTTP server waits for HTTP requests and sends HTTP responses.

Go provides the `net/http` package for building HTTP clients and servers.

The basic command used to start an HTTP server is:

```go
http.ListenAndServe(":8080", nil)
```

This starts a server that listens for incoming HTTP traffic on port `8080`.

Since `http.ListenAndServe` returns an error, it should normally be handled:

```go
log.Fatal(http.ListenAndServe(":8080", nil))
```

---

## 2. Handler

A handler is code that runs when an HTTP request reaches the server.

```text
Client
  │
  │ GET /health
  ▼
Go HTTP Server
  │
  ▼
Handler
  │
  ▼
Response
```

A handler function looks like this:

```go
func handler(w http.ResponseWriter, r *http.Request) {
}
```

It has two parameters:

### `r *http.Request`

`r` represents the incoming HTTP request.

The request contains information such as:

- HTTP method
- URL
- Headers
- Query parameters
- Request body

For example:

```go
r.Method
r.URL
r.Header
r.Body
```

### `w http.ResponseWriter`

`w` is used to send the HTTP response back to the client.

For example:

```go
fmt.Fprintln(w, "Hello from the server")
```

The request and response flow can be summarized as:

```text
r → incoming request
w → outgoing response
fmt.Fprintln → writes to w → response sent to the HTTP client
```

---

## 3. Connecting a Route to a Handler

After creating the server and handler, we must register the handler with a route.

```go
http.HandleFunc("/", handler)
```

This means:

> When a request arrives at `/`, run the `handler` function.

A complete example is:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the server")
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### Important

1. `http.HandleFunc` registers a route and its handler.
2. `http.ListenAndServe` starts the HTTP server.
3. Passing `nil` to `http.ListenAndServe` tells Go to use `http.DefaultServeMux`, which contains the routes registered using `http.HandleFunc`.

---

## 4. Restricting HTTP Methods

Without method restrictions, all of the following requests can reach the `/health` handler:

```text
GET /health
POST /health
PUT /health
DELETE /health
```

Suppose `/health` should accept only `GET` requests. We can check `r.Method` and reject other methods:

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
		return
	}

	fmt.Fprintln(w, "Server is healthy")
}
```

`http.MethodGet` is preferred over writing `"GET"` directly because it avoids spelling mistakes and clearly expresses the intended HTTP method.

In Go 1.22 and later, the method can also be included in the route pattern:

```go
http.HandleFunc("GET /health", healthHandler)
```

---

## 5. HTTP Status Codes

Common HTTP status codes include:

```text
200 → Request succeeded
201 → Resource created successfully
204 → Request succeeded, but there is no response body
400 → Client sent an invalid request
401 → Client is not authenticated
403 → Client is authenticated but not allowed to perform the action
404 → Resource or route was not found
405 → HTTP method is not allowed
500 → Server encountered an internal error
```

Go provides named constants for status codes:

```go
http.StatusOK                  // 200
http.StatusCreated             // 201
http.StatusNoContent           // 204
http.StatusBadRequest          // 400
http.StatusUnauthorized        // 401
http.StatusForbidden           // 403
http.StatusNotFound            // 404
http.StatusMethodNotAllowed    // 405
http.StatusInternalServerError // 500
```

To explicitly send a status code from a handler:

```go
w.WriteHeader(http.StatusCreated)
```

The status code must be written before the response body:

```go
w.WriteHeader(http.StatusCreated)
fmt.Fprintln(w, "Resource created")
```

If no status code is explicitly written, Go automatically sends `200 OK` when the response body is written.

---

## 6. Checking an HTTP Client Response

When an HTTP client receives a response, its status code can be checked using `resp.StatusCode`.

```go
if resp.StatusCode != http.StatusOK {
	// Handle the unexpected status code.
}
```

Example:

```go
resp, err := http.Get("http://localhost:8080/health")
if err != nil {
	log.Fatal(err)
}
defer resp.Body.Close()

if resp.StatusCode != http.StatusOK {
	log.Printf("Unexpected status code: %d", resp.StatusCode)
}
```

The response body should be closed using:

```go
defer resp.Body.Close()
```

This allows the HTTP client to release or reuse the underlying connection properly.

### 7. POST
- Example:
```
resp, err := http.Post(
    "http://localhost:8080/scan",
    "application/json",
    body,
)
```
- There are three things:
    1. Where we are sending it
    2. What kind of data are we sending
    3. Body  - The actual data 
