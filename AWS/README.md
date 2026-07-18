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

