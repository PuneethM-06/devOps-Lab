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

### create
`with open("servers.txt", "x") as file:`

### read
```
with open("servers.txt", "r") as file:
    data = file.read()
```
### readline
- reads one line 
```
with open("servers.txt") as file:
    line = file.readline()
```
### readlines()
```
with open("servers.txt") as file:
    lines = file.readlines()
```
- returns a list

### WRITING MULTIPLE LINES
- Instead of:
```
file.write("web01\n")
file.write("web02\n")
file.write("db01\n")
```
- Use:
```
servers = [
    "ABC",
    "DEF"
]
file.writelines(servers)
```

### tell()
- Used to say the current position in the file 

### seek()
- file.seek(0) - moves pointer back to the beginning 

### File Encoding 
```
with open("servers.txt", encoding="utf-8") as file:
```

## MODULE 2 PATHLIB
- Instead of:
```
import os

if os.path.exists("config.json")"
    print("Found")
```
- We can do:
```
from pathlib import Path

path = Path("config.json")

if path.exists():
    print("Found")
```
- We can also do
```
path.exists()      # Does it exist?
path.is_file()     # Is it a file?
path.is_dir()      # Is it a directory?
path.mkdir()       # Create a directory
path.unlink()      # Delete the file
path.read_text()   # Read the file
path.write_text()  # Write to the file
```

## MODULE 3 - argparse
- Accept command-line arguements instead of hardcoding it 

- Bad:
```
server = "web01"
env = "prod"
```
- Better:
`python deploy.py --server web01 --env prod`

## MODULE 4: SUBPROCESS
- Execute shell commands from python 

- instead of asking the user to run we can execute using subprocess

```
import subprocess

subprocess.run(["df", "-h"], capture_output=True, text=True)
```

### RETURNCODE

```
import subprocess
result = subprocess.run(["ls"], capture_output=True, text=True)

print(result.returncode)
```
- 0 means success
- Non-Zero - Failure

### stderr
- Executed when a command fails to understand the error or reason for failing 
```
import subprocess
result = subprocess.run(["ls"], capture_output=True, text=True)

print(result.stderr)
```

### TIMEOUTS
```
subprocess.run(
    ["sleep", "10"],
    timeout=5
)
```
- Raises an error when timeout occurs

## MODULE 5 JSON 
- Read, Write and parse JSON data in python 

| Python         | JSON        |
| -------------- | ----------- |
| `dict`         | Object `{}` |
| `list`         | Array `[]`  |
| `str`          | String      |
| `int`, `float` | Number      |
| `True`         | `true`      |
| `False`        | `false`     |
| `None`         | `null`      |

- Example:
```
server = {
    "name": "web01",
    "cpu": 80
}
```

```
{
    "name": "web01",
    "cpu": 80
}
```

## JSON LOADS
- Converts a JSON string to python object
```
import json 
data = '{"name":"Puneeth", "age":"32"};

server = json.loads(data)
print(server)

output:
{
    'name':'Puneeth',
    'age':'32'
}
```
## JSON DUMPS
- Converts a Python Object to Json string 
```
server = {
    'name':'Puneeth',
    'age':'32'
}
dumps = json.dumps(server)
print(dumps)

output:
{
    "name":"Puneeth",
    "age":"32"
}
```

## READING JSON FILES
- config.json
```
{
    "region": "us-east-1",
    "instance": "web01"
}
```
```
import json

with open("config.json","r") as file:
    config = json.load(file)
print(config["region"])
```

## WRITING JSON FILES
```
import json

config = {
    "name":"us-east-1"
}
with open("config.json", "w") as file:
    json.dump(config, file, indent=4)
```

1. load() vs loads()?
load() → Reads JSON from a file.
loads() → Reads JSON from a string.

2. dump() vs dumps()?
dump() → Writes JSON to a file.
dumps() → Converts an object to a JSON string.


### Example of json load 
```
suppose we have a config.json
{
    "region": "us-east-1",
    "instance": "web01"
}

import json

with open("config.json", "r") as file:
    config = json.load(file)

print(config)

ouput:
{
    "region": "us-east-1",
    "instance": "web01"
}
```
### Example of json loads 
```
import json

data = '{"region":"us-east-1","instance":"web01"}'

config = json.loads(data)

print(config)

{
    "region": "us-east-1",
    "instance": "web01"
}
```
### json.dump
```
config
config = {
    "region": "us-east-1",
    "instance": "web01"
}

import json

with open("config.json", "w") as file:
    json.dump(config, file, indent=4)

output:
{
    "region": "us-east-1",
    "instance": "web01"
}
```

### json dumps
```
import json

config = {
    "region": "us-east-1",
    "instance": "web01"
}

json_string = json.dumps(config)

print(json_string)

{"region": "us-east-1", "instance": "web01"}
```
