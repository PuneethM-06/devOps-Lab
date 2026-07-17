# Day 33 – HTTP Requests with `httpx` & JSON

## Interview Question 1 (Very Common)

You have written the following code:

```python
import httpx

response = httpx.get("https://jsonplaceholder.typicode.com/posts/1")
```

### Question

1. What is stored in the `response` variable?
2. What information can you get from this object?
3. Is `response` a Python dictionary? Why or why not?

---

## Interview Question 2 (Coding)

Write a Python program that:

- Fetches all users from:

```
https://jsonplaceholder.typicode.com/users
```

- Prints only the users' names.

---

## Interview Question 3 (Nested JSON)

Given the following JSON response:

```json
[
    {
        "id": 1,
        "name": "Alice",
        "company": {
            "name": "Google"
        }
    },
    {
        "id": 2,
        "name": "Bob",
        "company": {
            "name": "Microsoft"
        }
    }
]
```

Assume:

```python
data = response.json()
```

Write Python code to print:

```
Alice works at Google
Bob works at Microsoft
```

---

## Interview Question 4 (Coding)

Write a Python program that:

- Sends a GET request to:

```
https://jsonplaceholder.typicode.com/posts/1
```

- If the request is successful (`status_code == 200`), print:

```
Title: ...
Body: ...
```

- Otherwise print:

```
Request Failed
```

---

## Interview Question 5 (Very Common)

What is the difference between:

```python
response.text
```

and

```python
response.json()
```

Explain:

- What each one returns.
- When you would use each one.

---

## Interview Question 6 (Coding)

Write a Python program that:

- Sends a GET request to:

```
https://jsonplaceholder.typicode.com/posts
```

- Uses a query parameter to fetch only posts from **userId = 3**.
- Prints only the title of each post.

---

## Interview Question 7 (Very Common)

What is the difference between these two lines?

```python
httpx.post(url, json=payload)
```

and

```python
httpx.post(url, data=payload)
```

Explain:

1. What each one sends.
2. What `Content-Type` header is used by default.
3. When you would use each one.

---

## Interview Question 8 (Nested JSON)

Assume this API returns:

```json
[
    {
        "name": "Alice",
        "skills": [
            "Python",
            "Docker",
            "AWS"
        ]
    },
    {
        "name": "Bob",
        "skills": [
            "Linux",
            "Kubernetes"
        ]
    }
]
```

Assume:

```python
data = response.json()
```

Write Python code to print:

```
Alice
  - Python
  - Docker
  - AWS

Bob
  - Linux
  - Kubernetes
```

---

## Interview Question 9 (Production)

Why should HTTP requests always be wrapped inside a `try-except` block?

Explain:

- Why network requests can fail.
- Difference between `httpx.RequestError` and `httpx.HTTPStatusError`.
- What `response.raise_for_status()` does.
- Why it is preferred in production code.

---

## Interview Question 10 (Coding – Production Ready)

Write a Python program that:

- Sends a GET request to an API.
- Uses a timeout of **5 seconds**.
- Raises an exception if the server returns a **4xx** or **5xx** status code.
- Prints the JSON response if the request succeeds.
- Gracefully handles:
  - Request errors
  - HTTP status errors

---

## Bonus Interview Question (Senior DevOps)

A REST API requires the following headers:

```text
Authorization: Bearer <token>
Accept: application/json
```

Write a Python program using `httpx` that:

- Sends a GET request with these headers.
- Parses the JSON response.
- Prints the value of the `"message"` field.
- Handles any HTTP or network errors gracefully.

---

## Bonus Interview Question (Real-World Scenario)

You need to fetch all repositories for a GitHub user using the GitHub REST API.

Explain:

1. Which HTTP method would you use?
2. What endpoint would you call?
3. How would you pass authentication?
4. How would you handle pagination if the user has more than 100 repositories?
5. How would you handle API rate limits?
