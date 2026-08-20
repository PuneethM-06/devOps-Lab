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
