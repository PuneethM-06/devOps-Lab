# DAY 96 - GOLANG - FILE HANDLING, JSON, ENVIRONMENT VARIABLES PACKAGES

1. ### File handling 
- **File contains bytes, so Go reads the file as bytes**
- Example: `data, err := os.ReadFile("config.txt")`

2. ### File handling - Writing 
- This operation allows you to write the file 
- Example: `os.WriteFile("output.txt", data, 0644)`
```
"output.txt" → filename
data         → content to write
0644         → file permissions
```
- Example:
```
data := []byte("Cloud Platform kit")

err := os.WriteFile("output.txt", data, 0644)
if err != nil{
	fmt.Println("Failed to write file")
	return
}
fmt.Println("file written successfully ")
```

3. ### JSON 
- When an JSON is receievd there are 2 ways of converting 
```
JSON → Go struct
Go struct → JSON
```
- These have names,
1. **JSON -> Go is called Unmarshal**
2. **Go -> JSON is called Marshal**

4. ### Marshall - Go -> 
- Go's standard library has the `encoding/json` package 
- Example:
```
type Service struct {
	Name string 
	Port int 
	Status string 
}

service := Service {
	Name: "api",
	Port: 8080,
	Status: "running"
}
data, err := json.Marshal(service)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(string(data))
```

5. ### UnMarshall
- Here we will be turning JSON to Struct 
```
JSON
 ↓
json.Unmarshal()
 ↓
Go struct
```
- err := json.unmarshall(data, &service)
- Example:
```
data := []byte(`{
  "repository": "github.com/mycompany/payment-service",
  "branch": "main"
}`)
type ScanRequest struct {
	Repository string 
	Branch string
}
var scanRequest ScanRequest 
err := json.Unmarshal(data, &scanRequest)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(scanRequest.Repository, scanRequest.Branch)
```
5. ### Json Tags 
- Currently we have
```
type ScanRequest struct {
    Repository string
    Branch     string
}
```
- Go will normally map JSON like:
```
{
  "Repository": "...",
  "Branch": "..."
}
```
- but API conventionally use lowercase names 
```
{
  "repository": "...",
  "branch": "main"
}
```
- **So we can explicitly tell Go which JSON names belongs to each field**
```
type ScanRequest struct {
    Repository string `json:"repository"`
    Branch     string `json:"branch"`
}
```
```
type ScanRequest struct {
    Repository string `json:"repository"`
    Branch     string `json:"branch"`
}
```
- This works for both, Marshall and Unmarshall

6. ### Environment Variables 
- Environment variables are used instead of using hardcoded values 
- Example:
```
PORT=8080
ENVIRONMENT=production
AWS_REGION=ap-south-1
```
**Reading an Environment variable***
`os.Getenv("PORT")`
- `os.Getenv` returns a String always 
- Example:
```
APP_ENV=production
environment := os.Getenv("APP_ENV")
```
- If APP_ENV does not exist then it returns an **empty string**
- So we can write 
```

APP_ENV=production
environment := os.Getenv("APP_ENV")

if environment == "" {
	fmt.Println("APP_ENV does not exist")
}
```

7. ### Packages 
- Every go file belongs to a package.
- Example: `package main`
- **Package is a way to group related Go code**

8. ### Using code from another package 
- Example
```
package config

func GetPort() int {
    return 8080
}
```
- GetPort() -> capital G ->. It means that it is accessible outside config 
- getPort() -> small g -> It means that it is only usable inside config 