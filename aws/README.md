# AWS - CLOUD PRACTITIONER PAGE 1

### WHAT IS A SERVER COMPOSED OF?
1. CPU
2. RAM-Memory
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
1. **Compliance** - Sometimes Govt needs your data to be local 
2. **Latency** - Close to users
3. **Available services** - Not all regions have all services
4. **Pricing**

## AVAILABILITY ZONES
- Minimum is **3 and Maximum is 6**
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
2. Universal 2nd Factor - U2F security device - **Yubikey** is an example 

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
-  **Great for webservers and code repo**
- Balance between:
    - compute
    - memory
    - network

2. ### COMPUTE OPTIMIZED
- Great for **compute-intensive tasks**
- Requires **high performance processors**
- Example: Batch processing workloads, dedicated gaming servers, machine learning 
- They are of name C

3. ### MEMORY OPTIMIZED
- Fast performance for workloads that **process large datasets**
- They are used for :
    - High performance, relational and non relational databases
    - In-memory **databases optimized for BI**

4. ### STORAGE OPTIMIZED
- Great for **storage-intensive tasks**, for reading and writing large data sets 
- Uses cases:
    - **OLTP**
    - Relational & NoSQL databases 
    - Cache for in-memory databases

## INTRODUCTION TO SECURITY GROUPS
- They will **control how traffic will move IN & OUT** of the Ec2 instances
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
5. RDP - 3389
6. HTTP - 80

### SSH SUMMARY TABLE 
- For linux servers we can use this
- We connect **insiders of our servers** to perform maintaience
- SSH can be used for Mac, Linux and then windows >=10
- We can use **Putty for Windows < 10**
- EC2 instance connect - Available for all( Mac, windows and linux)

### RULE OF THUMB - NEVER ENTER ACCESS KEY AND ACCESS SECRET IN EC2 INSTANCE
- Instead make use of **IAM ROLES**

## EC2 INSTANCES PURCHASING OPTIONS 

## EC2 ON DEMAND
- **Linux or windows - Billing per second** 
- For all billing per hour
- Has the highest cost but no upfront payment 
- No commitments

## EC2 RESERVED INSTANCES
- **72% discount** compared to On-Demand 
- We reserve a specific type of instances (Regions, Os )
- Payment - No upfront, partial upfront and All upfront 
- **Scope: Region or Zone**
- Can BUY AND SELL in market place

## EC2 SAVINGS
- same 7**2% savings** compared to On-demand 
- Commit to a certain type of usage
- Example: **I want to use 10$ per hour for next 3 years**
- Anything above savings plan is **On-demand**

## EC2 SPOT INSTANCES
- **90%** discount 
- Instances can be lost, as we define a max price we can pay. Once moved out it will be terminated 
- Not suited for critical workloads

## EC2 DEDICATED HOSTS
- **dedicated physical server** 
- Used for compliance requirements
- We can use On-demand or Reserved
- Useful for **BYOL - BRING YOUR OWN LICENESE**
- Instance placement in hand 

## EC2 DEDICATED INSTANCES
- **Instance hardware is dedicated to you**
- No control over placement

## EC2 CAPACITY RESERVATIONS 
- Reserve On-Demand instance capacity in a specific AZ
- No time commitment and No **Billing Discounts** 

## SHARED RESPONSIBILITY MODEL

- AWS:
    - Infrastructure
    - Isolation on physical hosts
    - Replacing faulty hardware
- US:
    - Security group roles
    - IAM roles
    - OS patches
    - Data security 

## STORAGE FOR EC2 INSTANCES

1. ### EBS - ELASTIC BLOCK STORAGE

- It acts as a **NETRWORK volume** that is attached to EC2 instances
- They store the data even after termination 
- They can be **mounted only to a single EC2 instance**
- They are bound to specific **AVAILABILITY ZONE**
- Example: An EBS created us-east-1a cannot be attached to us-east-1b, but **we can move volumes across AZ**

