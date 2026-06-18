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

