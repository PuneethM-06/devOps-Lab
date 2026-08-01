# Terraform CI/CD & State Management — Interview Notes

## 1. What is `terraform fmt -check`?

Formats Terraform configuration according to the standard style.

- `terraform fmt` formats the files.
- `terraform fmt -check` verifies formatting without modifying files.
- Returns a non-zero exit code if formatting is incorrect, making it ideal for CI pipelines.

---

## 2. Why run `terraform init` before `terraform validate`?

`terraform validate` requires providers and modules to be initialized.

`terraform init`:
- Downloads providers
- Downloads modules
- Configures the backend
- Creates the `.terraform` directory

Without initialization, validation fails.

---

## 3. What does `terraform validate` do?

Validates the Terraform configuration.

Checks:
- Syntax
- Resource references
- Variable definitions
- Provider configuration
- Module structure

It does **not** contact AWS or create resources.

---

## 4. Why run `terraform plan` in CI?

To preview infrastructure changes before deployment.

It shows:
- Resources to create
- Resources to update
- Resources to destroy

This allows reviewers to verify changes before merging.

---

## 5. Why use `terraform plan -out=tfplan`?

Instead of displaying the plan only, Terraform saves the execution plan into a binary file.

Benefits:
- The exact reviewed plan can later be applied.
- Avoids recalculating the execution plan within the same workflow.

---

## 6. Why use `terraform show`?

Displays the contents of either:
- Terraform state
- A saved plan file

Example:

```bash
terraform show tfplan
```

Useful for converting the binary plan into human-readable output.

---

## 7. Why comment the Terraform plan on a Pull Request?

Allows reviewers to see infrastructure changes without downloading artifacts.

Benefits:
- Easier code review
- Better collaboration
- Detect unintended infrastructure changes early

---

## 8. Why couldn't we reuse the uploaded artifact after merging the PR?

GitHub Actions artifacts are scoped to a **single workflow run**.

Flow:

PR Workflow
→ Upload tfplan
→ Workflow ends

Merge
→ New workflow run
→ Previous artifact is unavailable

Artifacts cannot be downloaded across separate workflow runs using `actions/download-artifact`.

---

## 9. How was this solved?

Instead of downloading the previous plan:

Deployment workflow performs:

```bash
terraform init
terraform plan -out=tfplan
terraform apply tfplan
```

A fresh execution plan is generated and immediately applied.

---

## 10. Why is `terraform apply tfplan` better than `terraform apply`?

Using:

```bash
terraform apply tfplan
```

applies a previously generated execution plan.

Advantages:
- Predictable deployment
- No recalculation within the same workflow
- Ensures the reviewed plan is applied

---

## 11. What is `terraform import`?

Imports existing infrastructure into Terraform state.

It allows Terraform to manage resources that were created outside Terraform.

Example:

```bash
terraform import aws_s3_bucket.logs company-logs
```

---

## 12. Does `terraform import` generate Terraform code?

No.

It only updates the Terraform state.

You must first create the resource block in the `.tf` files and then import the existing resource.

---

## 13. What is the workflow for importing an existing resource?

1. Resource already exists in AWS.
2. Write the Terraform resource block.
3. Run `terraform import`.
4. Run `terraform plan`.
5. Verify that Terraform reports **No changes**.

---

## 14. What does `terraform state list` do?

Lists every resource currently tracked in the Terraform state.

Example:

```bash
terraform state list
```

It reads only the Terraform state and does not query AWS.

---

## 15. What does `terraform state show` do?

Displays detailed information about a specific resource stored in the Terraform state.

Example:

```bash
terraform state show aws_s3_bucket.logs
```

Useful for viewing IDs, ARNs, tags, and other tracked attributes.

---

## 16. What does `terraform state mv` do?

Moves or renames a resource within the Terraform state without modifying the actual infrastructure.

Example:

```bash
terraform state mv aws_s3_bucket.bucket aws_s3_bucket.logs
```

Commonly used when refactoring code or moving resources into modules.

---

## 17. What does `terraform state rm` do?

Removes a resource from Terraform state without deleting the actual infrastructure.

Example:

```bash
terraform state rm aws_s3_bucket.logs
```

Terraform simply stops managing the resource.

---

## 18. What does `terraform state pull` do?

Downloads or prints the current Terraform state in JSON format.

Example:

```bash
terraform state pull
```

Useful for:
- Inspecting remote state
- Backing up state
- Debugging

---

## 19. What does `terraform state push` do?

Uploads a state file to the configured backend.

Example:

```bash
terraform state push terraform.tfstate
```

Typically used only for:
- Disaster recovery
- State migration
- Restoring backups

Incorrect usage can corrupt Terraform state.

---

## 20. Why should Terraform state never be edited manually?

The state file is Terraform's source of truth.

Manual edits can:
- Corrupt state
- Lose tracked resources
- Cause incorrect plans
- Trigger unwanted infrastructure changes

State modifications should always be performed using Terraform state commands.

---

## 21. Why are remote backends recommended?

Remote state enables:
- Team collaboration
- State locking
- Centralized storage
- Consistent infrastructure management
- CI/CD integration

Common production backend:
- Amazon S3 (state storage)
- DynamoDB (state locking)

---

## 22. Explain the Terraform CI/CD pipeline built.

### Pull Request

```
terraform fmt
        ↓
terraform validate
        ↓
terraform plan
        ↓
terraform show
        ↓
Comment plan on PR
```

### Merge to main

```
terraform fmt
        ↓
terraform validate
        ↓
terraform plan
        ↓
terraform apply tfplan
```

This pipeline ensures infrastructure changes are validated, reviewed, and automatically deployed after merging to the main branch.