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

}
```
- Here `module` tells terraform I'm resusing another terraform configurations

