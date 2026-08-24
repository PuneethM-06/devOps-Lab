
# DAY 95 - GoLang Basics

1. ### Why do DevOps engineers prefer Go?
- Go has a combination that makes it extremely well suited to infrastructure and platform engineering.

1. **Compiles into a standalone binary**
- This is huge for DevOps, we can compile an code to a **executable binary without needing help of any compilers etc.**

2. **Fast and lightweight**
- Go is compiled and has **lower runtime overhead** than interpreted scripting languages

3. **Excellent Concurrency Support**
- Can handle **100's of concurrent workloads**

4. **Excellent networking support**
- Has **strong netwoking primitives**

2. ### Variables and constants
- A variable is simply a named place to store progammatic values
- Example:
```
name := "Puneeth"
age := 25
isEngineer =: true
```
3. ### Declaring a variable with var
- The **explicit syntax is**
` var name string = "Puneeth"`
```
var       name       string       =       "Puneeth"
│          │           │                    │
│          │           │                    └── value
│          │           └── data type
│          └── variable name
└── keyword saying "declare a variable"
```
**Note** - It is not mandate to infer the type in Go and it can do it for itself
- Example: `var name = "Puneeth"`

3. ### The shorthand `:=`
- This is called as the shorthand and instead of writing `var name string = "Puneeth"`, we can do `name := "Puneeth"`
- It is a shortend way of saying **Create this variable and infer its type from the value**

**Note** - Shorthand is used only for variables

4. ### Constants
- A constant is a value that should not change 
- Example: `const maxRetries = 5` or `const defaultPort = 8080` 

5. ### Go Important: Zero Values
- If a variable is declared without giving its value, Go gives it a **value Zero**
- Example:
```
var name string 
var age int 
var active bool
```
Go automatically gives these values
```
string → ""
int    → 0
bool   → false
```

6. ### Variables can change 
```
package main

import "fmt"

func main() {
    service := "nginx"

    fmt.Println(service)

    service = "api"

    fmt.Println(service)
}
```
- service gets value ngnix and later it is changed to api 

7. ### Data Types
| Type      | Stores          | Example   |
| --------- | --------------- | --------- |
| `string`  | Text            | `"hello"` |
| `int`     | Whole numbers   | `8080`    |
| `float64` | Decimal numbers | `99.5`    |
| `bool`    | True/false      | `true`    |

8. ### Type conversion
- There will be use cases for type conversion 
- Example:
```
port := 8080
portString := port
```
**Numeric conversion**
```
x := 10
y := float64(x)
```

9. ### Arrays
- An array has a **fixed size**
- Example:
```
ports := [3]int{8081,8082,8083}
```

**But arrays aren't something that is used most commonly, instead we make uses of slices**
- Example:
```
ports := []int{8080, 8081, 8082}
```
Unlike Arrays here we have `[]` instead of `[3]`
- Example:
```
services := []string{
    "api",
    "worker",
    "frontend"
}
```
**slices ensure we dont need to know the capacity in-hand before**

10. ### MAPS
- Maps are used based on `key-value` pair 
- Example:
```
serviceStatus := map[string]string {
    "api": "running",
    "worker": "failed"
}
```
- `serviceStatus[api] gives you running`
- In devOps, maps are extrememly useful
- Example:
```
servicePort := map[string]int{
    "api": 8080,
    "worker": 8081
}
```

11. ### Control Flow
**if**
- Basic syntax is:
```
if condition {

}
```
- Example:
```
production := true
if production {
    fmt.Println("Prod is true")
}
```
**Operators used**
```
==    equal
!=    not equal
>     greater than
<     less than
>=    greater than or equal
<=    less than or equal
```
and **Logical Operators**
```
&&    AND
||    OR
!     NOT
```
**if else**
- Basic syntax is
```
if condition {

}else {

}
```
**In Go, else must be on the same line as the closing } of the if block.**

**else if**
- basic syntax
```
if condition1 {

}else if condition 2{

}else{

}
```
12. ### for loops 
- Unlike laguages such as python or Java, Go does not have `while` or `do-while`, Go basically uses `for loop` for all looping
- Example:
```
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```
- Lets break that apart:
```
for i := 0; i < 5; i++ {
    │    │          │
    │    │          └── increment i
    │    └── continue while this is true
    └── create i
```

13. ### looping over collections
- Example:
```
services := []string{"api", "worker", "frontend"}

for index, service := range services {
    fmt.Println(index, service)
}
```
13. ### Switch 
- A switch is useful when we are checking one value against multiple possible values 
```
switch status {
case "running":
    fmt.Println("Service is running")
case "stopped":
    fmt.Println("Service is stopped")
case "failed":
    fmt.Println("Service has failed")
default:
    fmt.Println("Unknown status")
}
```
14. ### functions
- A function is a named block of code that performs a specific task 
- Basic syntax
```
func sayHello() {
    fmt.Println("Hello")
}
```
```
func       sayHello       ()       {
│              │           │        │
│              │           │        └── function body
│              │           └── parameters
│              └── function name
└── keyword meaning "function"
```
