# Day 9 - DNS Interview Questions

## DNS Fundamentals

### Q1. What is DNS and why is it required?

### Q2. Explain how a domain name gets translated into an IP address.

### Q3. What is the difference between a public IP address and a domain name?

---

## DNS Resolution Flow

### Q4. Explain the complete DNS resolution flow when a user opens `google.com`.

Expected flow:

Browser Cache → OS Cache → DNS Resolver → Root DNS Server → TLD DNS Server → Authoritative DNS Server → IP Address

### Q5. What is the role of a DNS Resolver?

### Q6. Does the Root DNS Server know the actual IP address of a domain?

### Q7. Which DNS server contains the actual DNS records for a domain?

---

## DNS Hierarchy & FQDN

### Q8. What is a TLD? Give some examples.

### Q9. What is an FQDN (Fully Qualified Domain Name)?

### Q10. Break down `api.dev.company.com.` into:

* Root
* TLD
* Domain
* Subdomains
* Hostname

---

## DNS Records

### Q11. Explain the following DNS records:

* A
* AAAA
* CNAME
* MX
* TXT
* NS

### Q12. What is the difference between an A record and a CNAME record?

### Q13. How does email delivery work using MX records?

---

## DNS Caching & TTL

### Q14. What is DNS caching and what is TTL?

### Q15. A DNS record was changed from `1.1.1.1` to `2.2.2.2`, but some users still reach the old IP. Why does this happen?

---

## Troubleshooting Scenario (Must Know)

### Scenario

`ping 8.8.8.8` works but `ping google.com` fails.

What would you check and how would you troubleshoot the issue?