2. ### EBS SNAPSHOT
- We can take a EBS and perform a **snapshot = BACKUP**
- We can COPY **SNAPSHOTS ACROSS REGIONS AND AZ**

### EBS SNAPSHOT ARCHIVE
- 75% cheaper to store an snapshot 
- Takes 24 to 72 hours for a restore

## AMI OVERVIEW
- AMI = **AMAZON MACHINE IMAGE**
- We can customize an EC2 using AMI
- We can install softwares, configs and OS.
- AMI are built for a **specific region** and can be copied

## EC2 IMAGE BUILDER
- **Automated creation of VM's and Container images**
- EC2 image builder creates a Builder EC2 instances where we customixze and then a our custom AMI will be created. EC2 builder will perform tests on it 
- And then we can make use of it 
- It can be **run on a schedule**

## EC2 INSTANCE STORE
- Since EBS are network volumes and are limited. WE can make use of EC2 instance store 
- And EC2 Instance **LOSE THEIR DATA IF STORED**

## EFS - ELASTIC FILE SYSTEM
- This can be attached to EC2 (shared network file system)
- It can be **mounted to 100s of EC2 instances**
- Works for **LINNUX EC2**
- Expesnsive 
- It can be attached **ACROSS AZ**

## EBS vs EFS
1. EBS one instance and one AZ, EFS can be done in other way. But we can move EBS to different AZ using Snapcshot

## EFS-IA
- Cost optimized, **92% discount** 
- Files that are not accessed frequently 
- We can create policy, Like *If a file not accessed for 60 days, move to EFS-IA*

## SHARED RESPONSIBILITY
- AWS
    - Infrastructure
    - Replication for data
    - Replacing faulty hardware
- US
    - snapshot / backup procedures
    - Setting up for data encryption 
    - Responsbility of any data on drives

## AMAZON FSX
- managed system for getting **3rd party file system**

### FSX FOR WINDOWS FILER SERVER
- Windows native shared file system 
- meant for **Windows based EC2**
- It is has **SMB and Windows NTFS**

### AMAZON FSX LUSTRE
- For **High Performance Computer (HPC)**
- Used for ML, video processing etc
- Can be connected to comput instances or server. 
- It can make use of S3 for storing data

# ELASTIC LOAD BALANCING AND AUTO SCALING 
### SCALABILITY 
- Means application can handle greater loads by adapting 
- There are two kinds:
    - Vertical scaling 
    - Horizontal scaling = Elasticity
### VERTICAL SCALING 
- **SCALE UP OR DOWN**
- Increase the size 
- Example; t2.micro to t2.large
- Example: Database
- Limit of hardware

### HORIZONTAL SCALING
- **SCALE OUT AND SCALE IN**
- Increasing the number of systems/ EC2
- We need distributed system 

## HIGH AVAILABILITY
- Goes hand in hand with horizontal scaling 
- **RUNNING IN ATLEAST 2 AZ**

### SCALBILITY - ACCOMODATE LARGER LOAD (SCALE UP AND DOWN)
### ELASTICITY - AUTO SCALING (SCALE OUT AND SCALE IN)
### AGILITY - REDUCE TIME AND FASTER

## WHAT IS LOAD BALANCING 
- Servers that forward internet traffic to multiple EC2 instances based on worklaod
- We can use **across multiple AZs**
- AWS takes care of ALB and we are responsible for configuring.
- 4 Kind of load balancers
    - ALB - Application Load balancers - HTTP and HTTPS - **Layer 7**
    - NLB - Network Load balacners - **Layer 4**
    - GATEWAY LOAD BALANCER - **Layer 3** **GENEVE protocol**

![Load balancers](image-2.png)

## AUTO-SCALING GROUPS
- In real-life, the load on websites can change and hence 
- ASG can scale out and scale in by number of request coming in 
- ASG can terminate unhealthy EC2 instances and replace them with healthy EC2

