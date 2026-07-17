## DAY 31 PYTHON BASICS

# MODULE 1: VARIABLES AND DATA TYPES

### VARIABLES
- Unlike J?ava, In python, the data type is determined at the run time and hence we can declare the variable directly 
- Example: `name = "Puneeth"`

### DATA TYPES IN PYTHON
| Type  | Example           | Use Case              |
| ----- | ----------------- | --------------------- |
| str   | `"web01"`         | Hostnames, file paths |
| int   | `8080`            | Ports, counts         |
| float | `78.45`           | CPU %, memory usage   |
| bool  | `True`            | Status flags          |
| list  | `["web1","web2"]` | Server lists          |
| dict  | `{"Name":"EC2"}`  | JSON/API responses    |
| None  | `None`            | Missing values        |

- `type()` is the function that is used for checking the data type.
```
name = "Puneeth"
print(type(name))
```
### TPYPE CONVERSION
- It is very common in automation for making string to int and int to string 
```
port = "8080"
port = int(port)
```
- similarly we can do for float and other data types

### MULTIPLE ASSIGNEMENT 
- Instead of:
```
cpu = 70
usage = 80
mem = 90
```
- we can do 
```
cpu, usage, mem = 70, 80, 90 
```
### SWAPPING VARIABLES
- We can swap variables making use of multiple assignments
```
a = 10
b = 20

a, b = b, a
```

### NAMING CONVENTION
- Python uses snake_case instead of camel_casing unlike Java
- Example: `int_cpu_usge`

### IDENTITY vs EQUALITY
- This is a common interview question 
```
a = [1, 2]
b = [1, 2]

print(a == b)
```
- output: True
- The above is Equality

- We can also do like:
` a is b`
- this is called as identity

### NOTE: For `None` we can never do ` value == None` instead we have to do ` value is None`

## MODULE 2 - OPERATORS

1. ### ARTHAMETIC OPERATORS
```
a = 10
b = 3

print(a + b)   # 13
print(a - b)   # 7
print(a * b)   # 30
print(a / b)   # 3.333...
print(a // b)  # 3
print(a % b)   # 1
print(a ** b)  # 1000
```

2. ### COMPARISON OPERATORS
```
==
!=
<
>
<=
>=
```

3. ### LOGICAL OPERATORS
```
and
or
not
```
