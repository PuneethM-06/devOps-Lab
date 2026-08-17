# DAY 86 0 SNYK: Scan Go Dependencies and Auto-Fix PRs

1. ### WHAT IS SNYK??
- Snyk is a **developer secuity platform that helps find security vulnerabilities in software**
```
Go project
   ↓
go.mod 
   ↓
Snyk
   ↓
Find vulnerable dependencies
```

2. ### GO DEPENDENCIES: go.mod, direct and indirect dependencies 
- A Go Project does not download and manage every dependency file-by-file 
- Instead, Go uses modules such as:
```
my-go-app/
├── main.go
├── go.mod
└── go.sum
```
- **go.mod**
```
module github.com/puneeth/my-app

go 1.25

require github.com/gin-gonic/gin v1.11.0
```
- In the above example, This tells go **My application is directly dependent on this module and version**
- This is called **direct dependency**
- Whereas, `gin` might be dependent on another modules and that becomes **indirect dependency** 
```
Your Application
        ↓
      Gin
        ↓
 ┌──────┼──────┐
 ↓      ↓      ↓
Dep A  Dep B  Dep C
```
- **DIRECT DEPENDENCY** - A dependency our application directly imports
- **INDIRECT DEPENDENCY** - Required by one of our dependency 
