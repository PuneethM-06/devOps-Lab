
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
15. ### Passing functional parameters
```
func main (){
	sayHello("Puneeth")
}
func sayHello(name string) {
    fmt.Println("Hello ", name)
}
```
15. ### return values - functions 
``` function add(a int, b int)int{
    return a+b
}
```
- a int → parameter a is an integer
- b int → parameter b is an integer
- final int → function returns an integer
- return a + b → sends the result back

16. ### Multiple return values
```
func divide(a, b int) (int, error) {
    ...
}
func main() {
    result, err := divide(10, 20)
}

```
- Example:
```
func calculate (a int, b int)(int, int){
	return a + b, a - b
}
func main(){
	sum, diff := calculate(10,20)
	fmt.Println("sum:", sum, "Diff:", diff)
}
```
```
 func divide(a int, b int)(int, error){
	if b == 0{
		return 0, errors.New("cannot divide by zero")
	} else{
		return a/b, nil 
	}
 }
 func main(){
	div, err := divide(10, 2)
 }
 ```

 17. ### Error propogation 
 - Pass the error upward until the layer that actually knows what to do with it can handle it.
 ```
 LOWER LEVEL
   ↓
DISCOVER error
   ↓
RETURN error
   ↓
HIGHER LEVEL
   ↓
RETURN error
   ↓
APPROPRIATE LEVEL
   ↓
HANDLE error
```
- Example:
```
func getPort()(int, error) {
	return 0, errors.New("port not configured")
}

func startService()error {
	result, err := getPort()
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
func main(){
	err := startService()
	if err != nil {
		fmt.Println(err)
	}
}
```
18. ### structs
- Lets you create your own data type made up of mutiple fields 
- Example
```
type Service struct {
    Name       string
    Port       int
    Production bool
    Status     string
}
```
In the above example, we are **creating struct**
```
service := Service{
    Name:       "api",
    Port:       8080,
    Production: true,
    Status:     "running",
}
```
- **defining service**

19. ### Accessing fields 
- we can access using `.` operator
- Example:
`fmt.Println(service.Name)`

20. ### Changing fields
- Example: `Service.Status = "stopped"`
