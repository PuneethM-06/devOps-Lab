# Day 98 — AWS SDK for Go v2, IAM Authentication, S3, STS & CloudWatch

## Goal

Build a Go application that interacts with AWS using the AWS SDK for Go v2.

The project will cover:

- AWS SDK for Go v2
- AWS authentication
- IAM roles and permissions
- S3
- STS
- CloudWatch
- Temporary credentials
- AWS SDK credential provider chain

The learning approach for this day is project-first.

Instead of learning the SDK only through theory, we will build the application and learn the AWS concepts as we interact with the actual services.

---

# 1. What Is AWS SDK for Go v2?

The AWS SDK for Go v2 allows Go applications to interact with AWS services programmatically.

Instead of manually constructing and signing AWS HTTP requests, the SDK provides service clients and APIs.

Conceptually:

Go Application
      ↓
AWS SDK for Go v2
      ↓
AWS API
      ↓
AWS Service

Examples:

- S3
- STS
- CloudWatch
- IAM
- Secrets Manager
- EC2
- Lambda
- And other AWS services

The SDK still uses HTTP underneath, but it handles much of the AWS-specific request construction, signing, serialization, and response handling.

---

# 2. SDK vs AWS CLI

AWS CLI is primarily used by humans and scripts from the command line.

AWS SDK is used by applications.

Example:

AWS CLI:

aws s3 ls

Go Application:

Go Application
      ↓
AWS SDK
      ↓
S3 API

The important distinction is:

AWS CLI → Command-line interaction with AWS

AWS SDK → Programmatic interaction with AWS from application code

---

# 3. AWS SDK Configuration

Before the Go application can interact with AWS, the SDK needs configuration.

Important configuration includes:

- AWS Region
- AWS Credentials
- Other SDK configuration

Conceptually:

AWS Configuration
├── Region
└── Credentials

Region answers:

> Where should the AWS request go?

Credentials answer:

> Which AWS identity is making the request?

Authentication identifies the caller.

Authorization determines whether that identity is allowed to perform the requested operation.

---

# 4. AWS Credential Provider Chain

The AWS SDK can obtain credentials from different sources depending on the environment.

Conceptually:

Go Application
      ↓
AWS SDK
      ↓
Credential Provider Chain
      ↓
Credentials

Possible credential sources include:

- Environment variables
- Shared AWS credentials/configuration
- Container credentials
- EC2 IAM roles
- Other supported credential providers

The application does not need to hardcode a specific credential source into its business logic.

This allows the same Go application to run in different environments.

For example:

Local:

Go Application
      ↓
AWS SDK
      ↓
Local AWS credentials
      ↓
AWS

EC2:

Go Application
      ↓
AWS SDK
      ↓
EC2 IAM Role
      ↓
Temporary Credentials
      ↓
AWS

EKS:

Go Application
      ↓
AWS SDK
      ↓
IAM Role / IRSA
      ↓
Temporary Credentials
      ↓
AWS

---

# 5. IAM Authentication

AWS authentication identifies the AWS identity making the request.

The application needs credentials so AWS can identify the caller.

After authentication, IAM policies determine what the identity is authorized to do.

Mental model:

Credentials
    ↓
Authentication
    ↓
Who are you?
    ↓
IAM Policy
    ↓
Authorization
    ↓
What can you do?

For production workloads, IAM roles should generally be preferred over long-lived IAM user access keys.

---

# 6. IAM Roles

An IAM role is an AWS identity that can be assumed by a trusted entity.

Examples:

- EC2
- Lambda
- ECS
- EKS workloads
- Other AWS services

An IAM role contains two important policy concepts.

## Trust Policy

Defines:

> Who can assume this role?

## Permissions Policy

Defines:

> What can the role do?

Mental model:

Trust Policy
    ↓
Who can assume me?

Permissions Policy
    ↓
What can I do?

Both are important.

---

# 7. Temporary Credentials

When a workload assumes an IAM role, AWS provides temporary security credentials.

These include:

- Access Key ID
- Secret Access Key
- Session Token

The credentials expire automatically.

Static credentials:

Valid until manually revoked or rotated.

Temporary credentials:

Expire automatically.

This reduces the lifetime of stolen credentials.

---

# 8. S3 Fundamentals

Amazon S3 is an object storage service.

The basic structure is:

S3
 └── Bucket
      ├── Object
      ├── Object
      └── Object

A bucket is the container.

An object is the actual stored data.

Objects have:

- Key
- Data
- Metadata

Example:

Bucket:

my-devops-bucket

Object:

logs/application.log

The object key identifies the object inside the bucket.

---

# 9. S3 Operations From Go

Using the AWS SDK for Go v2, the application can interact with S3 programmatically.

Operations we will implement:

- List buckets or objects
- Upload an object
- Download an object
- Read object metadata
- Delete an object

