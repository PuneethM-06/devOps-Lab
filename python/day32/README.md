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
