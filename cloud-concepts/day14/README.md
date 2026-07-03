# Day 14 – IaaS vs PaaS vs SaaS ☁️

## Infrastructure as a Service (IaaS)

The provider says:

> "Here is your infrastructure. Build whatever you want on top of it."

### You rent:

* Virtual Machines (VMs)
* Networking
* Storage
* Firewalls
* Load Balancers

### Cloud Provider Manages:

* Data centers
* Physical servers
* Storage hardware
* Networking infrastructure
* Virtualization layer

### Customer Manages:

* Operating System
* Patching
* Security updates
* Users and permissions
* SSH keys
* Applications
* Databases
* Monitoring

### Examples

* AWS EC2
* Azure Virtual Machines
* Google Compute Engine

### When to use IaaS?

* Need full control over infrastructure.
* Custom operating systems or runtimes.
* Legacy applications.
* Special networking requirements.

---

## Platform as a Service (PaaS)

The provider says:

> "Forget the servers. Give me your code and I'll run it for you."

### Provider Manages:

* Infrastructure
* Operating System
* Runtime
* Scaling
* Availability
* Patching

### Customer Manages:

* Application code
* Business logic
* Data

### Examples

* Heroku
* Google App Engine
* Azure App Service

### Advantages

* Faster development.
* Less operational overhead.
* Automatic scaling and patching.

### Disadvantages

* Less infrastructure control.
* Platform limitations.
* Higher risk of vendor lock-in.

---

## Software as a Service (SaaS)

The provider says:

> "Just use the software."

### Customer Manages:

* Users
* Configurations
* Data entered into the application

### Provider Manages:

* Everything else

### Examples

* Outlook
* Gmail
* Slack
* Microsoft Teams
* Zoom

### Advantages

* No infrastructure management.
* Quick adoption.
* Minimal operational overhead.

### Disadvantages

* Limited customization.
* Provider dependency.
* Data portability concerns.

---

# One-line Summary

IaaS = Manage Servers.

PaaS = Manage Code.

SaaS = Use Software.
