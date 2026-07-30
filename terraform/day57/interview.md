# Terraform Day 57 – Interview Questions

## 1. Why do we need multiple environments in Terraform?

**Answer:**
Multiple environments (Dev, Staging, Production) allow teams to develop, test, and deploy infrastructure safely. Changes are first validated in lower environments before reaching production, reducing the risk of outages.

---

## 2. What are Terraform Workspaces?

**Answer:**
Terraform Workspaces allow a single Terraform configuration to manage multiple state files. Each workspace has its own state while sharing the same infrastructure code.

Example:

```bash
terraform workspace new dev
terraform workspace new prod
terraform workspace select prod
```

---

## 3. What is `terraform.workspace`?

**Answer:**
`terraform.workspace` is a built-in variable that returns the currently selected workspace.

Example:

```hcl
instance_type = terraform.workspace == "prod" ? "t3.medium" : "t2.micro"
```

---

## 4. What are `.tfvars` files used for?

**Answer:**
`.tfvars` files provide values for input variables defined in `variables.tf`. They allow different environments to use different configurations without changing the Terraform code.

Example:

```hcl
instance_count = 2
instance_type  = "t2.micro"
```

---

## 5. What is the difference between `.tfvars` and `backend.hcl`?

**Answer:**

| `.tfvars` | `backend.hcl` |
|-----------|---------------|
| Stores variable values | Configures the backend |
| Defines desired infrastructure | Defines where Terraform state is stored |
| Used during `plan` and `apply` | Used during `terraform init` |
| Example: EC2 count, instance type | Example: S3 bucket, state key, DynamoDB lock table |

---

## 6. Why should Dev and Production never share the same state file?

**Answer:**
Terraform compares the desired configuration against the current state stored in the state file. If multiple environments share one state, Terraform treats them as a single infrastructure, which can lead to unintended creation, modification, or deletion of resources.

---

## 7. What happens if Dev and Production use the same backend?

**Answer:**
Both environments point to the same state file. Terraform cannot distinguish Dev resources from Production resources and may modify or destroy Production resources while applying Dev changes.

---

## 8. Why is state isolation important?

**Answer:**
State isolation ensures each environment manages only its own infrastructure, preventing accidental changes across environments and enabling independent deployments.

---

## 9. Why don't we create separate `main.tf` files for each environment?

**Answer:**
Maintaining separate Terraform configurations duplicates code, increases maintenance effort, and causes environments to drift apart. A single codebase with different `.tfvars` files is easier to maintain and follows Infrastructure as Code best practices.

---

## 10. Which files would you modify if Production needs 15 EC2 instances instead of 10?

**Answer:**
Only update `prod.tfvars`:

```hcl
instance_count = 15
```

The Terraform code (`main.tf`, modules, variables) remains unchanged because only the input values change.

---

## 11. Why won't Dev be affected when updating Production?

**Answer:**
Because Dev and Production use different backend configurations (or separate workspaces), each has its own Terraform state. Terraform only reads and modifies the Production state.

---

## 12. What is a Terraform backend?

**Answer:**
A backend determines where Terraform stores and reads its state. Common backends include S3, Azure Blob Storage, Google Cloud Storage, and Terraform Cloud.

---

## 13. What is stored in a Terraform state file?

**Answer:**
The state file stores Terraform's mapping between configuration resources and real cloud resources, including resource IDs, metadata, outputs, and dependency information.

---

## 14. Why can't Terraform rely only on AWS APIs instead of a state file?

**Answer:**
AWS APIs can return existing resources, but they don't know which resources belong to a particular Terraform project. The state file records the resources Terraform manages and maps configuration blocks to real infrastructure.

---

## 15. Explain the Terraform workflow during `terraform apply`.

**Answer:**

1. Read the backend configuration.
2. Load the current state from the backend.
3. Read the Terraform configuration.
4. Load variable values from `.tfvars`.
5. Compare current state with desired state.
6. Generate an execution plan.
7. Apply the required changes.
8. Update the state file.

---

## 16. What is environment promotion?

**Answer:**
Environment promotion is the practice of deploying the same Terraform code through Dev → Staging → Production. Only variable values, credentials, and backend configuration differ between environments.

---

## 17. What is a production-ready Terraform folder structure?

```text
terraform/
├── modules/
│   ├── ec2/
│   ├── vpc/
│   └── security-group/
│
├── environments/
│   ├── dev/
│   │   ├── backend.hcl
│   │   └── dev.tfvars
│   ├── staging/
│   │   ├── backend.hcl
│   │   └── staging.tfvars
│   └── prod/
│       ├── backend.hcl
│       └── prod.tfvars
│
├── main.tf
├── variables.tf
├── outputs.tf
├── providers.tf
└── versions.tf
```

---

## 18. What changes between Dev and Production?

**Changes:**
- Variable values (`.tfvars`)
- Backend/state
- AWS credentials or account

**Remains the same:**
- Terraform code
- Modules
- Infrastructure logic

---

## 19. What is the difference between current state and desired state?

**Current State:**
The infrastructure currently managed by Terraform, stored in the state file.

**Desired State:**
The infrastructure described in Terraform configuration files and variable values.

Terraform compares these two states to generate an execution plan.

---

## 20. Why is using one Terraform codebase for all environments considered a best practice?

**Answer:**
It provides a single source of truth, avoids code duplication, improves maintainability, ensures consistency across environments, and reduces configuration drift.