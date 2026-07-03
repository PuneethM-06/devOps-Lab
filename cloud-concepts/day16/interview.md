# Day 16 – Scalability, High Availability & Fault Tolerance: Interview Questions

## Fundamentals

1. Explain the difference between scalability, high availability, and fault tolerance with examples.
2. Can a system be scalable but not highly available? Give an example.
3. Can a system be highly available but not fault tolerant? Explain.
4. What is a Single Point of Failure (SPOF)? How do you identify one in an architecture?
5. What is the difference between vertical scaling and horizontal scaling?
6. Why is horizontal scaling generally preferred in cloud-native applications?
7. What are the limitations of vertical scaling?
8. Does Auto Scaling guarantee high availability? Why or why not?
9. What metrics would you monitor to decide when to scale an application?
10. When would you choose vertical scaling over horizontal scaling?

## Scenario-Based Questions

11. Your application runs on a single EC2 instance and CPU usage reaches 95% every evening. How would you solve this?
12. Your application servers are deployed across multiple AZs, but your database is hosted in a single AZ. Is the application highly available? Why?
13. A Load Balancer has only one backend server attached to it. Is the architecture highly available?
14. How would you design a highly available web application on AWS?
15. An entire AWS Availability Zone goes down. How should your application respond?
16. An entire AWS Region fails. How would you design for disaster recovery?
17. How would you make a database fault tolerant?
18. How would you perform a zero-downtime deployment?
19. Explain active-active and active-passive architectures and when you would use each.
20. How does Kubernetes improve application availability?

---

# Day 17 – Shared Responsibility Model: Interview Questions

## Fundamentals

1. What is the Shared Responsibility Model?
2. What is the difference between "Security of the Cloud" and "Security in the Cloud"?
3. Why does the Shared Responsibility Model exist?
4. Who is responsible for patching the operating system on an EC2 instance?
5. Who is responsible for patching the operating system in AWS Lambda?
6. Who is responsible for securing IAM users and permissions?
7. Who is responsible for encrypting application data?
8. Who is responsible if an S3 bucket is accidentally made public?
9. Why is IAM considered one of the most important customer responsibilities?
10. Why do most cloud breaches happen despite secure cloud providers?

## Scenario-Based Questions

11. Your EC2 instance was compromised because SSH was open to the internet. Who is responsible?
12. A power outage occurs in an AWS data center. Who is responsible?
13. An engineer accidentally commits database credentials to GitHub. Who is responsible?
14. An application contains an SQL Injection vulnerability. Who is responsible?
15. A hypervisor vulnerability allows a VM escape. Who is responsible?
16. Your RDS instance was accidentally deleted by an employee. Who is responsible?
17. You lose access to your KMS encryption key. Who is responsible?
18. How do responsibilities change between EC2, EKS, and Lambda?
19. Why do organizations still need DevSecOps if cloud providers already secure their infrastructure?
20. What operational responsibilities are transferred to AWS when moving from EC2 to Lambda?
