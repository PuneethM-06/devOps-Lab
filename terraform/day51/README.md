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
- By default, tf doesnt know about the provider and it needs to be defined explicitly as a plugin 
- Exaple:
```
provider "aws" {
    region = "us-east-1
}
```
