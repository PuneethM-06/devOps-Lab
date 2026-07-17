# Day 32 – JSON Interview Questions

## Question 1

**What is JSON, and why is it important in Python and DevOps?**

---

## Question 2

**What is the difference between `json.load()` and `json.loads()`?**

- Explain the difference.
- Give one example for each.

---

## Question 3

**What is the difference between `json.dump()` and `json.dumps()`?**

- Explain the difference.
- Give one example for each.

---

## Question 4

**Why do `loads()` and `dumps()` have an `s` at the end? What does the `s` represent?**

---

## Question 5

Suppose you have a file named **`config.json`**.

**How would you read it into a Python dictionary?**

Write the code.

---

## Question 6

An API returns the following response as a JSON string:

```text
'{"cpu":80,"memory":60}'
```

**Which JSON method would you use before accessing `cpu`, and why?**

Write the code.

---

## Question 7

You have the following Python dictionary:

```python
report = {
    "status": "SUCCESS",
    "servers": 5
}
```

You want to save it as **`report.json`**.

**Write the code you would use.**

---

## Question 8

You have the following Python dictionary:

```python
payload = {
    "name": "web01",
    "status": "running"
}
```

You need to send it as the **body of an HTTP POST request**.

- Which JSON method would you use?
- Why?

---

## Question 9

Explain the difference between:

```python
json.load(file)
```

and

```python
json.loads(data)
```

Don't just say "file" and "string".

Explain:
- What each method does.
- When you would use each one.

---

## Question 10 (Most Important)

Complete the table from memory.

| Method | Input | Output |
|---------|-------|--------|
| `json.load()` | ? | ? |
| `json.loads()` | ? | ? |
| `json.dump()` | ? | ? |
| `json.dumps()` | ? | ? |

---

# Bonus Question (Real DevOps)

You're writing a Python automation script that:

1. Reads `config.json`.
2. Calls an AWS API.
3. Saves the response to `report.json`.

**Which JSON methods would you use at each step, and why?**

---

# Interview Tips

### Easy Trick

**"s" = String**

| Method | Remember |
|---------|----------|
| `load()` | Read from a File |
| `loads()` | Read from a String |
| `dump()` | Write to a File |
| `dumps()` | Convert to a JSON String |

---

### Common Follow-up Questions

- Why is JSON important in DevOps?
- Where is JSON used in AWS?
- Why is `json.dumps()` used before sending an HTTP request?
- Why is `json.load()` used for configuration files?
- What Python data structures does JSON become?
- Why is `.get()` preferred over `[]` when reading API responses?