## AUTO SCALING STRATERGIES
1. ### MANUAL SCALING
- Scale manually

2. ### DYNAMIC SCALING 
- Do when a certain event occur like a cloudwatch trigger
- **TARGET SCALING** - I want the average CPU to stay around 40%

3. ### SCHEDULED SCALING 
- Anticiapating ahead of time 

4. ### PREDICTIVE SCALING 
- Use ML for past patterns and scale accordingly 

# AMAZON S3
- It is used for backup and storage 
- S3 buckets are regional level. 
- **NAMING** - GLOBALLY UNIQUE but now we can have **Account Regional Namespace** So we can make use of it at different regions
- Here AWS, will add  suffix themselves
- Naming conventions:
    - No uppercase, No underscore
    - Not an IP
    - Must start lowercase or Number 
    
## AMAZON S3 KEY
- Objects have keys 
- **KEY = FULL PATH**
- KEY is made of **PREFIX NAME + OBJECT**
- Max size to upload once is **50TB**

## AMAZON S3 - SECURITY
1. ### USER BASED
- IAM Policies 

2. ### RESOURCE BASED
- Bucket policies 
- Object Access Control List - ACL - can be disabled
- Bucket Access Control List - can be disabled 
- Encryption keys 

## S3 BUCKET POLICIES
- JSON based policies
- RESOURCES - bucket url
- Effect - Allow or deny
- Actions - Set of API to allow or deny 
- Principal - The account user to apply 

## AMAZON S3 - STATIC WEBSITE HOSTING 
- S3 can host static website and URL deoends on bucket region and aws-region 

## AMAZON S3 VERSIONING 
- We can version files in Amazon S3
- It is enabled at bucket level 
- Everytime we make a change or add a new file we can keep track of it and then it helps is restoring and maintaining
- Suspending versioning will not remove old versions 

## S3 REPLICATION OVERVIEW
### CRR - CROSS REGION RELICATION 
### SRR - SAME REGION REPLICATION 

- We **must enable versioning**
- We need IAM permission for reading and writing 

Use case
- CRR
    1. Compliance, low latency
- SRR
    1. Logs and make replications in prod and other env 

# AMAZON S3 STORAGE CLASSES
1. ### DURABILITY - how many times object is lost (11 9's)
2. ### AVAILABILITY - how readily it is available 99.99%

1. ### S3 STORAGE CLASS
- 99.99% availability
- Low latency and high throughput 
- used for frequent data access
- Use case: Big data analytics, mobile and gaming applications 

2. ### S3 INFREQUENT ACCESS
- Less frequently accessed, but **requires rapid access** when needed
- It is lesser in cost compared to general S3, but will be costed for retrieval 
- 99.99% available 
- Use case - Disaster recovery and Backus

3. ### S3 ONE ZONE INFREQUENT ACCESS
- High durability in a single AZ
- Data lost when AZ is destroyed
- **99.5% availability**
- Use case - store secondary copy of data

4. ### S3 GLACIER STORAGE CLASSES
- Low cost object storage
- Pricing - Price for storage + Object retrieval cost

5. ### S3 GLACIER INSTANT RETRIEVAL
- Millisecond and **rapid retrieval**
- Minimum storage is **90 days**

6. ### S3 GLACIER FLEXIBLE RETRIEVAL
- Expedited (1 to 5 hours), Standard(3 to 5 hours), Bulk(5 to 12 hours)
- Minimum storage is 90 days 

7. ### S3 GLACIER DEEP ARCHIVE
- minimum storage is **180 days**
- standard(12 hours), Bulk(48 hours)

8. ### S3 INTILLIGENT-TIERING
- smoothly moving between storage classes based on usage
- We have to pay auto-tiering fee 


## S3 EXPRESS ONE ZONE

- High performance, but **single AZ** storage class
- Stored in **Directory**
- Handles millions of requests in single digit millisecond latency
- upto **10x better performance** than S3 standard
- Highly durable and available 

