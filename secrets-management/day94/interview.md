# DAY 94 - HASHICORP VAULT 

1. ### WHAT IS HASHICORP VAULT AND WHAT PROBLEM DOES IT SOLVE 
- Hashicorp vault is a secrets management system used to securely store, control access to, and manage sensitive data 
- Vault can do more than simply return the stored data for the application for its runtime usage 
- It is capable of generating **dynamic secrets on demand**
- By this way, static secrets can be used and bad practices unless they are rotated manually, but with this it solves that problem

- Vault helps avoid:
    - Hardcoding secrets
    - Distributing the same credentials everywhere
    - Long-lived shared credentials
    - Manual credential rotation
    - Uncontrolled access to sensitive data

2. ### STATIC SECRETS vs DYNAMIC SECRETS
- A **static secret** is a credential that already exists and remains the same until someone changes or rotates it.
- A **dynamic secret** is generated when it is requested.
```
Application
      │
      │ Request credentials
      ▼
HashiCorp Vault
      │
      │ Generates new credentials
      ▼
Database
      │
      ▼
username: v-app-123
password: random-password
TTL: 1 hour
```
3. ### VAULT CORE CONCEPTS
1. **Authentication**
- It authenticates the identity of the user or the application requesting the creds 
2. **Policies**
- It checks policies identified with that identity
```
Application A
      │
      ▼
Vault Policy
      │
      ├── Can read database credentials ✅
      └── Cannot read production API keys ❌
```
3. **Secret engines**
- A secret engine is the vault component responsible for managing the secrets 
- Examples:
    - **KV secrets engine** → Stores static key-value secrets
    - **Database secrets engine** → Generates dynamic database credentials
    - **PKI secrets engine** → Generates certificates

4. ### SECRET ENGINES
- A **Secrets Engine** is the Vault component responsible for storing, generating, or managing secrets.
1. **KV Secrets Engine**
- Used to store static key-value secrets.

2. **Database Secrets Engine**
- Used to generate dynamic database credentials.

3. **PKI Secrets Engine**
- Used to generate and manage certificates.

5. ### VAULT AUTHENTICATION AND AUTHORIZATION
- It answers the question:
    1. WHO are you
    2. WHAT are you allowed to do 
- Authentication Example
```
Application
      ↓
Authenticate with Vault
      ↓
Vault verifies identity
      ↓
Vault Token
```
- Authorization 
```
Application
      ↓
Authenticated successfully
      ↓
Vault checks policy
      ↓
Allowed? ── Yes → Access secret
         └─ No  → Access denied
```
6. ### LEASES and TTL
**TTL - Time to Live**
```
Vault generates database credentials

Username: v-user-123
Password: random-password

TTL: 1 hour
```
- Those creds are invalid after 1 hour 

**Lease**
- When Vault generates a dynamic secret, it associates that secret with a lease.
- Think of the lease as Vault's way of tracking:
    - Which secret was generated
    - How long it is valid
    - When it should expire
    - Whether it can be renewed or 

7. ### RENEWAL AND REVOCATION 
- **A lease can be renewed to extend its lifetime, if renewal is allowed.**
```
Application
      ↓
Credentials have TTL: 1 hour
      ↓
Application still needs them
      ↓
Lease renewed
      ↓
TTL extended
```
- **Revocation** - Revocation means Vault invalidates the secret **before or when its lease ends.**
```
Vault-generated credentials
      ↓
TTL expires or lease is revoked
      ↓
Vault revokes credentials
      ↓
Credentials can no longer access the database
```
