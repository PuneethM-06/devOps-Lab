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

Child Modules do **not** communicate directly. The **Root Module** acts as the bridge by passing outputs from one module as inputs to another.

### Step 1 - Networking Module

Creates a VPC and exposes its ID.

```hcl
resource "aws_vpc" "main" {
  cidr_block = var.cidr_block
}

output "vpc_id" {
  value = aws_vpc.main.id
}
```

After creation:

```text
Output:
vpc_id = vpc-12345
```

---

### Step 2 - Root Module

The Root Module calls the Networking Module and uses its output as the input for the Compute Module.

```hcl
module "networking" {
  source = "../modules/networking"

  cidr_block = "10.0.0.0/16"
}

module "compute" {
  source = "../modules/compute"

  vpc_id = module.networking.vpc_id
}
```

Here,

```hcl
module.networking.vpc_id
```

means:

> Get the `vpc_id` output from the Networking Module and pass it to the Compute Module.

---

### Step 3 - Compute Module

The Compute Module simply accepts the VPC ID and uses it.

```hcl
variable "vpc_id" {}

resource "aws_security_group" "web" {
  name   = "web-sg"
  vpc_id = var.vpc_id
}
```

---

### Communication Flow

```text
Networking Module
        │
        │ Output: vpc_id
        ▼
     Root Module
        │
        │ Input: module.networking.vpc_id
        ▼
   Compute Module
        │
        ▼
Creates Security Group inside the VPC
```

### Key Takeaway

> Child Modules expose values using **outputs**. The Root Module receives those outputs and passes them as **inputs** to other Child Modules.