## S3 ENCRYPTION
1. **SERVER SIDE ENCRYPTION**
- **BY DEFAULT ENCRYPTED BY AWS**

2. **CLIENT SIDE ENCRYPTION**
- Client does the encryption 

## IAM ACCESS ANALYZER FOR S3
- Ensured only intended people can access your S3
- Powered by IAM Access Analyzer

## SHARED RESPONSBILITY MODEL

- AWS
    - Durability
    - Availability
    - Internal and CVE
- US
    - S3 versioning 
    - S3 bucket policies
    - Loggina nd monitoring 
    - Storage class
    - Encryption on client side 

## AWS SNOWBALL FAMILY 
- Highly-secure, portable device to move data in and out of AWS
- **SNOWBALL EDGE STORAGE OPTIMIZIED** - 210 B STORAGE
- **SNOWBALL EDGE COMPUTE OPTIMIZIED** - 28 TB STORAGE

### SNOWBALL EDGE PRICING 
- **Device usage** and **data transfer out of AWS**
- Data into AWS is **free**

### AWS STORAGE GATEWAY
- BRIDGE **ON-PREM TO CLOUD**
- Hybrid storage services allow on-prem to seamlessly use the AWS cloud

# DATABASE SECTION

1. ## RELATIONAL DATABASES
- Each column is linked to each other 
- can use SQL for querying 
- Good for **OLTP**

2. ## NOSQL DATABASES
- Non relational databases
- Flexible schema 
- JSON is the common way of expressing data 

## RDS - RELATIONAL DATABSE
- **Relational databases**
- It can use **SQL** for querying 
- Aurora is AWS database
- Since RDS and Aurora are AWS managed:
    - Automated is provisioned
    - OS patching 
    - Monitor Dashboard
    - we can do maintainence
- ###### WE CANNOT SSH

## AURORA
- **Aurora is a proprietary**
- PostgresSQL and MySQL are supported
- Aurora costs more than 20%

## AURORA SERVERLESS
- **automated auto-scaling**
- PostgresSQL and MySQL are both supported
- **PAY PER SECOND**

### AURORA DATABASES SHARE SAME STORAGE VOLUME

## RDS DEPLOYMENTS
1. ## READ REPLICAS
- We are gonna create read replicas allowing applications to manage loads
- We can **CREATE 15 READ REPLICAS**
- **WRITING DATA IS DONE TO THE MAIN DB**

### MULTI AZ
- We use this during AZ failover
- We are gonna read write to the main RDS, but we create a **Failover DB** in AWS RDS
- We can have one AZ as a Failover DB

### MULTI REGIONS
- This is for **READ REPLICAS**
- But here we are dealing with Multiple regions instead of AZ

But **WRITE CAN BE DONE TO THE MAIN DB ALONE AND NOT FOR READ REPLICAS**

## AWS ELASTICACHE 
- **In-memory databases**
- Helps reduce load off databases for read intensive workloads using CACHE.
- AWS manages this DB with availability, OS updates and patching etc.

## DYNAMODB
- **NOSQL DATABASES**
- **SERVERLESS DB**
- It is highly available with **3AZ**
- It can handle **millions of request at a time single digit milli-second**

## DYNAMODB ACCELERATOR - DAX
- It is the **in-memory cache for DynamoDB**
- 10x performance improvement 

## DYNAMODB - GLOBAL TABLES
- Make a DynamoDB table accessible with low latency in multiple-regions 
- Users can read and write to these table since it has **active-replication**

## REDSHIFT OVERVIEW
- It is used for PostgresSQL, **DATAWAREHOUSE**
- It is good for **OLAP**
- **Load data every hour, not every second**
- it is good for **analysing and processing data**
- Integrated for **AWS Quicksight or Tableau**

## REDSHIFT SERVERLESS
- WE dont have to worry about scaling and AWS will handle it 

