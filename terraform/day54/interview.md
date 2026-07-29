# Terraform Day 54 - Interview Questions

## 1. What is Terraform State?

**Answer:**

Terraform State is a file that records the current infrastructure managed by Terraform. It maps Terraform configuration to real cloud resources and enables Terraform to plan, apply, and destroy infrastructure safely.

---

## 2. What is the default Terraform state file?

**Answer:**

```
terraform.tfstate
```

---

## 3. Why does Terraform need a state file?

**Answer:**

Terraform uses the state file to:

- Track managed resources
- Map configuration to cloud resources
- Detect changes
- Plan infrastructure updates efficiently

---

## 4. Is `terraform.tfstate` a Terraform configuration file?

**Answer:**

No.

It is a JSON data file generated and managed by Terraform.

---

## 5. Name the important fields inside `terraform.tfstate`.

**Answer:**

- version
- terraform_version
- serial
- lineage
- outputs
- resources

---

## 6. What is the purpose of `serial`?

**Answer:**

It is the version number of the state file and increments whenever Terraform successfully updates the state.

---

## 7. What is `lineage`?

**Answer:**

A unique identifier for a Terraform state that distinguishes it from other state files.

---

## 8. What is the most important section of the state file?

**Answer:**

The `resources` section because it maps Terraform resources to actual cloud resources.

---

## 9. Can you edit `terraform.tfstate` manually?

**Answer:**

No. It should only be modified through Terraform commands.

---

## 10. What is Infrastructure Drift?

**Answer:**

Infrastructure Drift occurs when the actual infrastructure differs from the Terraform configuration, usually because of manual changes outside Terraform.

---

## 11. How does Terraform detect drift?

**Answer:**

By comparing:

- Terraform configuration
- Terraform state
- Actual infrastructure in the cloud

using `terraform plan`.

---

## 12. How is drift corrected?

**Answer:**

Run:

```bash
terraform apply
```

Terraform updates the infrastructure to match the configuration.

---

## 13. What is a Terraform Backend?

**Answer:**

A backend determines where Terraform stores its state and how state operations such as locking are performed.

---

## 14. What is the default backend?

**Answer:**

Local backend.

---

## 15. Why use an S3 backend?

**Answer:**

To centrally store Terraform state so multiple team members use the same state file.

---

## 16. How do you configure an S3 backend?

**Answer:**

Using a backend block:

```hcl
terraform {
  backend "s3" {
    bucket = "my-terraform-state"
    key    = "terraform.tfstate"
    region = "us-east-1"
  }
}
```

Then run:

```bash
terraform init
```

to initialize the backend.

---

## 17. What does `terraform init` do after adding a backend?

**Answer:**

It initializes the backend and configures Terraform to read and write the state from the remote backend.

---

## 18. What is State Locking?

**Answer:**

State locking prevents multiple users or processes from modifying the Terraform state simultaneously.

---

## 19. Why is state locking important?

**Answer:**

It prevents race conditions, state corruption, and accidental overwrites when multiple users run Terraform operations at the same time.

---

## 20. Which AWS service is commonly used for Terraform state locking?

**Answer:**

Amazon DynamoDB.

---

## 21. What happens if two engineers run `terraform apply` simultaneously without locking?

**Answer:**

A race condition can occur where both users overwrite each other's state updates, leading to an inconsistent or corrupted state file.

---

## 22. Which AWS services are commonly used together for remote state management?

**Answer:**

- Amazon S3 → Stores the Terraform state.
- Amazon DynamoDB → Provides state locking.

---

# Quick Revision

- `terraform.tfstate` stores Terraform's state.
- State is required for `plan`, `apply`, and `destroy`.
- `serial` increments after every successful state update.
- `lineage` uniquely identifies the state.
- Drift occurs when infrastructure differs from Terraform configuration.
- `terraform plan` detects drift.
- `terraform apply` reconciles drift.
- Local backend stores state on your machine.
- Remote backend stores state in Amazon S3.
- `terraform init` initializes the configured backend.
- DynamoDB provides state locking.
- S3 + DynamoDB is the recommended backend setup for team environments.