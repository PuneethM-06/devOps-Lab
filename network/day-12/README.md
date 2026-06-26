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

