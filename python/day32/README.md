## PYTHON DAY 32

## FILE HANDLING

### 1. OPENING A FILE 
- Basic syntax:
```
file = open("servers.txt", "r")
data = file.read()
print(data)

```

### OPENING USING `with open()`
```
with open("servers.txt", "r") as file:
    data = file.read()
```

### FILE MODES
| Mode   | Meaning                              |
| ------ | ------------------------------------ |
| `"r"`  | Read (default)                       |
| `"w"`  | Write (creates or overwrites)        |
| `"a"`  | Append                               |
| `"x"`  | Create new file (fails if it exists) |
| `"rb"` | Read binary                          |
| `"wb"` | Write binary                         |

### write
```
with open("server.txt", "w") as file:
    file.write("Health OK")
```
- To keep in mind. Write erases all the exisiting file 

### append
with open("server.txt". "w") as file:
    file.write("Health OK\n")
