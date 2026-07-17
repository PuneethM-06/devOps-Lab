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

### NOTE: is compares object identity (whether two variables refer to the same object in memory).

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
- Unlike Java, python doesn't use &&, ||, ! for logical operators

4. ### MEMBERSHIP OPERATORS
```
in
not in
```
- Example:
```
servers = ["web01", "web02", "db01"]

if "web01" in servers:
    print("Found")
```


5. ### ASSIGNMENT OPERATORS
```
count += 1
count -= 1
count *= 2
count /= 2
```

### TERNARY OPERATORS
- Instead of
```
if cpu > 80:
    status = "High"
else:
    status = "Normal"
```
- Write
```
status = "High" if cpu > 80 else status = "Normal"
```
## MODULE 3 -  STRINGS
1. ### CREATING STRINGS
- Python supports:
    - single quotes ' ' 
    - double quotes "" 
    - tripe quotes for multi line strings """" """

2. ### STRING FORMATING 
- Traditional way of handling 
```
server = "web01"
print("checking" + server)
```
- MODERN WAY OF HANDLING IS USING fstrings
```
server = "web01"
print(f"checking + {server}")
```
- Multiple variables
```
print(f" Checking CPU={CPU}% and MEM={MEM}%")
```

### COMMON STRING METHODS
1. upper() - convert to uppercase
2. lower() - convert to lowercase
3. strip() - rmeove trailing whitespaces
4. replace() - replace("string to be replaced", "new string")
5. split() - split strings - split(",")
6. join() - Opposite of split - print(",".join(servers))
7. startswith() - search strings - filename.startswith("error")
8. endswith() - search strings that end with - filename.endswith(".log")
9. string slicing - string[start:end]
10. Length - len(str)
11. Escape character - print("Hello\nworld")

### STRINGS ARE IMMUTABLE IN NATURE AND CANNOT BE MODIFIED AT DESIRED LOCATIONS 

### MODULE 4 - LISTS
- Unlike strings, Lists are mutable and can be done by making use of index positions 

1. Creating a List - `servers = ["web01", "web02", "db01"]`
2. Accessing Elements - 
```
servers = ["web01", "web02", "db01"]

print(servers[0])
```
3. Slicing lists
```
servers = ["web01", "web02", "web03", "db01"]

print(servers[:2])
```
4. Adding elements
```
servers.append("cache01")
```
5. Removing elements
```
servers.remove("web02")
```
6. Length
```
len(servers)
```
7. Looping through lists
```
for server in servers:
    print(server)
```
- With index
```
for index, server in enumerate(servers):
    print(index, server)
```
8. sorting 
```
servers.sort()
```
9. Reverse
```
servers.reverse()
```
10. Copying lists
```
a = ["web01", "web02"]

b = a
```
- Note: Both variables point to the same list and hence doing `b.append("ABC")` now a will be ``` a = ["web01", "web02", "ABC"]`

### LIST COMPREHENSION
- One of python's best feature
- Instead of
```
sqaures = [] 
for x in range(5):
    squares.append(x*x)
```
- We can do
```
sqaures = [x*x for x in range(5)]
```

### NOTE:
1. `sort` Modifies the exisiting list, while `sorted` creates a new list
2. `append` appends to last of the list while `insert(0, "ABC")` inserts at a specific location

## MODULE 5: DICTIONARIES

1. Creating dictionaries
```
server = {
    "name": "web01",
    "cpu": 78,
    "memory": 62
}
```
2. Accessing values
```
print(server["cpu"]) or print(server.get("cpu"))
```

3. updating values
` server["cpu"] = 90`

4. Adding new keys
`server["disk"] = 90`

5. Removing keys
` server.pop("disk")

6. Looping through dict
```
for key in servers:
    print(key)
```
or
```
for key, value in server.items():
    print(key, valur)
```

### NESTED DICT
```
instance = {
    "id": "123"
    "name": "puneeth"
    "state": {
        "name": "ABC"
    }
}

Accessing them 
print(instance["state"]["name"])
```

### LIST OF DICT
```
instances = [
    {
        "id": "i-111",
        "state": "running"
    },
    {
        "id": "i-222",
        "state": "stopped"
    }
]
```
```
Loop:
for instance in instances:
    print(instances["id"])
```
### CONDITIONS IN PYTHON 
- Already known and hence skipping the notes here

### LOOPS
1. for Loop:
- Basic Examples:
```
servers = ["web01", "web02", "db01"]

for server in servers:
    print(server)

for i in range(5):
    print(i)
```

2. enumerate()
- Very useful when you need both the index and the value.
```
for index, server in enumerate(servers):
    print(index, server)
```

3. while Loop
- Runs until a condition becomes false
```
count = 1

while count <= 5:
    print(count)
    count += 1
```

4. break
- stops the loop immediately
```
servers = ["web01", "web02", "db01"]

for server in servers:
    if server == "web02":
        break
    print(server)
```
5. continue
- skip the current iteration 
```
for server in servers:
    if server == "web02":
        continue
    print(server)
```
6. pass
- Placeholder that does nothing 
```
if cpu > 90:
    pass
```
## FUNCTIONS 
- writes reusable blocks of code 
- It can take single and multiple parameters. Not adding in notes since known 
- `**kwargs` accept any number of arguements
```
def create_server(**config):
    print(config)

create_server(cpu=4, memory=16)

output:
{
    "cpu": 4,
    "memory": 16
}
```

## MODULE 9 SCOPE
- understand where variables exist and who can access them.
1. Variable created inside a function exist inside that function 
2. Variables defined outside the functions are global 

### MODIFYING GLOBAL VAR
```
count = 0

def increment():
    count += 1

output: UnboundLocalError

Right way of doing 
count = 0

def increment():
    global count
    count += 1
```

- Whenever python looks for variable it looks in this order 
- L → Local
- E → Enclosing
- G → Global
- B → Built-in

### CONSTANTS
- Instead of globals that change we define CONSTANTS
- Example:
```
MAX_RETRIES = 3
AWS_REGION = "us-east-1"
DEFAULT_TIMEOUT = 30
```

## MODULE 10 - ERROR HANDLING 

- Exception is an error that occurs while running the program 
- Basic syntax:
```
try:
    num = int("100")
except ValueError:
    print("Invalid number")
```
