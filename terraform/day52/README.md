## TERRAFORM DAY 51

### INFRASTRUCTRE AS CODE
- Defining cloud resources through code is called IaC

## DECLARATIVE vs IMPERATIVE
1. **IMPERATIVE** 
- Here we tell **HOW TO DO IT**
- Example: Create a VPC, create a subnet, so on and etc

2. **DECLARATIVE**
- Here we tell **WHAT WE WANT** and we do not worry about how it is created
- Example: I want a VPC, 2 Ec2 instances etc.

- **IaC is declarative in nature**

## HOW TERRAFORM WORK:
```
main.tf

↓

terraform init

↓

Downloads Provider

↓

terraform plan

↓

Compares Desired State

↓

Shows Changes

↓

terraform apply

↓

Creates Infrastructure

↓

Updates terraform.tfstate
```

## PROVIDER

- By default, tf doesnt know about the provider and it needs to be defined explicitly as a plugin 

- Example:
provider "aws" {
    region = "us-east-1"
}

## RESOURCE
- Everything Terraform creates is a resource
- Example:
```
resource "aws_s3_bucket" "logs" {

}
```
- Here:
    - aws_s3_bucket is resource type
    - logs is the logical/resource name given 

## STATE FILE
- State files are used by terraform to **remember what was created by them**

## DESIRED STATE
- Terraform doesn't execute instructions one by one instead it tries to make the actual infrastructure match for the configuration we have written 
- Example: We write we need 2 Ec2 instances, and we have 1 EC2 already so when we do `terraform apply` it will create 1 more matching the desired state 

## TERRAFORM FMT
- It automatically formats our terraform configuration file according to official terraform style guide

## TERRAFORM VALIDATE
- validates if the terraform configuration i syntactically and structurally correct

## TERRAFORM INIT
- it initializes the current working directory as a terraform project