## AMAZON EMR
- It stands for **ELASTIC MAPREDUCE**
- It used to do Bigdata using **Hadoop cluster**
- Has Auto-scaling
- Use-case: Data processing, Ml and bigdata

## ATHENA OVERViEW
- **Serverless** query service to perform analytics agaist **S3 OBJECTS**
- Uses standard SQL query to analyse data 
- Pricing 5$ per TB

## AMAZON QUICKSIGHT
- It is serverless for creating interactive dashboards using ML
- It is automatically scalable 
- It can run on Aurors, Redshift

## DOCUMENTDB
- It is the same for **MongoDB**
- It is a **NOSQL**
- It stores JSON data
- It is stored in **3AZ**

## AMAZON NEPTUNE
- It is fully managed **GRAPH DATABASE**
- A popular graph dataset 
- It is across 3AZ

## AMAZON TIMESTREAM
- It is for **Time series data**
- It is data that is **evolving overtime**
- It is **serverless**

## AWS GLUE
- It is a **ETL** service
- IT is to Extract a dataset, Transform the data and then load it for analysis 

## DMS - DATABASE MIGRATION SERVICE
- **Extract the data from source DB and inject in Target DB**
- Source DB remains during migration 

## WHAT IS A DOCKER
- Docker is a software development platform for deploying apps
- We will package the apps something called as container 
- Running the container runs the application 
- Scalability can be done in seconds 

## ECS - ELASTIC CONTAINER SERVICE
- This is used to launch docker containers to run on AWS
- We need to manage the infra by creating EC instance before
- AWS will take care about **scaling**
- ECS will **place containers on EC2 instances** 

## FARGATE - SERVERLESS
- Launch docker containers on AWS
- This is **serverless and infrastructure is handled by AWS.**

## ECR -  ELASTIC CONTAINER REGISTRY
- Private docker **registry to store docker images**
- It will be later leveraged by ECS and Fargate 

## AMAZON EKS
- EKS = Elastic k8s service
- It is **used to manage k8s cluster on AWS**
- In EKS we will have **EKS Node**, containers will be run inside these nodes
- Can be used in cloud like GCP, Azure or AWS

## LAMBDA OVERVIEW
- **SERVERLESS**
- We only have **virutal functions**
- It **RUNS ON DEMAND**
- Pricing:
    - Pay **Per request**
    - Pay **Per compute time**
- It will be **Event driven**
- **TIME LIMIT IS 15 MINS** 


## API GATEWAY OVERVIEW
- It is used for building **serverless HTTPS API**
- We want an external client accessing lambda function, since lambda is not exposed as an external API that is when we get API Gateway in picture 
![API Gateway](image.png) 
- Supports **RESTful API and Websockets API**

## AWS BATCH
- Helps in batch processing 
- Run 1000000's of batch processing 
- It starts with **start and end time**
- Batch will dynamically launch **EC2 instances and spot instances**

## AWS LIGHTSAIL
- We can get VM's, storage DB and networking 
- Great for people with **Little cloud experience**
- Use cases: simple web applications, websites etc.

## GITHUB SKELETONW WORKFLOW 
```
Workflow
│
├── Event
│
└── Jobs
      │
      ├── Job 1
      │      │
      │      └── Steps
      │             │
      │             ├── Action
      │             ├── Action
      │             └── Shell Command
      │
      └── Job 2
             │
             └── Steps
```
## GITHUB ACTIONS VOCABULARY 

1. **name** - gives your workflow human readable name 
2. **on** - says when should the workflow start
3. **job** - These are tasks that will be executed once the workflow stars executing
4. **runs-on** - Tells where the jobs should run Example: ubuntu-latest, windows-latest, macos-latest
5. **steps** - organized way of executing commands or following the workflow in a job 
6. **uses** - Execute a resuable action created by GitHub or the community
7. **run** - Used to execute commands 
8. **with** - When we use an action, we need input with the action and that's when we use with 

