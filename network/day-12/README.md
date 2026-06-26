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

### LINUX FIREWALLS
- There are three names that we will hear:
1. iptables
2. ufw
3. firewalld

- Here linux has a built in framework what we call as a Netfilter, which is used by iptables to configure our firewalls 

| Scenario                               | Chain     |
| -------------------------------------- | --------- |
| `ssh ubuntu@server`                    | ✅ INPUT   |
| `curl https://google.com`              | ✅ OUTPUT  |
| Linux machine acting as a router       | ✅ FORWARD |
| Server calling an external payment API | ✅ OUTPUT  |

INPUT - MACHINE TO SERVER
OUTPUT - SERVER TO MACHINE 
FORWARD - MACHINE TO VPN ROUTER TO SERVER

- iptables example
```
sudo iptables -A INPUT -p tcp --dport 22 -j ACCEPT
sudo iptables -A INPUT -p tcp --dport 443 -j ACCEPT
sudo iptables -P INPUT DROP
```
- This means that port 22 and 443 are accepted and everything else can be dropped 

## TROUBLE SHOOTING ORDER
1. Is the application running?
2. Is it listening on port 443?
3. Can I connect locally?
4. Check OS firewall (UFW/iptables).
5. Check Security Group.
6. Check Network ACL.
7. Check DNS and routing.
