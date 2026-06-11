# Day 06 - Networking Fundamentals Interview Questions

## Networking Basics

1. Explain how a request reaches an application.
2. What is DNS?
3. Why do we need DNS?
4. What happens when you enter `google.com` in a browser?
5. What is an IP address?
6. Why do computers communicate using IP addresses instead of domain names?
7. What is the difference between a domain name and an IP address?
8. What is a TCP connection?
9. Why is TCP commonly used for web traffic?
10. What is a port?
11. Why are ports needed if we already have IP addresses?
12. What is the purpose of port 80?
13. What is the purpose of port 443?

## DNS

14. What does the `dig` command do?
15. What does the `nslookup` command do?
16. What is the difference between `dig` and `nslookup`?
17. What information can you obtain from `dig`?
18. What is the purpose of the QUESTION SECTION in a dig response?
19. What is the purpose of the ANSWER SECTION in a dig response?
20. What does TTL stand for?
21. What is the purpose of TTL in DNS?
22. What does the following command do?

```bash id="n7zv4q"
dig github.com +short
```

23. What does the `+short` option do?
24. What does an `A` record represent?
25. What does it mean if DNS resolution fails?

## ping and Connectivity

26. What does the `ping` command do?
27. Which protocol does ping use?
28. What is ICMP?
29. What does the following command do?

```bash id="v3j8ls"
ping -c 4 google.com
```

30. What does packet loss mean?
31. What could cause packet loss?
32. What does latency mean in ping output?
33. What does the response below indicate?

```text id="j2pm1x"
64 bytes from ...
```

34. If ping fails, does that always mean the website is down?
35. Why do many production servers block ICMP traffic?

## HTTP and curl

36. What is HTTP?
37. What is HTTPS?
38. What is the difference between HTTP and HTTPS?
39. What does the `curl` command do?
40. Why is curl commonly used by DevOps engineers?
41. What happens when you run:

```bash id="2qfivq"
curl https://google.com
```

42. What does `curl -I` do?
43. Why would you use `curl -I` instead of `curl`?
44. What does `curl -v` do?
45. Why is verbose mode useful during troubleshooting?
46. What information can you see in verbose mode?
47. What is an HTTP request?
48. What is an HTTP response?

## HTTP Status Codes

49. What does HTTP status code 200 mean?
50. What does HTTP status code 301 mean?
51. What does HTTP status code 302 mean?
52. What does HTTP status code 404 mean?
53. What does HTTP status code 500 mean?
54. What does HTTP status code 503 mean?
55. What is the difference between 404 and 500?
56. What is the difference between 500 and 503?

## TCP and Connections

57. What is the TCP three-way handshake?
58. Explain SYN, SYN-ACK, and ACK.
59. Why is a TCP handshake necessary?
60. What happens before an HTTP request is sent?

## Listening Ports

61. What does it mean for a process to be listening on a port?
62. Why must an application listen on a port?
63. What does the following command do?

```bash id="0frygw"
ss -tlnp
```

64. What does the `-t` option mean?
65. What does the `-l` option mean?
66. What does the `-n` option mean?
67. What does the `-p` option mean?
68. How would you identify the process listening on port 8080?

## macOS Port Troubleshooting

69. What does the following command do?

```bash id="9ey4it"
lsof -iTCP -sTCP:LISTEN -n -P
```

70. How would you check if port 8000 is listening?

```bash id="rld85v"
lsof -iTCP:8000 -sTCP:LISTEN
```

71. What information does lsof provide?
72. What does LISTEN mean in the output?

## Local Testing

73. What does the following command do?

```bash id="2kt2gj"
python3 -m http.server 8000
```

74. Why is this command useful for learning networking?
75. How would you verify that the server is listening?
76. How would you test the server locally?

```bash id="7e3i6r"
curl http://localhost:8000
```

77. What should happen if the server is running correctly?

## localhost

78. What is localhost?
79. What IP address does localhost resolve to?
80. What does `127.0.0.1` represent?
81. Why is localhost important for application testing?
82. Are `localhost` and `127.0.0.1` the same thing?
83. What happens when you run:

```bash id="o8mcp8"
ping localhost
```

## Downloading Files

84. What does the `wget` command do?
85. What is the difference between `wget` and `curl`?
86. When would you use wget instead of curl?
87. What does the following command do?

```bash id="r58ifz"
wget https://example.com/file.zip
```

## Troubleshooting Scenarios

88. A user reports that a website is down. What is the first thing you would check?

89. DNS resolution is failing. Which commands would you use?

90. A website is returning HTTP 503. What does that indicate?

91. A domain resolves correctly, but curl cannot connect. What would you investigate next?

92. An application is running but users cannot access it. Which networking checks would you perform?

93. How would you verify that an application is listening on the correct port?

94. How would you determine whether a problem is DNS-related or application-related?

95. A server responds to ping but curl returns an error. What does that tell you?

## Frequently Asked DevOps Questions

96. Explain the complete flow from browser to application.

97. What happens when you type a URL into a browser?

98. What is the difference between DNS, IP addresses, ports, and applications?

99. How would you troubleshoot a website that is inaccessible?

100. Explain the role of DNS, TCP, ports, and HTTP in a web request.
