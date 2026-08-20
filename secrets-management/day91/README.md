# DAY 91 - AWS SECRETS MANAGER 

1. ### WHAT ARE SECRETS AND WHY DO WE NEED SECRET MANAGEMENT?
- A Secret is a sensitive piece of information that should not be exposed to unauthorized users or systems 

2. ### PROBLEMS WITH HARDCODED SECRETS
1. Exposure
2. Rotation becomes painful 
3. No centralized access control 
4. Poor auditing 
5. Accidential exposure

3. ### WHAT IS AWS SECRETS MANAGER
- A centralized AWS managed service for storing, managinf and retrieving secrets securely 
- AWS Secrets Manager can help you with:
    1. Store secrets
    2. Control access **using IAM**
    3. Encrypt secrets **using KMS**
    4. **Rotate** supported secrets automatically

4. ### HOW DOES AWS SECRETS MANAGER STORE AND PROTECT SECRETS?
**High Level**
```
You / Application
       │
       │ Store secret
       ▼
AWS Secrets Manager
       │
       ├── Encrypts the secret using AWS KMS
       │
       ├── Controls access using IAM
       │
       └── Stores and manages versions
```
1. **Secrets manager** - Stores and manages the secrets 
2. **AWS KMS** - handles the encryption at rest 
3. **IAM** - decides who or what is allowed to retrieve it
```
Application
     │
     │ "I need this secret"
     ▼
AWS Secrets Manager
     │
     │ Is this identity allowed?
     ▼
IAM Authorization
     │
     ├──  No → Access denied
     │
     └──  Yes
             │
             ▼
        Secret retrieved
```
> Secrets Manager stores the secret, KMS protects it through encryption, and IAM controls access to it.

5. ### SECRETS, VERSIONS AND STAGING LABELS
- Everytime a password is rotated it creates a new version of it.
```
Today:
DB_PASSWORD = password-v1

Later, after rotation:
DB_PASSWORD = password-v2
```
Applications usally arent aware of version and instead they use **staging labels by AWS**
- The two important ones are:
    1. **AWSCURRENT** - Latest version 
    2. **AWSPREVIOUS** - The version before latest 

6. ### ENCRYPTION AND ROLE OF AWS KMS
- **KMS** - Provides protection to secrets at rest by encrypting them 

7. ### IAM PERMISSIONS FOR ACCESSING SECRETS 
```
{
  "Effect": "Allow",
  "Action": "secretsmanager:GetSecretValue",
  "Resource": "arn:aws:secretsmanager:region:account-id:secret:my-db-secret"
}
```
- `"Action": "secretsmanager:GetSecretValue",` is important and this is the role 
