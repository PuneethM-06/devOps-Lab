## DAY 55 - TERRAFORM

## DATA SOURCES IN TERRAFORM
- A data source in terraform allows terraform to read already exisiting resources in cloud and not create from first

1. **resource** - resource is used when we want to create a new resource in AWS. Not make use of exisiting one 
- Example:
```
resource "aws_s3_bucket" "logs" {
    bucket = var.bucket_name
}
```

2. **data** - data is used when we want to make use the exisiting resource. **Read something that already exists in AWS**
- Example:
```
data "aws_s3_bucker" "demo" {
    bucket = "kryptxx-demo-bucket"
}   
```
- The above code snippet says, go to AWS and fetch information about this bucket 

## MODULES
- Modules are terraform configurations that are written oce and is reused anywhere needed 
- A Terraform module is simply a collective of terraform files that work together to perform a specific task 
```
EC2 Module
├── main.tf
├── variables.tf
└── outputs.tf
```
- The above modules task is to create EC2 instances only 

- A company might organize terraform files like 
```
terraform/

modules/
│
├── networking/
├── compute/
└── storage/

environments/
├── dev/
├── test/
└── prod/
```
## ROOT MODULE vs CHILD MODULE
- Here **keyword is module**
- Example:
```
module "web_server" {
    source = "./modules/ec2"
}
```
- Here `module` tells terraform I'm resusing another terraform configurations
- `source` tells terraform where to find the module.

## Root Module vs Child Module

### Project Structure

```text
terraform/
│
├── modules/
│   ├── networking/
│   ├── compute/
│   ├── database/
│   └── storage/
│
├── dev/
├── stage/
└── prod/
```

### Child Modules
- Live inside `modules/`
- Reusable building blocks
- One responsibility (Networking, Compute, Storage, etc.)
- Know **how** to create infrastructure

### Root Modules
- Usually environment folders (`dev/`, `stage/`, `prod/`)
- Entry point of the deployment
- Call child modules using `module` blocks
- Decide **what** infrastructure to deploy

### Flow

```text
terraform apply (inside prod/)
        │
        ▼
    Root Module
        │
        ├──► Networking Module → Creates VPC
        ├──► Compute Module    → Creates EC2
        ├──► Database Module   → Creates RDS
        └──► Storage Module    → Creates S3
```

### Easy Way to Remember

> **Root Module = Manager (orchestrates the deployment)**
>
> **Child Module = Worker (implements one reusable piece of infrastructure)**

### Interview Definition

> **The Root Module is the entry point where Terraform commands are executed. It orchestrates the deployment by calling Child Modules. Child Modules are reusable Terraform configurations that implement a specific infrastructure component.**

> How do modules communicate?
- Modules communicate through the root module. One module exposes outputs, and the root module passes those outputs as inputs to another module.

> What are module inputs?
- Module inputs are variable that are defined in the child module 

```
Networking Module
        │
        │ Output
        ▼
     vpc_id
        │
        ▼
   Root Module
        │
        │ Input
        ▼
 Compute Module
 ```
## Module Communication (Inputs & Outputs)

Child Modules communicate through the Root Module.

### Step 1 - Networking Module

The Networking module creates a VPC.

```hcl
resource "aws_vpc" "main" {
  cidr_block = var.cidr_block
}
```

After creating the VPC, it exposes the VPC ID.

```hcl
output "vpc_id" {
  value = aws_vpc.main.id
}
```

Think of it as:

```
Networking Module

Creates VPC
      │
      ▼
Returns:
vpc_id = vpc-12345
```

---

### Step 2 - Root Module

The Root Module calls the Networking module.

```hcl
module "networking" {
  source = "../modules/networking"

  cidr_block = "10.0.0.0/16"
}
```

Terraform now knows:

```
module.networking.vpc_id

↓

vpc-12345
```

The Root Module then passes this value to another module.

```hcl
module "compute" {
  source = "../modules/compute"

  vpc_id = module.networking.vpc_id
}
```

Read this line as:

> Take the `vpc_id` output from the Networking Module and pass it as the `vpc_id` input to the Compute Module.

---

### Step 3 - Compute Module

The Compute module doesn't create or know anything about the VPC.

It simply accepts a VPC ID as an input.

```hcl
variable "vpc_id" {
  type = string
}
```

It uses that value while creating resources.

```hcl
resource "aws_security_group" "web" {
  name   = "web-sg"
  vpc_id = var.vpc_id
}
```

---
### Communication Flow

```
             Root Module

        module "networking"
                │
                ▼
      Creates VPC
                │
                ▼
        Output: vpc_id
                │
                │ module.networking.vpc_id
                ▼
        module "compute"
                │
                ▼
       Input: var.vpc_id
                │
                ▼
 Creates Security Group inside the VPC
```

```

> Child Modules never communicate directly. The Root Module passes outputs from one module as inputs to another.