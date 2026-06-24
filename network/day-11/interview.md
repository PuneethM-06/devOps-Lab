# Day 11 - HTTP, HTTPS, TLS & Status Codes Interview Questions

## HTTP Fundamentals

1. What is HTTP and why is it required?

2. What does it mean when we say HTTP is a stateless protocol?

3. Explain the Request → Response model used by HTTP.

4. What is the difference between TCP and HTTP?

---

## HTTP Requests & Responses

5. What are the main components of an HTTP request?

6. What are the main components of an HTTP response?

7. What is the purpose of HTTP headers?

8. Why do POST requests usually contain a request body while GET requests typically do not?

---

## HTTP Methods

9. Explain the purpose of the following HTTP methods:

   * GET
   * POST
   * PUT
   * PATCH
   * DELETE

10. What is the difference between PUT and PATCH?

11. Which HTTP method would you use to create a new resource and why?

---

## Status Codes

12. What are the five major categories of HTTP status codes?

13. What is the difference between:

    * 401 Unauthorized
    * 403 Forbidden

14. When would a server return a 404 Not Found response?

15. What does a 500 Internal Server Error indicate?

16. What is a 502 Bad Gateway error and where is it commonly seen?

17. What is the difference between 502 Bad Gateway and 503 Service Unavailable?

18. What does a 429 Too Many Requests response indicate?

---

## HTTPS & TLS

19. What is the difference between HTTP and HTTPS?

20. Why is HTTPS considered more secure than HTTP?

21. What is TLS and why is it required?

22. At a high level, explain the TLS handshake process.

23. Does the TLS handshake happen before or after the TCP three-way handshake?

---

## TLS Certificates

24. What is a TLS certificate?

25. What information is typically stored inside a TLS certificate?

26. Why does a browser show "Your connection is not private"?

27. What is the most common certificate-related production issue?

---

## curl

28. What is curl and why is it commonly used by DevOps engineers?

29. How would you view only the response headers of a website using curl?

30. How would you enable verbose output in curl?

31. What does the `-k` option do in curl?

32. How would you send a POST request using curl?

---

## End-to-End Flow

33. Explain what happens when a user enters:

https://google.com

in a browser until the webpage is displayed.

34. Explain the role of each of the following in a web request:

    * DNS
    * TCP
    * TLS
    * HTTP

35. Describe the complete flow:

Browser → DNS Lookup → IP Address → Port 443 → TCP Handshake → TLS Handshake → HTTP Request → HTTP Response → Browser Render
