# Day 97 - HTTP Client/Server, REST APIs, Routing and Middleware

1. ### HTTP Server
- An HTTP Server simply waits for HTTP request and sends HTTP response
- The basic Go Package is `net/http`
- The basic Go command to start a server is **http.ListenAndServe(":8080",nil)**
- **Listen for incoming HTTP traffic onm port 8080**

2. ### Handler
- Handler is a code that runs when an HTTP request comes in
```
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
- Example:
```
func handler(w http.ResponseWriter, r *http.Request){

}
```
- It has two parameters:
    1. **r *http.Request**
        - This is responsible for **representing the incoming http request**
        - The incoming request contains a lot of things such as **Headers, Query Parameters, Request Body, Method etc.**
        - so we can do things like **r.Method, r.URL**
    2. **w http.ResponseWriter**
        - used to send the **response back to the client**
        
```
r → incoming request
w → outgoing response
```
