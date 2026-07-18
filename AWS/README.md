# AWS - CLOUD PRACTITIONER

### WHAT IS A SERVER COMPOSED OF?
1. CPU
2. RAM -Memory
3. Storage - Database
4. Networking Aspects - Routers, DNS etc.

### WORKFLOW
- Client will send data through a Router using Internet. Client and Server have their own IP addresses
- Router is responsible for routing the data packet to a Switch.
- Switch is responsible for sending the data packet to the right server 

### TRADITIONAL IT APPROACH / DATA CENTRES DISADVANTAGES
1. Pay rent for data centre
2. Power supply, cooling was a concern
3. Adding and replacing hardware - MAINTAINENCE
4. Scaling is LIMITED

## WHAT IS CLOUD COMPUTING?
- It is the **on-demand delivery** of compute power, database storage, applications and other IT resources.
- **pay-as-you-go pricing**
- Access almost these resources **instantly**

### NOTE: WE NEVER PAY FOR THE INFRASTRUCTURE, WE PAY FOR THE RESOURCES USED.

## DEPLOYMENT MODELS IN CLOUD
1. ### PRIVATE CLOUD
- Private infrastructure
- Cloud servivce is not exposed to public 
- Complete control
- More security 

2. ### PUBLIC CLOUD
- AWS, GCP and Azure
- Cloud services are owned and operated by third-party
- They are delivered over internet

3. ### HYBRID CLOUD 
1. Private + Public cloud
2. Few services On-prem while some are on Cloud

## FIVE CHARACTERISTICS OF CLOUD COMPUTING A
1. Completely **ON-DEMAND**
2. Broad Network Access
3. Multi-tenancy and resource pooling  - Multiple customers can use the same services having complete privacy
4. Rapid Elasticity and Scalability - Quickly and easily scalable 
5. Measured services

## SIX ADVANTAGES OF CLOUD
1. Tade **CAPITAL EXPENSES FOR OPERATIONAL EXPENSES, CAPEX FOR OPEX** 
2. Benefit of scaling 
3. stop guessing capaciy 
4. Increased speed and agility
5. stop spending money and running data centres

## DIFFERENT TYPES OF CLOUD COMPUTING 

1. ### INFRASTRUCTURE AS A SERVICE - IaaS
- Provides **networking, computers, data storage space**
- Highest level of flexibility
- Provide building blocks for cloud IT.
- EC2

2. ### PLATFORM AS A SERVICE - PaaS
- Remove the need to manage the underlying infrastructure
- Manage on deployment and managing applications
- Elastic Beanstalk

3. ### SOFTWARE AS A SERVICE - SaaS
- Completed product that is run and managed by service providers
- Gmail

### ON-PREM WE MANAGE:
1. Applications 
2. Data
3. Runtime
4. Middleware
5. O/S
6. Virtulization
7. Servers
8. Storage
9. Networking 

### IaaS
1. Applications
2. Data
3. Runtime
4. Middleware
5. O/S is maintained by us, rest is done by CSP

### PaaS
1. Application
2. Data
- Is managed by rest is done by CSP


## PRICING ON CLOUD
- It is mainly pay-as-you-go but it depends 
1. COMPUTE - Pay for compute time 
2. Storage - Data stored
3. Network - Pay when Data transfer is outside of cloud

## AWS REGIONS
- Regions have name. Example:  us-east-1
- It is going to be a cluster of Data centres
- **MOST SERVICES ARE REGION SPECIFIC**

### HOW TO CHOOSE A AWS REGION 
1. Compliance - Sometimes Govt needs your data to be local 
2. Latency - Close to users
3. Available services - Not all regions have all services
4. Pricing

## AVAILABILITY ZONES
- Minimum is 3 and Maximum is 6
- It is one or more discrete Data centres
- They're seperated from eachother, to avoid disasters

## AWS POINTS OF PRESENCE OR EDGE LOCATIONS
- Have more than 400 
- It is used for CDN for low latency 

# IAM SECTION

- IAM - IDENTITY AND ACCESS MANAGEMENT 
- It is a global service
- **USERS** - People within in orgaization 
- **GROUP** - People can be grouped

### NOTE: Group contains users, and a single user can belong to multiple groups
- These groups and users will be assigned permissions to ensure what services can be accessed and also to restrict what can be accessed.

- **LEAST PREVILIGE PERMISSION** - Not to give more permissions than needed

