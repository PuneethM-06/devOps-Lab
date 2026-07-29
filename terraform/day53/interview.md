# Day 53 – Terraform Variables, Locals, Outputs & Expressions

---

## 1. What is a Terraform variable?

**Answer:**
A variable is an input to a Terraform configuration that allows values to be provided externally instead of hardcoding them.

---

## 2. Why should we use variables?

**Answer:**
- Reusable configurations
- Environment-specific values
- Easier maintenance
- Avoid hardcoding

---

## 3. What is the syntax to reference a variable?

```hcl
var.bucket_name
```

---

## 4. What is terraform.tfvars?

**Answer:**
A file that provides values for variables declared in `variables.tf`.

Terraform automatically loads it.

---

## 5. Difference between variables.tf and terraform.tfvars?

| variables.tf | terraform.tfvars |
|--------------|------------------|
| Declares variables | Assigns values |
| Defines type | Defines actual data |
| Usually doesn't change | Changes per environment |

---

## 6. What is a local?

**Answer:**
A local is a reusable value defined inside Terraform.

Locals help avoid duplication.

---

## 7. Difference between Variables and Locals?

| Variables | Locals |
|------------|---------|
| External input | Internal reusable value |
| `var.` | `local.` |

---

## 8. Why use locals?

**Answer:**
- Reduce duplicate code
- Store common tags
- Improve readability
- Easier maintenance

---

## 9. How do you reference a local?

```hcl
local.common_tags
```

---

## 10. What is an Output?

**Answer:**
Outputs expose useful information after Terraform creates infrastructure.

Examples:
- Bucket ARN
- Bucket Name
- EC2 Public IP
- Database Endpoint

---

## 11. How do you display outputs?

```bash
terraform output
```

Single output

```bash
terraform output bucket_name
```

---

## 12. Can outputs be sensitive?

**Answer:**
Yes.

```hcl
sensitive = true
```

---

## 13. What is a Terraform expression?

**Answer:**
Any Terraform code that evaluates to a value.

Examples:

```hcl
var.bucket_name
```

```hcl
local.common_tags
```

```hcl
aws_s3_bucket.demo.arn
```

```hcl
upper(var.bucket_name)
```

---

## 14. What are common Terraform expressions?

- Variable references
- Local references
- Resource references
- Conditional expressions
- Function calls

---

## 15. Give an example of a conditional expression.

```hcl
var.environment == "prod" ? "t3.large" : "t3.micro"
```

---

## 16. Name some commonly used Terraform functions.

- upper()
- lower()
- length()
- join()
- split()
- lookup()

---

## 17. Explain the Terraform project structure.

```
versions.tf
providers.tf
variables.tf
terraform.tfvars
locals.tf
main.tf
outputs.tf
```

---

## 18. Why split Terraform code into multiple files?

**Answer:**

- Better readability
- Easier maintenance
- Separation of concerns
- Production best practice

---

## 19. What is the flow of a Terraform project?

```
terraform.tfvars
        │
        ▼
variables.tf
        │
        ▼
locals.tf
        │
        ▼
main.tf
        │
        ▼
Terraform creates resources
        │
        ▼
outputs.tf
```

---

## 20. Scenario

**Question:**
Your company wants the same tags on 100 AWS resources. How would you implement it?

**Answer:**
Store the tags in `locals.tf` and reference them using `local.common_tags`.

---

## 21. Scenario

**Question:**
The bucket name changes for each environment. How would you implement it?

**Answer:**
Declare a variable in `variables.tf`, assign values in `terraform.tfvars`, and reference it using `var.bucket_name`.

---

## 22. Scenario

**Question:**
How would you retrieve an S3 bucket ARN after deployment?

**Answer:**
Create an output:

```hcl
output "bucket_arn" {
  value = aws_s3_bucket.demo.arn
}
```

Retrieve it with:

```bash
terraform output bucket_arn
```

---

# Quick Revision

- Variable → External input
- terraform.tfvars → Variable values
- Local → Internal reusable value
- Expression → Code that evaluates to a value
- Output → Information returned after deployment