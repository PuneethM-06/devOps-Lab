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
