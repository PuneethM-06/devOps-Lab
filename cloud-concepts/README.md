# Day 13 – Cloud Fundamentals ☁️

## Topics Covered

* What is Cloud Computing
* Why Cloud Computing became popular
* Traditional On-Prem Infrastructure
* Cloud vs On-Prem
* Cloud Trade-offs
* Pay-as-you-go Pricing Model
* Vendor Lock-in
* When to choose Cloud vs On-Prem

---

## What is Cloud Computing?

Cloud computing is the on-demand delivery of computing resources such as servers, storage, networking, and databases over the internet. The underlying infrastructure is managed by a Cloud Service Provider (CSP), and customers pay only for the resources they consume.

Examples of CSPs:

* AWS
* Microsoft Azure
* Google Cloud Platform (GCP)

---

## Why Cloud?

Traditional infrastructure requires:

* Buying servers
* Setting up networking
* Managing data centers
* Hardware maintenance
* Capacity planning

Cloud solves many of these challenges by providing infrastructure on demand.

---

## On-Prem vs Cloud

| Feature            | On-Prem       | Cloud                         |
| ------------------ | ------------- | ----------------------------- |
| Initial Cost       | High          | Low                           |
| Hardware Ownership | Company       | CSP                           |
| Maintenance        | Company       | CSP (Physical Infrastructure) |
| Scaling            | Slow          | Fast                          |
| Provisioning       | Days or Weeks | Minutes                       |
| Global Reach       | Difficult     | Easy                          |
| Pricing            | Fixed Cost    | Pay-as-you-go                 |

---

## Advantages of Cloud

* Pay only for what you use.
* Rapid provisioning of infrastructure.
* Easy scalability.
* Access to hundreds of managed services.
* Global infrastructure.

---

## Disadvantages of Cloud

* Vendor lock-in.
* Internet dependency.
* Compliance and data residency challenges.
* Costs can increase if resources are not managed properly.

---

## What is Pay-as-you-go?

You are charged only for the resources that you consume.

Example:

If an EC2 instance runs for 2 hours and is terminated, you pay only for:

* Compute usage
* Storage consumed
* Data transfer (if applicable)

You do **not** pay for the physical server or data center.

---

## What is Vendor Lock-in?

Vendor lock-in happens when an application becomes heavily dependent on one cloud provider's services, making migration to another provider difficult and expensive.

Example:

Application uses:

* AWS Lambda
* DynamoDB
* SQS
* CloudFormation

Migrating to Azure or GCP would require significant redesign and migration effort.

---

## When to Choose On-Prem?

* Strict compliance requirements
* Government workloads
* Highly sensitive data
* Legacy applications
* Full infrastructure control required

---

## When to Choose Cloud?

* Startups
* Rapidly growing applications
* Variable traffic patterns
* Global applications
* Faster time to market

---

## Key Takeaways

✅ Cloud is still someone else's infrastructure.

✅ Cloud is much more than renting a server.

✅ Cloud provides elasticity and on-demand resources.

✅ Cloud is not always cheaper than on-prem.

✅ Cloud and on-prem both have valid use cases.
