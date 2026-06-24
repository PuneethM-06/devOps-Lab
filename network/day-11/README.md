## NETWORK - DAY 11

## WHAT IS HTTP
- HTTP -  HyperText Transfer protocol
- It is a protocol that is used to communicate between client and server 

### HTTP IS STATELESS -  It does not rememeber from whom the request comes and each request is independent of each other

## HTTP STRUCTURE

1. **METHOD** - This tells the server which action to perform 
Example:
- GET
- POST
- PATCH
- DELETE

2. **PATH** - This tells the server which resources it wants 
Example:
- /users
- /orders

3. **HTTP VERSION** - Tells the version being used.

4. **HEADERS** - Headers provide extra information 
- Example:
1. Authorization
2. Content-Type

5. **BODY** - Optional 

### HTTP BODY - Already known
- PUT - replace all
- PATCH - Update specific field

## HTTP RESPONSES AND STATUS CODE

1. ### 2XX - SUCCESS
```
200 OK
201 Created
204 No Content
```

2. ### 3XX - REDIRECT
```
200 OK
201 Created
204 No Content
```

3. ### 4XX - ERROR
```
400 Bad Request
401 Unauthorized
403 Forbidden
404 Not Found
429 Too Many Requests
```
4. ### 5XX - SERVER ERROR
```
500 Internal Server Error
502 Bad Gateway
503 Service Unavailable
504 Gateway Timeout
```
## HTTP vs HTTPS

1. **HTTP** - HTTP sends data in plain text and it can be interpreted or breached through network
2. **HTTPS** - HTTPS is HTTP + TLS
Basically, data is encrypted before it is sent 

**TLS GIVES**:
1. Encryption
2. Intergrity
3. Authentication 

## TLS HANDSHAKE
- We know that TCP handshake is to establish a connection
- But the data is not yet encrypted, so we need TLS to make that encryption.

### HIGH LEVEL TLS HANDSHAKE 

1. **STEP1** 
- client asks for a use of HTTPS in this step

2. **STEP2**
- Server sends its TLS certificate in which it proves it identity, like I am really google.com

3. **STEP3**
- client verifies the certificte

4. **STEP4**
- Client and server agree on encryption keys

5. **STEP5**
- encrypted communication begins

## CERTIFICATES
### WHAT DOES A CERTIFICATE CONTAIN?
- At a high level; it might contain 
1. Domain Name
2. Owner Information 
3. Validity Period
4. Public key
5. Issuer

## CURL

- Curl is basically a browser without UI

### Basic Request
`curl google.com`
- This performs a GET operation and sends the response

### VIEW HEADERS
```
curl -I https://google.com

output:

HTTP/2 200
content-type: text/html
```
### VERBOSE MODE
- Gives complete information about TLS, TCP HTTP and also HTTPS request
```
curl -v https://google.com

output:
DNS Lookup
TCP Connection
TLS Handshake
HTTP Request
HTTP Response
```

### POST REQUEST
```
curl -X POST https://api.example.com/users
or
curl -X POST \
-H "Content-Type: application/json" \
-d '{"name":"Puneeth"}' \
https://api.example.com/users
```

### Ignore Certification Validation 
```
curl -k https://example.com
```
