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
