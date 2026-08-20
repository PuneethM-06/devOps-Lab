# AWS Secrets Manager — Interview Questions

## 1. How does an application retrieve a secret from AWS Secrets Manager at runtime?

The application requests the secret using its AWS identity. IAM checks whether that identity has permission to access the requested secret. If the request is authorized, AWS Secrets Manager returns the secret, and the application uses it at runtime.

---

## 2. What is the difference between IAM and AWS KMS in the context of AWS Secrets Manager?

IAM controls who or what is allowed to access a secret.

AWS KMS encrypts and protects the secret data at rest.

In simple terms:

- IAM → Who can access the secret?
- KMS → How is the secret data protected?

---

## 3. What is secret rotation, and why is it important?

Secret rotation is the process of replacing an existing secret with a new one.

It is important because if a secret is exposed or compromised, an attacker may continue using it for as long as it remains valid. Rotation limits how long a credential can remain useful.

Rotation can be:

1. **Manual** — Someone manually changes the secret.
2. **Automatic** — AWS Secrets Manager can rotate supported secrets using a configured rotation process.

---

## 4. What are `AWSCURRENT` and `AWSPREVIOUS` in AWS Secrets Manager?

AWS Secrets Manager stores different versions of a secret and uses staging labels to identify them.

- `AWSCURRENT` points to the version currently in use.
- `AWSPREVIOUS` points to the version that was current before the latest rotation.

Example:

Version 1 → password-v1 → AWSPREVIOUS  
Version 2 → password-v2 → AWSCURRENT

---

## 5. Why is the principle of least privilege important when giving an application access to AWS Secrets Manager?

The application should receive only the permissions required to perform its job and nothing more.

For example, an application that only needs to retrieve a database password should be given permission to access that specific secret, rather than access to all secrets.

This reduces the blast radius if the application or its identity is compromised.