Conceptually:

Go Application
      ↓
AWS SDK for Go v2
      ↓
S3 Client
      ↓
S3 API
      ↓
Bucket / Object

The SDK provides the S3 client and API operations so we don't need to manually construct AWS HTTP requests.

---

# 10. S3 IAM Permissions

The IAM role or identity needs the appropriate S3 permissions.

Examples:

s3:ListBucket

s3:GetObject

s3:PutObject

s3:DeleteObject

Permissions should follow least privilege.

For example, if the application only needs to read objects:

Application Role
      ↓
s3:GetObject
      ↓
Specific S3 bucket/path

It should not automatically receive permissions to delete or modify unrelated resources.

---

# 11. AWS STS

STS stands for:

> AWS Security Token Service

STS provides temporary security credentials.

It is closely connected to IAM roles and role assumption.

Conceptually:

Identity
    ↓
STS
    ↓
Temporary Credentials
    ↓
AWS Services

One important STS operation is:

AssumeRole

AssumeRole allows a trusted identity to assume an IAM role and receive temporary credentials with that role's permissions.

---

# 12. AssumeRole

The basic flow is:

Application / Identity
        ↓
STS AssumeRole
        ↓
IAM Role
        ↓
Temporary Credentials
        ↓
AWS Service

The role's trust policy determines whether the caller can assume the role.

The role's permissions policy determines what the caller can do after assuming it.

Mental model:

Trust Policy
    ↓
Can I assume this role?

Permissions Policy
    ↓
What can I do after assuming it?

---

# 13. AssumeRole From Go

The AWS SDK for Go v2 can be used to interact with STS.

The application can:

1. Create an STS client.
2. Request `AssumeRole`.
3. Receive temporary credentials.
4. Use those credentials to create another AWS service client.
5. Access the AWS resource permitted by the assumed role.

Conceptually:

Go Application
      ↓
AWS SDK
      ↓
STS AssumeRole
      ↓
Temporary Credentials
      ↓
S3 Client
      ↓
S3

This demonstrates how temporary credentials work in practice.

---

# 14. CloudWatch Fundamentals

Amazon CloudWatch is an AWS monitoring and observability service.

For this day, we will focus on two basic concepts:

## Metrics

Metrics are numerical measurements over time.

Examples:

- Request count
- Error count
- CPU usage
- Latency
- Queue depth

## Logs

Logs contain application or system events.

Conceptually:

Application
    ├── Metrics ──→ CloudWatch Metrics
    └── Logs ─────→ CloudWatch Logs

---

# 15. CloudWatch From Go

The AWS SDK for Go v2 provides CloudWatch clients that allow applications to interact with CloudWatch APIs.

For the project, we will demonstrate publishing a custom metric.

Example:

Application
      ↓
AWS SDK for Go v2
      ↓
CloudWatch
      ↓
Custom Metric

Example custom metric:

ApplicationRequests = 10

The goal is not to learn every CloudWatch feature.

The goal is to understand how a Go application can interact with CloudWatch programmatically.

---

# 16. Complete AWS SDK Mental Model

The final architecture we should understand by the end of the project:

Go Application
      ↓
AWS SDK for Go v2
      ↓
Credential Provider Chain
      ↓
AWS Identity / IAM Role
      ↓
Authentication
      ↓
IAM Authorization
      ↓
AWS API
      ↓
S3 / STS / CloudWatch

For role assumption:

Go Application
      ↓
AWS SDK
      ↓
STS AssumeRole
      ↓
Temporary Credentials
      ↓
AWS Service

For S3:

Go Application
      ↓
AWS SDK
      ↓
S3 Client
      ↓
S3 API
      ↓
Bucket / Object

For CloudWatch:

Go Application
      ↓
AWS SDK
      ↓
CloudWatch Client
      ↓
CloudWatch API
      ↓
Metrics / Logs

---

# Project Goal

Build a Go application that demonstrates:

1. Loading AWS SDK configuration
2. Obtaining AWS credentials through the SDK
3. Creating an S3 client
4. Uploading an object
5. Reading an object
6. Deleting an object
7. Understanding the IAM permissions required
8. Using STS
9. Assuming an IAM role
10. Using temporary credentials
11. Creating a CloudWatch client
12. Publishing a custom metric

The project should reinforce the following principle:

> Go application → AWS SDK → AWS authentication → IAM authorization → AWS API → AWS service

---

# Final Takeaway

The AWS SDK is not the security mechanism itself.

The SDK is the application interface to AWS.

IAM provides identity and authorization.

STS provides temporary credentials.

IAM roles allow workloads to avoid long-lived static credentials.

S3 provides object storage.

CloudWatch provides monitoring capabilities.

The SDK connects the Go application to all of these AWS services programmatically.