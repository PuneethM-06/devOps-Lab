# Day 8 – Networking Fundamentals (IP Addressing, CIDR & Subnetting)

## 1. What is an IP Address?

An IP Address (Internet Protocol Address) is a unique logical address assigned to a device on a network. It helps devices identify each other and communicate over a network.

Example:

```text
192.168.1.10
```

---

## 2. What is the difference between Public and Private IP Addresses?

### Public IP

* Accessible over the internet.
* Globally routable.
* Example: 8.8.8.8

### Private IP

* Used inside private networks.
* Not directly accessible from the internet.

Private IP Ranges:

```text
10.0.0.0     - 10.255.255.255
172.16.0.0   - 172.31.255.255
192.168.0.0  - 192.168.255.255
```

---

## 3. How many bits are there in an IPv4 Address?

* IPv4 Address = 32 bits
* 1 Octet = 8 bits
* 4 Octets = 32 bits

Example:

```text
192.168.1.10
```

---

## 4. Why can each octet contain values only from 0–255?

Because:

```text
2^8 = 256
```

Possible values:

```text
0 - 255
```

---

## 5. What is CIDR Notation?

CIDR (Classless Inter-Domain Routing) represents how many bits belong to the network portion.

Examples:

```text
/8
/16
/24
```

Example:

```text
192.168.1.45/24
```

means:

* 24 bits = Network
* 8 bits = Host

---

## 6. What is a Subnet Mask?

A subnet mask separates the network portion from the host portion.

Common Subnet Masks:

| CIDR | Subnet Mask     |
| ---- | --------------- |
| /8   | 255.0.0.0       |
| /16  | 255.255.0.0     |
| /24  | 255.255.255.0   |
| /32  | 255.255.255.255 |

---

## 7. What is the Network Portion and Host Portion?

Example:

```text
10.20.30.40/16
```

Network Portion:

```text
10.20
```

Host Portion:

```text
30.40
```

---

## 8. What is a Network Address?

A Network Address is obtained by setting all host bits to 0.

Example:

```text
10.20.30.40/16
```

Network Address:

```text
10.20.0.0
```

---

## 9. What is a Broadcast Address?

A Broadcast Address is obtained by setting all host bits to 1.

Example:

```text
192.168.1.45/24
```

Broadcast Address:

```text
192.168.1.255
```

---

## 10. What is the Usable Host Range?

Example:

```text
192.168.1.45/24
```

Network Address:

```text
192.168.1.0
```

Broadcast Address:

```text
192.168.1.255
```

Usable Hosts:

```text
192.168.1.1 - 192.168.1.254
```

---

## 11. How do you calculate the number of hosts in a subnet?

### Step 1

```text
Host Bits = 32 - CIDR
```

### Step 2

```text
Total Addresses = 2^(Host Bits)
```

### Step 3

```text
Usable Hosts = Total Addresses - 2
```

---

## 12. Why do we subtract 2?

Because two addresses are reserved:

1. Network Address
2. Broadcast Address

These cannot be assigned to devices.

---

## 13. Common CIDR Calculations

| CIDR | Host Bits | Total Addresses | Usable Hosts |
| ---- | --------- | --------------- | ------------ |
| /24  | 8         | 256             | 254          |
| /25  | 7         | 128             | 126          |
| /26  | 6         | 64              | 62           |
| /27  | 5         | 32              | 30           |
| /28  | 4         | 16              | 14           |

---

## 14. How does subnet splitting work?

Example:

```text
192.168.1.0/24
```

Split into four /26 subnets:

```text
192.168.1.0   - 192.168.1.63
192.168.1.64  - 192.168.1.127
192.168.1.128 - 192.168.1.191
192.168.1.192 - 192.168.1.255
```

---

# Important Interview Questions

### Q1. What is an IP Address?

### Q2. What is the difference between Public and Private IPs?

### Q3. What are the three Private IP ranges?

### Q4. How many bits are there in an IPv4 address?

### Q5. Why is an octet limited to 0–255?

### Q6. What is CIDR notation?

### Q7. What does /24 mean?

### Q8. What is a subnet mask?

### Q9. What is the subnet mask of /16 and /24?

### Q10. What is the difference between Network Portion and Network Address?

### Q11. What is the Network Address of 192.168.1.45/24?

### Q12. What is the Broadcast Address of 192.168.1.45/24?

### Q13. What is the usable host range of 192.168.1.45/24?

### Q14. How many usable hosts are available in a /27 subnet?

### Q15. Why do we subtract 2 while calculating usable hosts?

