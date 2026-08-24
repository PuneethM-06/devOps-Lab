
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
