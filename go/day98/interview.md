# Day 100 — AWS SDK for Go v2, IAM, S3, STS & CloudWatch

## 1. What is the AWS SDK for Go v2?

The AWS SDK for Go v2 is a set of Go libraries that allows applications to interact with AWS services programmatically.

It provides service-specific clients and APIs, so we don't need to manually construct and sign AWS HTTP requests.

Examples:

- S3 client
- STS client
- CloudWatch client

The SDK handles much of the AWS-specific request and response handling.

---

## 2. How does the AWS SDK obtain credentials?

The AWS SDK uses a credential provider chain to find available credentials.

Depending on the environment, credentials can come from:

- Environment variables
- Shared AWS credentials/configuration
- Container credentials
- IAM roles

This means the same Go application can run in different environments without hardcoding credentials into the source code.

For example:

Local → Local AWS credentials

EC2 → EC2 IAM Role

EKS → IAM Role through IRSA

---

## 3. What is the difference between authentication and authorization in AWS?

Authentication determines:

> Who are you?

Authorization determines:

> What are you allowed to do?

Credentials are used to authenticate the caller.

IAM policies determine what actions that identity is authorized to perform.

Example:

Credentials
    ↓
Authentication
    ↓
Identify caller
    ↓
IAM Policy
    ↓
Authorization
    ↓
Allow / Deny action

---

## 4. What is AWS STS and what does AssumeRole do?

AWS STS (Security Token Service) provides temporary security credentials.

`AssumeRole` allows a trusted identity to assume an IAM role and receive temporary credentials associated with that role.

The trust policy determines who can assume the role.

The permissions policy determines what the role can do.

Flow:

Identity
    ↓
STS AssumeRole
    ↓
IAM Role
    ↓
Temporary Credentials
    ↓
AWS Service

---

## 5. Why should applications use IAM roles and temporary credentials instead of long-lived access keys?

Long-lived access keys remain valid until they are manually rotated, disabled, or deleted.

If they are leaked, an attacker may continue using them until they are revoked.

IAM roles provide temporary credentials that expire automatically.

This:

- Reduces the lifetime of stolen credentials
- Avoids storing long-lived credentials in applications
- Reduces credential distribution and rotation problems
- Works well with AWS workloads such as EC2, ECS, Lambda, and EKS

The IAM role should also follow the principle of least privilege so the workload has only the permissions it actually needs.

---

## 6. What is S3?

Amazon S3 is an object storage service.

The basic structure is:

S3
 └── Bucket
      ├── Object
      ├── Object
      └── Object

A bucket is the container and an object is the data stored inside it.

Objects are identified using an object key.

---

## 7. What are common S3 operations an application can perform?

Common operations include:

- List objects
- Upload objects
- Download objects
- Read object metadata
- Delete objects

A Go application can perform these operations using the S3 client provided by the AWS SDK for Go v2.

---

## 8. What IAM permissions might a Go application need to interact with S3?

The required permissions depend on what the application needs to do.

Examples:

- `s3:ListBucket`
- `s3:GetObject`
- `s3:PutObject`
- `s3:DeleteObject`

The permissions should follow least privilege.

If an application only needs to read objects, it should not receive unnecessary write or delete permissions.

---

## 9. What is the difference between S3 and CloudWatch?

S3 is primarily an object storage service.

CloudWatch is an AWS monitoring and observability service.

S3 stores objects such as:

- Files
- Backups
- Logs
- Artifacts

CloudWatch handles monitoring data such as:

- Metrics
- Logs
- Alarms

---

## 10. What is CloudWatch?

Amazon CloudWatch is an AWS monitoring and observability service.

Two important concepts are:

### Metrics

Numerical measurements over time.

Examples:

- CPU usage
- Request count
- Error count
- Latency

### Logs

Records of application or system events.

A Go application can interact with CloudWatch programmatically using the AWS SDK for Go v2.

---

## 11. What is the overall flow when a Go application accesses an AWS service?

The general flow is:

Go Application
    ↓
AWS SDK for Go v2
    ↓
Credential Provider Chain
    ↓
AWS Credentials
    ↓
Authentication
    ↓
IAM Authorization
    ↓
AWS API
    ↓
AWS Service

For example:

Go Application
    ↓
AWS SDK
    ↓
S3 Client
    ↓
S3 API
    ↓
S3 Bucket

---

## Key Mental Models

### SDK

> SDK = Application interface for interacting with AWS programmatically.

### Credentials

> Credentials = Used to authenticate the AWS caller.

### IAM

> IAM = Controls what the authenticated identity is authorized to do.

### STS

> STS = Provides temporary security credentials.

### S3

> S3 = Object storage.

### CloudWatch

> CloudWatch = Monitoring and observability.

### IAM Role

> IAM Role = AWS identity that workloads can assume to obtain permissions and temporary credentials.