## IAM POLICIES INHENRITANCE
- Attaching policy at a group level applies policy to everyone in the group 
- **INLINE POLICY** used for applying policy for a user not in any group 
- IAM policy consists of:
    - Version 
    - ID
    - Statement
    - SID - Statement ID
    - Affect - allow or deny
    - Principal - Account, user or role 
    - Action - List of API calls
    - Resources - List of resources to which the action will be applied 

### IAM PASSWORD POLICY
- Stronger password = More security
- We can set a customized password security:
    - Number of characters
    - character case
    - non-alphanumeric characters
    - Allow IAM users to change their own pwd

### MULTI-FACTOR AUTHENTICATION 
- MFA is a combination of **password we know + secruity device we own**

### MFA OPTIONS
1. Virtual MFA device - Google authenticator, Authy
2. Universal 2nd Factor - U2F security device - Yubikey is an example 

### HOW CAN USERS ACCESS AWS
1. AWS CONSOLE - Protected by MFA and password
2. CLI - command line interface
3. SDK - AWS Software development kit - Code

- CLI and SDK are protected by **Access Keys**

- Similar to terminal in CLI, we can make use of Cloudshell. But it is important to note that Cloudshell is not available for all regions 

## IAM ROLES FOR SERVICES
- Some AWS services will need permissions to perform certain operations 
- So we make use of **IAM ROLES**

## IAM SECURITY TOOLS
- **IAM Credentials report** - a report that lists all your account users and the status of their credentials 

- **IAM ACCESS ADVISOR** - Shows service permissions granted to a user and when were they last accessed

## IAM BEST PRACTICES
- Principle of least privelige
- Assign users to groups
- Create a strong pwd
- Create Roles fore giving permissions to AWS services


## SHARED RESPONSIBILITY MODEL
AWS is responsible for:
    - Infrastructure and security
    - Config and vulnerability analysis
We are responsible for:
    - Creating roles
    - enabling MFA
    - Assigning permissions
    - Rotate keys often

# EC2 INSTANCES 

- EC2 instances are capable of:
    1. Renting VM's
    2. Storing data 
    3. Distribute load across

- Operating system offered: Linux, Windows and Mac OS
- CPU
- RAM
- Network attached - EFS & EBS
-  Networking and Security group

- **BOOTSTRAPPING** - means launching commands when a machine starts 
- It runs only run once at the instance start 
- It does installing updates, softwares etc.
- EC2 runs as a root user

### NAMING CONVENTION
- m5.2xlarge

- m - instance class
- 5 - Generation
- 2xlarge - Size

## EC2 INSTANCE TYPES

1. ### GENERAL PURPOSE 
-  Great for webservers and code repo
- Balance between:
    - compute
    - memory
    - network

2. ### COMPUTE OPTIMIZED
- Great for compute-intensive tasks
- Requires high performance processors
- Example: Batch processing workloads, dedicated gaming servers, machine learning 
- They are of name C

3. ### MEMORY OPTIMIZED
- Fast performance for workloads that process large datasets
- They are used for :
    - High performance, relational and non relational databases
    - In-memory databases optimized for BI

4. ### STORAGE OPTIMIZED
- Great for storag-intensive tasks, for reading and writing large data sets 
- Uses cases:
    - OLTP
    - Relational & NoSQL databases 
    - Cache for in-memory databases

## INTRODUCTION TO SECURITY GROUPS
- They will control how traffic will move IN & OUT of the Ec2 instances
- They have rules that refer by IP address of by Security group 
- They regulate to:
    - Access to ports
    - Validate IP range 
    - control inbound and outbound network

### NOTE: SECURITY GROUP BY DEFAULT WILL ALLOW OUTBOUND RULE 

- Security groups can be attached to multiple instances
- Locked down to a region or a VPC combination 
- It is a good practise to maintain a seperate security group for SSH

### CLASSIC PORTS TO KNOW 
1. SSH - Port 22
2. FTP - Port 21
3. SFTP - Port 22
4. HTTPS - Port 443
5. 3389 - RDP
6. 80 - HTTP

### SSH SUMMARY TABLE 
- For linux servers we can use this
- We connect **insiders of our servers** to perform maintaience
- SSH can be used for Mac, Linux and then windows >=10
- We can use Putty for Windows < 10
- EC2 instance connect - Available for all( Mac, windows and linux)
