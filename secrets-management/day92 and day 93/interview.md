# Day 92–93 — IAM Roles Everywhere and External Secrets Operator

## IAM Roles Everywhere

The main principle is:

> Do not store or distribute long-lived AWS access keys to workloads. Use IAM roles and temporary credentials instead.

---

### Why Static AWS Credentials Are a Problem

Static AWS credentials usually include:

- `AWS_ACCESS_KEY_ID`
- `AWS_SECRET_ACCESS_KEY`

These credentials can be stored in:

- Application code
- `.env` files
- Configuration files
- CI/CD systems
- Developer machines
- Kubernetes Secrets

The problem is that they are long-lived.

If they are leaked, they remain usable until someone manually rotates, disables, or deletes them.

Distributing the same credentials across multiple applications also creates multiple copies and makes rotation difficult.

A better approach is:

Application / Workload
        ↓
IAM Role
        ↓
Temporary Credentials
        ↓
AWS Service

---

### IAM User vs IAM Role

An IAM User is an AWS identity that can have permissions and long-term credentials.

An IAM Role is an AWS identity that can be assumed by a trusted entity.

Examples include:

- EC2
- Lambda
- ECS
- EKS workloads
- Other AWS services

When a role is assumed, AWS provides temporary credentials.

---

### Trust Policy vs Permissions Policy

An IAM role has two important parts:

Trust Policy
    → Who can assume this role?

Permissions Policy
    → What can this role do?

For example:

EC2
    ↓
Trust policy allows EC2
    ↓
IAM Role
    ↓
Permissions allow `s3:GetObject`
    ↓
S3

The mental model is:

> Trust policy gets an identity into the role.

> Permissions policy determines what it can do after assuming the role.

---

### Temporary Credentials

When a workload assumes an IAM role, it receives temporary security credentials:

- Access Key ID
- Secret Access Key
- Session Token

These credentials expire automatically.

Static credentials:

Valid until manually revoked or rotated.

Temporary credentials:

Expire automatically after a limited period.

This reduces the lifetime of stolen credentials.

---

### IAM Roles for AWS Workloads

AWS workloads can receive permissions without storing access keys.

Examples:

EC2 Instance
    ↓
IAM Role

Lambda Function
    ↓
Execution Role

ECS Task
    ↓
Task Role

EKS Pod
    ↓
IAM Role through IRSA

The workload receives temporary credentials and uses them to access AWS services.

---

### Least Privilege and Blast Radius

A workload should receive only the permissions it actually requires.

For example:

Application Role
    ↓
`secretsmanager:GetSecretValue`
    ↓
Specific database secret

Instead of:

Application Role
    ↓
Access to all secrets

Least privilege reduces the blast radius if a workload is compromised.

---

## IRSA — IAM Roles for Service Accounts

IRSA stands for IAM Roles for Service Accounts.

It allows a Kubernetes ServiceAccount to be associated with an IAM role.

The flow is:

Kubernetes Pod
        ↓
Kubernetes ServiceAccount
        ↓
IRSA
        ↓
IAM Role
        ↓
Temporary AWS Credentials
        ↓
AWS Service

The ServiceAccount provides the identity inside Kubernetes.

The IAM role provides permissions in AWS.

This removes the need to store static AWS access keys inside the Kubernetes cluster.

---

# External Secrets Operator

## Why Kubernetes Secrets Alone Aren't Enough

Kubernetes Secrets can store sensitive information for workloads inside the cluster.

However, manually copying secrets from an external secret manager creates duplication.

AWS Secrets Manager
        ↓
Manual copy
        ↓
Kubernetes Secret

If the secret is later rotated in AWS Secrets Manager, the Kubernetes Secret may still contain the old value.

Kubernetes Secrets also do not automatically provide:

- A centralized external source of truth
- Automatic synchronization
- Secret lifecycle management

This is where External Secrets Operator helps.

---

## What Is External Secrets Operator?

External Secrets Operator, or ESO, is a Kubernetes operator that retrieves secrets from external secret management systems and synchronizes them into Kubernetes Secrets.

For example:

AWS Secrets Manager
        ↓
External Secrets Operator
        ↓
Kubernetes Secret
        ↓
Application

AWS Secrets Manager remains the source of truth.

ESO is responsible for retrieving and synchronizing the secret.

---

## How External Secrets Operator Works

ESO needs to know:

1. Where to fetch the secret from
2. Which secret to fetch
3. Which Kubernetes Secret to create or update

The flow is:

External Secret Store
        ↓
SecretStore / ClusterSecretStore
        ↓
ExternalSecret
        ↓
External Secrets Operator
        ↓
Kubernetes Secret
        ↓
Application

ESO continuously reconciles the desired state and can synchronize updates based on its configured refresh behavior.

---

## `SecretStore` vs `ClusterSecretStore`

### `SecretStore`

A `SecretStore` is namespace-scoped.

It can be used only within the namespace where it is created.

### `ClusterSecretStore`

A `ClusterSecretStore` is cluster-scoped.

It can be referenced across multiple namespaces in the cluster.

In simple terms:

- `SecretStore` = Namespace-specific
- `ClusterSecretStore` = Cluster-wide

Both define where and how ESO connects to an external secret provider.

---

## `ExternalSecret`

An `ExternalSecret` defines:

- Which `SecretStore` or `ClusterSecretStore` to use
- Which external secret to fetch
- Which Kubernetes Secret to create or update

Mental model:

> `SecretStore` = Where

> `ExternalSecret` = What to fetch and what to create

Example:

AWS Secrets Manager
    ↓
`production/database-password`
    ↓
ExternalSecret
    ↓
Kubernetes Secret: `database-secret`

---

## How ESO Authenticates to AWS

External Secrets Operator should not use static AWS access keys.

Instead, it can use IRSA.

The flow is:

ESO Pod
    ↓
Kubernetes ServiceAccount
    ↓
IRSA
    ↓
IAM Role
    ↓
Temporary AWS Credentials
    ↓
AWS Secrets Manager

The IAM role contains the required permissions, such as:

`secretsmanager:GetSecretValue`

The role should follow the principle of least privilege and only access the secrets that ESO actually requires.

---

# Complete End-to-End Flow

AWS Secrets Manager
    ↓
Stores the secret and acts as the source of truth
    ↓
SecretStore / ClusterSecretStore
    ↓
Defines where ESO connects
    ↓
ExternalSecret
    ↓
Defines which secret to fetch and which Kubernetes Secret to create
    ↓
External Secrets Operator
    ↓
Uses ServiceAccount + IRSA + IAM Role
    ↓
Receives temporary AWS credentials
    ↓
Fetches the secret from AWS Secrets Manager
    ↓
Creates or updates the Kubernetes Secret
    ↓
Application Pod consumes the secret

# Final Mental Model

Don't distribute static AWS credentials
        ↓
Use IAM roles
        ↓
Workloads receive temporary credentials
        ↓
For EKS, use ServiceAccount + IRSA
        ↓
ESO authenticates to AWS
        ↓
SecretStore defines where to connect
        ↓
ExternalSecret defines what to fetch
        ↓
ESO synchronizes the secret
        ↓
Kubernetes Secret is created or updated
        ↓
Application consumes the secret