# DAY 91 AND 92 
1. ### WHY STATIC AWS CREDENTIAL ARE A PROBLEM
- Credentials like Access Key ID and Secret Access key are long lived and they need to be changed/deleted by someone manually.
- Due to which if they are leaked they can be a threat to the infrastructure 
- It is not also not a right idea to give distribute them to appplications and hence using short term creds are the right way of doing it 


2. ### IAM ROLE vs IAM USER
- **IAM USER** - Is an AWS identity representing a person or application to which roles/permissions can be attached
```
IAM User
   │
   └── Permissions
```
- **IAM ROLE** - It is identity with set of permissions 
- The entity could be:
    1. EC2 instance 
    2. Lambda
    3. ECS task
    4. K8s workload

3. ### WHAT AN IAM ROLE ACTUALLY IS
- It ha 2 major parts:
    1. **Permisions policy** - What the role is allowed to do 
    2. **Trust Policy** - Who or what is allowed to assume the role 
4. ### TEMPORARY CREDENTIALS 
These typically includes:
    1. Access Key ID
    2. Secret Access Key
    3. Session token 
```
Static credentials:
Valid until manually revoked/rotated 

Temporary credentials:
Valid for a limited time
        │
        ▼
Expire automatically 
```
5. ### HOW WORKLOADS GET CREDETIALS WITHOUT HARDCODING THEM 
- This is exactly where IAM roles come into picture 
- Different AWS workloads can receive an IAM role
```
EC2 Instance ──────► IAM Role
Lambda Function ───► Execution Role
ECS Task ──────────► Task Role
EKS Pod ───────────► IAM Role through IRSA
```
```
Application / Workload
        │
        │ Has an IAM role associated with it
        ▼
AWS provides temporary credentials
        │
        ▼
Application uses AWS APIs
```
6. ### ASSUME ROLE AND TRUST POLICIE
- An identiy does not randomly start using an IAM role. It must be allowed by the **trust policy**
```
EC2 Instance
      │
      ▼
Trust Policy says:
"EC2 is allowed to assume this role"
      │
      ▼
IAM Role
      │
      ▼
Permissions:
s3:GetObject
```
- So AssumeRole is essentially the process where a trusted entity takes on an IAM role and receives temporary credentials with that role's permissions.