## DAY 12 - NETWORKING 

### WHAT IS A FIREWALL?
- A security system that monitors and control incoming and outgoing traffic based on operational rules defined

```
Internet
    |
    |
+-----------+
| Firewall  |
+-----------+
    |
    |
 Linux Server
 ```
- A firewall inspects
    - Source IP
    - Destination IP
    - Protocol
    - Destination Port
- Example: `ALLOW TCP 22 FROM 203.0.113.10` everything else is denied 

### INBOUND vs OUTBOUND TRAFFIC
- Inbound and outbound traffic are relative to the machine we're talking about 
```
Laptop: 192.168.1.10
Server: 10.0.0.5
Traffic: Laptop --------> Server
                 Port 22
```
From the servers perspective this is an INBOUND traffic 

- If server sends the response to ubuntu then it is OUTBOUND TRAFFIC

### WHAT INFORMATION DOES A PACKET HAVE?
- Source IP
- Destination IP
- Protocol
- Destination Ports

### PACKET FILTERING 
- It is the process where the firewalls examine the data packets to decide whether to allow or drop/refuse the packet based on certain fields 

### STATEFUL vs STATELESS FIREWALLS

- **SECURITY GROUP = STATEFUL**
- **NETWORK ACL = STATELESS**


- stateful firewall has memory and it keeps a stable table to identify the next time the request comes where as stateless firewalls do not have memory to remember about it 
- Stateful firewall can allow the response/return traffic automatically where as in stateless we have to define it else the response will be dropped 

```
STATEFUL

Laptop --------> EC2
      SSH Request

Laptop <-------- EC2
      SSH Response (Automatically Allowed)
```

```
STATELESS
Laptop --------> EC2
      SSH Request (Allowed)

Laptop <-------- EC2
      SSH Response (No outbound rule)

DROP ❌
```
### SECUIRTY GROUP
- SG's operate at the instance level
- It only has allow and there is not explicit DENY, everything else which is not allowed is deined automatically

### NETWORK ACL
- They operate at the subnet level 
- They have both ALLOW and DENY

