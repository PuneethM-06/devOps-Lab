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

