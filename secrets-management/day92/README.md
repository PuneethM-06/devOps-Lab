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

7. ### LEAST PRIVILGE AND BLAST RADIUS
- If the application using BadRole gets compromised, the attacker inherits a massive set of permissions.**This is blast radius**
- The more permissions an identity has, the more damage can potentially be done if that identity is compromised.

8. ### IRSA: IAM ROLES AND KUBERNETES SERVICE ACCOUNT
- **IRSA stands for IAM Roles for Service ACcount**s
- **This is used to help pods running in k8s service access AWS services**
```
Kubernetes Pod
      │
      ▼
Kubernetes ServiceAccount
      │
      ▼
Associated IAM Role
      │
      ▼
Temporary AWS Credentials
      │
      ▼
AWS Secrets Manager
```
9. ### WHY K8S SECRETS ALONE ARENT ENOUGH ### WHY K8S SECRETS ALONE AREN'T ENOUGH

Kubernetes provides a `Secret` resource for storing sensitive information such as:

- Database passwords
- API keys
- Tokens

However, Kubernetes Secrets alone are not a complete secrets management solution.

The main problem is the lack of a centralized external source of truth.

Without an external secrets management system, secrets may need to be manually copied from a central secret store into Kubernetes:

AWS Secrets Manager
        │
        │ Manual copy
        ▼
Kubernetes Secret
        │
        ├── Cluster A
        ├── Cluster B
        └── Cluster C

This creates multiple copies of the same secret.

If the secret is later rotated in AWS Secrets Manager, the Kubernetes Secret may still contain the old value:

AWS Secret updated
        │
        ▼
Kubernetes Secret still has old value

Someone or something must update the Kubernetes Secret.

Another important point is that Base64 encoding is not encryption. Kubernetes Secrets are encoded using Base64, which means they should not be treated as secure merely because their values are encoded.

Kubernetes Secrets are useful for delivering secret data to applications inside the cluster, but they do not automatically provide:

- A centralized external source of truth
- Automatic synchronization with external secret stores
- Secret lifecycle management
- Automatic updates when an external secret is rotated

This is where External Secrets Operator helps:

AWS Secrets Manager
        │
        │ Source of truth
        ▼
External Secrets Operator
        │
        │ Synchronizes
        ▼
Kubernetes Secret
        │
        ▼
Application / Pod
```
 