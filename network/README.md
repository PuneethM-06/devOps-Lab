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