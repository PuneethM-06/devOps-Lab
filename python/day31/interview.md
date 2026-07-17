# Python for DevOps – Interview Notes (Day 31)

## Variables & Data Types

### What is dynamic typing?
Python determines variable types at runtime, so you don't declare them explicitly.

### Example

```python
x = 10
x = "hello"
```

### Interview Questions

**Q. Difference between Python and Java variable declarations?**
- Java is statically typed.
- Python is dynamically typed.

**Q. Difference between `==` and `is`?**
- `==` compares values.
- `is` compares object identity (memory reference).

**Q. Why use `is None` instead of `== None`?**
`None` is a singleton object, so `is None` is the recommended and idiomatic check.

---

# Operators

## Important Operators

### Arithmetic
- `+`
- `-`
- `*`
- `/`
- `//`
- `%`
- `**`

### Comparison
- `==`
- `!=`
- `<`
- `>`
- `<=`
- `>=`

### Logical
- `and`
- `or`
- `not`

### Membership
- `in`
- `not in`

### Identity
- `is`
- `is not`

### Interview Questions

**Q. Difference between `/` and `//`?**

- `/` → Floating-point division
- `//` → Floor (integer) division

**Q. Difference between `==` and `is`?**

- `==` compares values.
- `is` compares object identity.

**Q. Why use `in`?**

Checks membership inside a list, tuple, set, or dictionary.

---

# Strings

## Frequently Used Methods

- `strip()`
- `split()`
- `join()`
- `replace()`
- `startswith()`
- `endswith()`
- `lower()`
- `upper()`
- `len()`

### Interview Questions

**Q. Why prefer f-strings?**

- More readable
- Faster
- Cleaner

Example

```python
print(f"CPU Usage: {cpu}%")
```

---

**Q. Difference between `split()` and `join()`?**

`split()`

```
String → List
```

`join()`

```
List → String
```

---

**Q. Are strings mutable?**

No.

Strings are immutable.

---

# Lists

## Frequently Used Methods

- `append()`
- `insert()`
- `remove()`
- `pop()`
- `sort()`
- `copy()`
- `len()`

### Interview Questions

**Q. append() vs insert()?**

- `append()` adds at the end.
- `insert()` adds at a specific index.

---

**Q. sort() vs sorted()?**

- `sort()` modifies the original list.
- `sorted()` returns a new sorted list.

---

**Q. Why use enumerate()?**

Provides both index and value while looping.

Example

```python
for index, server in enumerate(servers):
    print(index, server)
```

---

# Dictionaries

## Frequently Used Methods

- `get()`
- `items()`
- `keys()`
- `values()`
- `pop()`
- `copy()`

### Interview Questions

**Q. `dict["key"]` vs `dict.get("key")`?**

`dict["key"]`

- Raises `KeyError` if missing.

`dict.get("key")`

- Returns `None`
- Can return a default value

---

**Q. Why are dictionaries important?**

Almost every JSON object returned by an API becomes a Python dictionary.

---

**Q. Most common API structure?**

```python
[
    {
        "id": "i-123",
        "state": "running"
    }
]
```

(List of dictionaries)

---

# Conditionals

## Interview Questions

**Q. Why use `is None`?**

Recommended Python style for checking `None`.

---

**Q. elif vs multiple if?**

`elif`

- Stops after the first match.

Multiple `if`

- Evaluates every condition.

---

**Q. Why use truthy/falsy checks?**

Instead of

```python
if len(items) > 0:
```

Write

```python
if items:
```

Cleaner and more Pythonic.

---

# Loops

## Types

- `for`
- `while`

## Keywords

- `break`
- `continue`
- `pass`

### Interview Questions

**Q. for vs while?**

`for`

- Iterate over collections.

`while`

- Repeat until a condition changes.

---

**Q. break vs continue?**

`break`

- Exits the loop completely.

`continue`

- Skips the current iteration.

---

**Q. Why use enumerate()?**

Cleaner than

```python
for i in range(len(items)):
```

---

# Functions

## Important Concepts

- Parameters
- Return values
- Default arguments
- Keyword arguments
- `*args`
- `**kwargs`
- Type hints
- Docstrings

### Interview Questions

**Q. print() vs return?**

`print()`

- Displays output.

`return`

- Sends data back to the caller.

---

**Q. Why use functions?**

- Reusability
- Easier testing
- Better readability
- Easier maintenance

---

**Q. Why should functions return instead of print?**

Returned values can be reused.

Example

```python
result = add(10, 20)

if result > 20:
    ...
```

---

# Scope

## LEGB Rule

- Local
- Enclosing
- Global
- Built-in

### Interview Questions

**Q. Local vs Global variables?**

Local

- Exists only inside a function.

Global

- Defined outside functions.

---

**Q. Should we use global variables?**

Avoid them whenever possible.

Instead

```python
count = increment(count)
```

Not

```python
global count
```

---

**Q. What is LEGB?**

Python searches variables in this order:

1. Local
2. Enclosing
3. Global
4. Built-in

---

# Error Handling

## Important Keywords

- `try`
- `except`
- `raise`
- `finally`
- `else`

## Best Practices

✅ Catch specific exceptions.

```python
except FileNotFoundError:
```

✅ Raise your own exceptions when validating input.

```python
raise ValueError("Invalid CPU")
```

✅ Use `finally` for cleanup.

```python
finally:
    connection.close()
```

✅ Log errors instead of only printing them.

```python
logging.exception(e)
```

❌ Don't use

```python
except:
    pass
```

---

### Interview Questions

**Q. What is an exception?**

An error that occurs during program execution.

---

**Q. Difference between except and finally?**

`except`

- Runs only if an exception occurs.

`finally`

- Always runs.

---

**Q. Why avoid bare except?**

- Hides bugs
- Makes debugging difficult

Prefer

```python
except ValueError:
```

or

```python
except Exception as e:
```

---

**Q. When should you use raise?**

When your code detects invalid input or an unrecoverable condition.

Example

```python
if cpu > 100:
    raise ValueError("Invalid CPU")
```

---

**Q. Should every function have try/except?**

No.

Catch exceptions only:

- Around operations that can fail
- Where you can recover
- Where you want to add context

Handle unexpected failures in `main()`.

Example

```python
def main():
    try:
        process_servers()
    except Exception as e:
        logging.exception(e)
```

---

# Quick Revision (Most Important)

## Must Know

- Dynamic Typing
- `==` vs `is`
- f-Strings
- `split()` / `join()`
- List Comprehension
- `append()` / `sort()` / `copy()`
- `dict.get()`
- `dict.items()`
- Truthy/Falsy
- `for` / `while`
- `break` / `continue`
- Functions with `return`
- `*args` / `**kwargs`
- LEGB Rule
- `try`
- `except`
- `raise`
- `finally`
- Catch Specific Exceptions

---

# Common DevOps Pattern

```python
def main():
    try:
        config = load_config()
        instances = get_instances(config)
        process_instances(instances)
    except FileNotFoundError:
        logging.error("Configuration file missing")
    except Exception as e:
        logging.exception(e)

if __name__ == "__main__":
    main()
```

This is the pattern you'll commonly see in production Python automation scripts.