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
