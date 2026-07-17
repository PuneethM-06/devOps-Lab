## DAY 33 -PYTHON

## MODULE 1 - HTTP REQUEST
- HTTP REQUEST is a standardized text sent by client over the internet for Server to get data or response
```
Python Program
        │
HTTP Request
        │
        ▼
API Server
        │
HTTP Response
        ▼
Python Program
```

### WHAT IS AN API?
- API = Application Programming interface
- API acts as an intermediate/bridge that helps in delivering the request from client to server

### HTTP METHODS
| Method | Purpose          | Real Example      |
| ------ | ---------------- | ----------------- |
| GET    | Read data        | Get user          |
| POST   | Create data      | Create issue      |
| PUT    | Replace          | Replace profile   |
| PATCH  | Update partially | Change username   |
| DELETE | Delete           | Delete repository |

### CRUD
- C - Create → POST
- Read → GET
- Update → PUT/PATCH
- Delete → DELETE

### INSTALLING HTTPX AND GET REQUEST
- Command used for installing HTTPX - `pip3 install httpx`
- Example:
```
import httpx
```

### GET REQUEST
```
import httpx

response = httpx.get("https://jsonplaceholder.typicode.com/posts/1")
print(response)
```
### UNDERSTANDING RESPONSE OBJECT

```
import httpx

response = httpx.get("https://jsonplaceholder.typicode.com/posts/1")

print(response.status_code)

print(response.headers)

print(response.text)

print(response.url)
```

## STATUS CODES
| Code | Meaning      |
| ---- | ------------ |
| 200  | Success      |
| 201  | Created      |
| 204  | No Content   |
| 400  | Bad Request  |
| 401  | Unauthorized |
| 403  | Forbidden    |
| 404  | Not Found    |
| 500  | Server Error |

### READING TEXT RESPONSE

```
response = httpx.get("
    "https://jsonplaceholder.typicode.com/posts/1"
)
print(response.text)
```
### JSON
- JSON stands for JavaScript object notation
- It is the string that is sent over network 

```
import httpx

x = httpx.get("https://jsonplaceholder.typicode.com/posts/1")
data = response.json()
print(data) or print(data["title"])

output:
{
'id':1,
'title':'...',
'body':'...',
'userId':1
}
```
## PASSING QUERY PARAMETERS
- Params are passed to extract specific data 

- Example:
```
import httpx

params = {
    "userId":1
}

response = httpx.get(
    "https://jsonplaceholder.typicode.com/posts",
    params=params
)
data = response.json()
print(data)
```

## CUSTOM HEADERS
```
headers = {
    "User-Agent": "MyApp"
}

response = httpx.get(
    "https://httpbin.org/headers",
    headers=headers
)

print(response.json())
```

### TIMEOUT
- Never let our program or wait for response forever and hence we make use of this

```
response = httpx.get(
    "https://example.com",
    timeout=5
)
```
## EXCEPTION HANDLING 
- Easy to catch errors and debug

```
import httpx

try:
    response = httpx.get(
        "https://example.com"
        timeout=5
    )
    response.raise_for_status()
    print(response.json())

except httpx.HTTPStatusError as e:
    print("HTTP error": e)
```
### Why raise_for_status()?
- without it, a 404 response doesnt raise an exception and we have to checkout ourselves and hence.

## POST REQUEST

```
import httpx

payload = {
    "title":"Hello",
    "body":"Testing",
    "userId":1
}

response = httpx.post(
    "https://jsonplaceholder.typicode.com/posts",
    json=payload
)
print(response.status_code)
print(response.json())
```

### NOTE:
1. What is the difference between response.text and response.json()?

Expected Answer

response.text returns the response body as a string.
respons

"response is an httpx.Response object returned by the server after making an HTTP request. It contains metadata such as the status code, headers, URL, cookies, and the response body. The body can be accessed as plain text using response.text or parsed into a Python dictionary using response.json(). The response object itself is neither JSON nor a Python dictionary."