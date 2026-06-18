## NETWORK DAY-09

## WHAT IS DNS AND WHAT PROBLEM IT SOLVES?
- DNS stands for Doman Name System
- It is a distributed system that translates Domain Name to IP Addresses

## HOW DNS WORKS STEP BY STEP
```
Browser
   ↓
Local DNS Cache
   ↓
DNS Resolver
   ↓
Root DNS Server
   ↓
TLD Server (.com)
   ↓
Authoritative DNS Server
   ↓
IP Address Returned
   ↓
Browser Connects to Server
```

1. When we search for `www.google.com` in the browser, first it checks if the IP of this domain is cached in the browser. If not, we check the cache of the OS to see if the IP address of this domain is cached. If yes, we return else we reach out to the DNS Resolver
2. we ask `DNS Resolver` whose job is to resolve the domain name. Meaning, find the IP 
3. The Resolver asks `Root DNS Resolver` for the IP of the Domain name, Root DNS Resolver doesnt know the full IP, but it will identify the .com in the and says ask the `.com TLD Server`
4. Once the `Root DNS Resolver` asks `TLD`, TLD replies ask `Authoritative DNS Server`.
And it returns the IP and the website loads

## DNS HIERARCHY
```
                .
              (Root)
                 |
        -----------------
        |       |       |
      .com    .org    .net  (Domain)
        |
    google.com (sub-domain)
        |
   mail.google.com
   ```

1. ### Root Domain - `.`
- This is the top most in the hierarchy

2. ### TLD (Top-Level Domain)
- Example: `.com, .org`

3. ### Second Level Domaim
- Example: `mail.google.com`
- Here:
    - . is the root
    - .com is the TLD
    - google. is the domain
    - mail is the sub-domain

## FQDN DULLY QUALIFIED DOMAIN NAME
-  An FQDN contains sub-domain + domain + TLD + Root
- Example: mail.google.com.
- It is fully qualified because suppose we are inside a company network and we want to type mail, the system might assume the remaining part and hence it is not fully qualified in this scenario and hence is fully qualified  

## DNS RECORDS
- DNS servers doesnt just store IP they store DNS records

### A RECORD 
- Maps a domain to an IPv4 address
- example: maps myapp.com → 192.168.1.10 (IPv4 address)

### AAAA RECORD
- Maps a domain to an IPv6 address
- Example: myapp.com → 2001:db8::1

### CNAME RECORD
- maps one domain name to another domain name 
- example: www.myapp.com -> app.com

### MX RECORD
 - MX stands for Mail Exchange
 - without this the mail delivery will fail 
 - Example: when someone sends a email, mail servers checks MX records for that company and sends the mail
 - Example: support@company.com, the mail server checks for MX Record for company.com and delivers the mail

 ### TXT Records
 - It stores text information
 - This is for the user to prove the ownership of a webiste, it is similar to a bearer token.
 - Example: suppose google asks you to autheticate yourself as the owner then, you place the value in TXT record and authenticate yourself

 ## NS SERVERS
 - Basically it talks about who is responsible for this domain 
 - Example: google.com may have ns1.cloudflare.com as its authoritative name servers. so when someones wants an information they reach out to these servers


## DNS CACHING 
-  TTL is 36000 - 1 hour 
- concept remains same as you know 
- caching is generally stored in Local, OS and browser

## DevOps work
1. ### QUERY FOR MX RECORDS
- dig google.com MX

2. #### Which file contains DNS resolver configuration
- cat /etc/resolv.conf