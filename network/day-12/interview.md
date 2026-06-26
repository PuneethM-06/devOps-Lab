# Day 12 - Important Interview Questions

## Firewall Fundamentals

1. What is a firewall and why do we need it?

2. What is the difference between inbound and outbound traffic?

3. What is the Principle of Least Privilege?

4. What is the difference between DROP and REJECT?

5. Why should SSH (22) not be exposed to `0.0.0.0/0`?

---

# Stateful vs Stateless

6. What is a stateful firewall?

7. What is a stateless firewall?

8. Why are Security Groups stateful?

9. Why are Network ACLs stateless?

10. Why does SSH fail if a Network ACL only allows inbound port 22 but has no outbound rule?

11. Why does HTTPS work with Security Groups without explicitly allowing return traffic?

---

# AWS Security Groups & Network ACLs

12. What is a Security Group?

13. What is a Network ACL?

14. Difference between Security Groups and Network ACLs.

15. Which one operates at the instance level?

16. Which one operates at the subnet level?

17. Which one supports explicit DENY rules?

18. Can an EC2 instance have multiple Security Groups attached?

19. Can a subnet have multiple Network ACLs attached?

20. Which is evaluated first: Network ACL or Security Group?

21. What happens if the Security Group allows traffic but the Network ACL denies it?

22. Explain the path of a packet from the internet to an EC2 instance.

---

# Linux Firewalls

23. What is Netfilter?

24. What is iptables?

25. What is UFW?

26. What is the relationship between UFW and iptables?

27. What are the INPUT, OUTPUT and FORWARD chains?

28. Give a real-world example of each chain.

29. What does this command do?

```bash
iptables -A INPUT -p tcp --dport 22 -j ACCEPT
```

30. What does:

```bash
iptables -P INPUT DROP
```

mean?

31. Why is setting `INPUT` policy to `DROP` considered a security best practice?

32. How do you check listening ports on Linux?

33. How do you check firewall rules?

---

# Troubleshooting (VERY IMPORTANT)

34. A website hosted on EC2 is not opening. How would you troubleshoot it?

35. Nginx is inactive (`dead`). What is your next step?

36. Nginx is running on port 8080 while users are hitting port 443. What is the issue?

37. Security Group allows 443, but UFW denies 443. Will the website work? Why?

38. If `curl localhost` works but the website is inaccessible externally, where is the problem likely to be?

39. If `curl localhost` fails, where is the problem likely to be?

40. Why is `curl localhost` an important troubleshooting step?

41. Can Security Groups and UFW both block the same request?

42. What would you check if everything looks correct but the website still doesn't open?

43. What logs would you inspect when debugging Nginx issues?

44. What are common causes of a `502 Bad Gateway` error?

45. Explain your troubleshooting methodology for connectivity issues.
