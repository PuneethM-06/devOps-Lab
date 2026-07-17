# Day 32 – Module 5: JSON (Interview Notes)

## What is JSON?

JSON (JavaScript Object Notation) is a lightweight data-interchange format used by:

* REST APIs
* AWS SDK (boto3)
* Kubernetes
* Docker
* GitHub API
* Configuration files

Python represents JSON using dictionaries and lists.

---

# Python vs JSON

| Python      | JSON        |
| ----------- | ----------- |
| dict        | Object `{}` |
| list        | Array `[]`  |
| str         | String      |
| int / float | Number      |
| True        | true        |
| False       | false       |
| None        | null        |

---

# json.load()

**Purpose:** Read JSON from a file and convert it into a Python object.

Example

```python
import json

with open("config.json", "r") as file:
    config = json.load(file)

print(config)
```

Input:

* JSON File

Output:

* Python Dictionary

---

# json.loads()

**Purpose:** Convert a JSON string into a Python object.

Example

```python
import json

data = '{"region":"us-east-1","instance":"web01"}'

config = json.loads(data)

print(config)
```

Input:

* JSON String

Output:

* Python Dictionary

---

# json.dump()

**Purpose:** Write a Python object into a JSON file.

Example

```python
import json

config = {
    "region": "us-east-1",
    "instance": "web01"
}

with open("config.json", "w") as file:
    json.dump(config, file, indent=4)
```

Input:

* Python Dictionary

Output:

* JSON File

---

# json.dumps()

**Purpose:** Convert a Python object into a JSON string.

Example

```python
import json

config = {
    "region": "us-east-1",
    "instance": "web01"
}

json_string = json.dumps(config)

print(json_string)
```

Input:

* Python Dictionary

Output:

* JSON String

---

# Easy Way to Remember

**The letter "s" stands for String.**

| Method  | Reads/Writes | Source/Destination          |
| ------- | ------------ | --------------------------- |
| load()  | Read         | JSON File → Python Object   |
| loads() | Read         | JSON String → Python Object |
| dump()  | Write        | Python Object → JSON File   |
| dumps() | Write        | Python Object → JSON String |

---

# Common Interview Questions

### Q1. Difference between load() and loads()?

* `load()` reads JSON from a file.
* `loads()` reads JSON from a string.

---

### Q2. Difference between dump() and dumps()?

* `dump()` writes JSON to a file.
* `dumps()` converts a Python object into a JSON string.

---

### Q3. Why is JSON important?

JSON is the standard format used by REST APIs, AWS services, Kubernetes, Docker, GitHub, and most modern cloud applications.

---

### Q4. Which methods are used most in DevOps?

* `json.load()` → Read configuration files.
* `json.loads()` → Parse JSON strings from APIs.
* `json.dump()` → Save reports/configuration.
* `json.dumps()` → Prepare request payloads for APIs.

---

# Real DevOps Examples

### Read Configuration

```python
with open("config.json") as f:
    config = json.load(f)
```

---

### Parse API Response

```python
data = json.loads(api_response)
```

---

### Save Report

```python
with open("report.json", "w") as f:
    json.dump(report, f, indent=4)
```

---

### Send API Payload

```python
payload = {
    "name": "web01",
    "status": "running"
}

body = json.dumps(payload)
```

---

# Quick Revision

* `load()` → File → Python Object
* `loads()` → String → Python Object
* `dump()` → Python Object → File
* `dumps()` → Python Object → String

## Interview Trick

**"s" = String**

If there is an **`s`**, it works with **strings**.
If there is **no `s`**, it works with **files**.

"JSON stands for JavaScript Object Notation. It is a lightweight data-interchange format used to exchange structured data between systems. In Python, JSON is typically converted into dictionaries and lists using the json module. In DevOps, JSON is widely used for REST APIs, AWS services, Kubernetes, Docker, GitHub APIs, and configuration files, making it one of the most important data formats to understand."
