## DAY 6 OF LINUX

#### UNDERSTANDING HOW REQUEST REACHES AN APPLICATION 
```
Browser
↓
DNS
↓
IP Address
↓
TCP Connection
↓
Port
↓
Application
↓
HTTP Response
```
- suppose we hit a website, say `www.google.com` DNS helps in giving us the IP address of that. 
- Browser doesnt know where google.com lives and hence DNS helps us in identifying that 
- Once the IP address is given by the DNS, a TCP connection happens where certain protocols and security checks happen there.
- Example can be: Three-way handshake
- From there, from a certain port the request reaches the application
---
### PING IN LINUX

```
Example: ping google.com
or
ping -c google.com
```
- ping sends an ICMP request to a host and waits for ICMP replies from host
- If the host responds with 64 bytes from... then we are connected
-`ping -c 4 google.com` means ping with 4 packets and receive 4 packets

#### NOTE:
- this can be usefull suppose we send 10 packets and if you receieve 8 then it might be because of network connectivity or firewall etc.

#### nslookup
- it can be used to obtain IP address
---

### dig in LINUX
```
dig github.com
```
Example output:
```
dig github.com

; <<>> DiG 9.10.6 <<>> github.com
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 52142
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 1232
;; QUESTION SECTION:
;github.com.			IN	A

;; ANSWER SECTION:
github.com.		3	IN	A	20.207.73.82

;; Query time: 68 msec
;; SERVER: fe80::6298:49ff:feeb:4b31%15#53(fe80::6298:49ff:feeb:4b31%15)
;; WHEN: Thu Jun 11 12:21:39 IST 2026
;; MSG SIZE  rcvd: 55
```
> QUESTION SECTION shows whats being asked 
- Answer section answers the IP

```
;; ANSWER SECTION:
github.com.		3	IN	A	20.207.73.82
```
**TIME-TO-LIVE** - Here it is 3 seconds

**ANOTHER EXAMPLE WITH SHORTER VERSION OF ANSWER**
```
dig github.com +short 
20.207.73.82
```

### nslookup IN LINUX
Example
```
nslookup github.com
Server:		fe80::6298:49ff:feeb:4b31%15
Address:	fe80::6298:49ff:feeb:4b31%15#53

Non-authoritative answer:
Name:	github.com
Address: 20.207.73.82
```
## HTTP REQUESTS IN LINUX

`curl` - can be used as a browser in a terminal
- Example:
```
curl https://google.com
<HTML><HEAD><meta http-equiv="content-type" content="text/html;charset=utf-8">
<TITLE>301 Moved</TITLE></HEAD><BODY>
<H1>301 Moved</H1>
The document has moved
<A HREF="https://www.google.com/">here</A>.
</BODY></HTML>
```

### USING HEADERS ONLY
here can use `curl -I https://github.com`
- **This is used to check if the application is healthy**
- Example:
```
curl -I https://google.com
HTTP/2 301 
location: https://www.google.com/
content-type: text/html; charset=UTF-8
content-security-policy-report-only: object-src 'none';base-uri 'self';script-src 'nonce-ZTrplGzHcx3Ls1muKgmsqg' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp
date: Thu, 11 Jun 2026 06:59:47 GMT
expires: Sat, 11 Jul 2026 06:59:47 GMT
cache-control: public, max-age=2592000
server: gws
content-length: 220
x-xss-protection: 0
x-frame-options: SAMEORIGIN
alt-svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000
```
---
### NOTE: Here we can see the status code that helps us to understand the problem
----
 ### verbose mode
 - this is one of the important trouble shooting commands 
 - Here we can see: 
    - DNS resolution
    - TCP connection
    - Request headers
    - Response headers

## LISTENING TO PORTS
```
ss -tlnp
```
`-t` - TCP sockets
`-l` - Listening sockets
`n` - show numeric ports
`p` - show process names


### command for listing all the ports
```
lsof -iTCP -sTCP:LISTEN -n -P
```
This shows all the programs that are waiting for network connections

### command to check if a port is listening 
```
lsof -iTCP:8200 -sTCP:LISTEN
or lsof -i :8000
```

### Local Testing 
```
python3 -m http.server:8200
```

### DOWNLOADING FILES
```
wget https://example.com/file.zip
```
