# Day 52 Interview Questions

## Terraform Basics

### 1. What is Infrastructure as Code (IaC)?

**Answer:**
Infrastructure as Code (IaC) is the practice of provisioning and managing infrastructure using code instead of manual configuration. It enables automation, consistency, version control, and repeatability.

---

### 2. What are the advantages of IaC?

**Answer:**

- Automation
- Consistency
- Repeatability
- Version Control
- Faster deployments
- Reduced human errors

---

### 3. What is Terraform?

**Answer:**
Terraform is an open-source Infrastructure as Code (IaC) tool developed by HashiCorp that allows you to provision and manage infrastructure using declarative configuration files.

---

### 4. What does Declarative mean?

**Answer:**
Declarative means you describe the desired end state, and Terraform determines how to achieve it.

Example:

"I want one S3 bucket."

Terraform figures out the API calls automatically.

---

### 5. Declarative vs Imperative?

| Declarative | Imperative |
|-------------|------------|
| Describes desired state | Describes every step |
| Terraform | Bash Script |
| Simpler | More control |
| Easier maintenance | More code |

---

## Terraform Workflow

### 6. Explain the Terraform workflow.

**Answer:**

1. Write configuration
2. terraform fmt
3. terraform init
4. terraform validate
5. terraform plan
6. terraform apply
7. terraform destroy

---

### 7. What is a Provider?

**Answer:**
A Provider is a plugin that enables Terraform to communicate with a cloud platform or service such as AWS, Azure, or GCP.

Examples:

- AWS Provider
- Azure Provider
- Google Provider

---

### 8. What is a Resource?

**Answer:**
A Resource is an infrastructure object managed by Terraform.

Examples:

- EC2 Instance
- S3 Bucket
- VPC
- IAM User

---

### 9. What is Desired State?

**Answer:**
Desired State is the infrastructure configuration defined in Terraform files that Terraform tries to create and maintain.

---

### 10. What is terraform.tfstate?

**Answer:**
terraform.tfstate is a state file that stores information about resources managed by Terraform. It maps Terraform configuration to real infrastructure.

---

## Commands

### 11. What does terraform fmt do?

**Answer:**
Formats Terraform configuration files according to Terraform's standard style.

---

### 12. Does terraform fmt modify infrastructure?

**Answer:**
No.

It only formats code.

---

### 13. What does terraform validate do?

**Answer:**
Checks Terraform configuration for syntax and configuration errors without creating infrastructure.

---

### 14. Why must terraform init be run before validate?

**Answer:**
Because init downloads the required provider plugins. Validate requires the provider schemas to verify the configuration.

---

### 15. What does terraform init do?

**Answer:**

- Downloads providers
- Creates `.terraform`
- Creates `.terraform.lock.hcl`
- Initializes the working directory

---

### 16. Does terraform init create infrastructure?

**Answer:**
No.

It only prepares the working directory.

---

### 17. What is `.terraform`?

**Answer:**
A local working directory created by Terraform that stores downloaded providers and initialization data.

---

### 18. Should `.terraform` be committed to Git?

**Answer:**
No.

It should be added to `.gitignore`.

---

### 19. What is `.terraform.lock.hcl`?

**Answer:**
A dependency lock file that records the exact provider versions used to ensure consistent deployments.

---

### 20. Should `.terraform.lock.hcl` be committed?

**Answer:**
Yes.

It should be committed to version control.

---

### 21. What does terraform plan do?

**Answer:**
Generates an execution plan by comparing the desired state, current state, and actual infrastructure.

No infrastructure is changed.

---

### 22. What do the symbols in terraform plan mean?

| Symbol | Meaning |
|---------|---------|
| + | Create |
| ~ | Update |
| - | Delete |
| -/+ | Replace |

---

### 23. What does terraform apply do?

**Answer:**
Executes the execution plan and creates, updates, or deletes infrastructure.

---

### 24. What does terraform destroy do?

**Answer:**
Deletes all infrastructure managed by the current Terraform configuration.

---

## Scenario Questions

### 25. Why should you run terraform plan before apply?

**Answer:**
To review the proposed infrastructure changes before they are executed and avoid unintended modifications.

---

### 26. What happens if the provider is not initialized?

**Answer:**
Terraform commands like validate or plan fail because the provider plugin is missing.

---

### 27. What happens if terraform.tfstate is deleted?

**Answer:**
Terraform loses track of the infrastructure it manages and can no longer accurately determine the current state.

---

### 28. Why shouldn't you manually edit terraform.tfstate?

**Answer:**
Manual changes can corrupt the state file and cause Terraform to make incorrect infrastructure changes.

---

### 29. Explain the Terraform lifecycle.

**Answer:**

Write Configuration
↓

fmt
↓

init
↓

validate
↓

plan
↓

apply
↓

Infrastructure Created
↓

destroy
↓

Infrastructure Deleted

---

### 30. Which files should be committed to Git?

**Commit:**

- *.tf
- .terraform.lock.hcl
- README.md

**Ignore:**

- .terraform/
- terraform.tfstate
- terraform.tfvars (if sensitive)