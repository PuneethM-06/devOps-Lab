## DAY 10 - NETWORKING 

### WHY DO WE NEED PORTS?
- IP addresses are responsible for identifying the host.
- whereas ports help you identify where the data should flow in for that host.

- Example: IP address identify the host and then port says the data should go to ssh or postgres

```
ss --tlnp
```
- This shows all the ports that are listening

### WHAT IS TCP
- TCP = Transmission Control Protocol
- TCP is Responsible for:
    - **RELIABLE DELIVERY**: TCP detects missing packets and retransmits them to ensure reliable delivery.
    - **ORDERED DELIVERY**: When packets arrrive they dont arrive in order but are mismatched. TCP is responsible for the right ordering of the packet
    - **ERROR HANDLING** - TCP verifies data wasn't corrupted during transit

## TCP 3-WAY HANDSHAKE
- This is a protocol that is followed before there is a transmission of packets take place. TCP doesnt directly start sending packets.

1. **STEP1 - SYN**
- This is the inital step before data packets are transmitted and here SYN -  stands for synchronization.
- Here; **client** asks to establish a connection

2. **STEP2 - SYN ACK**
- This is the second step where the server responds to the request with an **ACK - Acknowdlegement** saying, Yes! we can have an connection.

3. **STEP3 - ACK**
- In this step, the client now sends a ACK saying. "Great! lets have an connection"

And then there will be a flow of data 

## UDP 
-  UDP stands for User Datagram protocol
- The reason we need UDP even though we have TCP is that because, sometimes speed matters more than perfection and TCP comes with a cost

- unlike TCP, UDP doesnt
    - check for packets
    - check for ordering 
    - check for missing packets
In simple terms, it sends the packet and says Good luck!

## CLIENT PORTS vs SERVER PORT (EPHEMERAL PORTS)
- We all know google listens on port 443, so if millions of users are connection to google how will google identify uniquely
- Here is where, client ports come into picture. 
-  The operating system automatically assigns a temporary port called an Ephemeral Port to the client side of a connection. called as **Ephemeral port** and google makes use of this to identify unique user 

## TCP CONNECTION STATES
```
when we run: ss -tan
```
This command is used to identify the TCP connection states

1. ### LISTEN
- Example: 0.0.0.0:22 LISTEN
- This means that port 22 is waiting for a connection to happen 

2. ### ESTABLISHED
- 192.168.1.10:55000 -> 10.0.0.5:22 ESTABLISHED
- TCP handshake completed and data is flowing 

3. ### TIME WAIT
- This tells connection has been closed but we will wait for sometime before we remove it completely
- This is to ensure that lost packets don't interfere the new connection 

4. ### CLOSE WAIT
- This often indicates a problem, it means that the remote side closed the connection, but your application hasnt cleaned it up.
- It is an indication of app bug, data leaks

## OVERALL FLOW
```
Browser
   ↓
DNS Lookup
   ↓
IP Address Returned
   ↓
Connect to Port 443
   ↓
Client Gets Ephemeral Port
   ↓
TCP 3-Way Handshake
   ↓
TLS Handshake
   ↓
HTTP Request
   ↓
HTTP Response
   ↓
Browser Renders Page
```

## COMMON PORTS
| Port | Service          |
| ---- | ---------------- |
| 22   | SSH              |
| 53   | DNS              |
| 80   | HTTP             |
| 443  | HTTPS            |
| 3306 | MySQL            |
| 5432 | PostgreSQL       |
| 6379 | Redis            |
| 8080 | Alternative HTTP |

## PORT RANGES
1. 0 - 1023 - **FAMOUS PORTS**
- 22   SSH
- 53   DNS
- 80   HTTP
- 443  HTTPS

1023 - 49151  **REGISTERED PORTS**
- 3306  MySQL
- 5432  PostgreSQL
- 6379  Redis
- 8080  HTTP Alternative

49152 - 65535 **EPHEMERAL PORTS**
