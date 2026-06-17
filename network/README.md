## DAY-08 NETWORKING 

### WHAT IS AN IP ADDRESS?
- IP is Internet Protocol adderss
- IP address is given to every device that is connected over internet so that other devices know where to send the data
- Example: 192.168.1.37

### IPv4 STRUCTURE
- Example: Lets take an IP and divide it into 4 parts 192.168.1.10 -> 192| 168| 1 | 10
- each section is called as `octect`
- It is called as octect because each section contains 8 bits. Hence, 8*4 = 32
- IP address are actually binary numbers underneath it 

### PUBLIC AND PRIVATE IP's

#### Private IP:
- Private IP are addresses that someone cannot connect directly from the internet
- Private IP range from 

| Private IP Range      | Start Address | End Address     | 
| --------------------- | ------------- | --------------- |
| Class A Private Range | 10.0.0.0      | 10.255.255.255  | 
| Class B Private Range | 172.16.0.0    | 172.31.255.255  | 
| Class C Private Range | 192.168.0.0   | 192.168.255.255 |

#### Public IPs
- Public IPs are reachable through internet 

## NETWORK ID AND HOST ID
- Example:  192.168.1.0/24
- so we need to think like the first 24 bits identify the host and remaining is the netwok 
- In the above example: 192.168.1.0 is the network ID and 192.168.1.24 is host and together we write 192.168.1.24

| Address                     | Meaning           |
| --------------------------- | ----------------- |
| 192.168.1.0                 | Network Address   |
| 192.168.1.1 - 192.168.1.254 | Usable Hosts      |
| 192.168.1.255               | Broadcast Address |

## SUBNET MASK
- subnet mask acts as a marker to seperate host address and network address
- Example:
| CIDR | Subnet Mask     |
| ---- | --------------- |
| /8   | 255.0.0.0       |
| /16  | 255.255.0.0     |
| /24  | 255.255.255.0   |
| /32  | 255.255.255.255 |
- For /16, The first two octect belongs to network address and the rest two belongs to device address
 
 ## NOTE: IMPORTANT QUESTION
1.  > 192.168.1.0/24; How many devices can exist in this subnet?
 -  Answer: 2^(host bits) - 2^0 = 2

 2. > `/25` How many host bits are there?
 - In an IPv4 total bits is 32 so; 32-25 = 7
 > How many total address?
 - 2^(Bits) = 2^7 = 128
 > Usable hosts
 - 128-2 = 126 usable hosts

 ## FORMULA FOR ANY SUBNET
1. Host bits = 32 - CIDR
2. Total address = 2^(Host bits)
3. Usable hosts = Total address - 2


## BROADCAST ADDRESS
- consider the example 10.20.30.40/24
- Here; network address is 10.20.30 right and then the host portion is 40
- For broadcast to make the largest network address possible it becomes 10.20.30.255/24
Here 255 is the key

1. > Question 192.168.10.50/24 what is network and broadcast address?
- Answer: 
    - Network address - 192.168.10.0
    - Broadcast address - 192.168.10.255

## NOTE:
/24 = One big network
/25 = Split it into 2 smaller networks
/26 = Split it into 4 smaller networks

- Example:
> Question: 192.168.1.0/26
- Answe:
    - 192.168.1.0 - 192.168.1.63 
    - 192.168.1.64 - 192.168.1.127 
    - 192.168.1.128 - 192.168.1.191 
    - 192.168.1.192 - 192.168.1.255

### NOTE:
First usable host = Network Address + 1
Last usable host = Broadcast Address - 1