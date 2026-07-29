## TERRAFORM - DAY53

## TERRAFORM VARIABLES
- Terraform variables is a **nmaed input that allow you to pass values into our terraform configuration files** instead of hardcoding them
- Instead of :
```
resource "aws_s3_bucket" "logs" {
    bucket = "aws-s3-logs"
}
```
- We can do:
```
resource "aws_s3_bucket" "logs" {
    bucket = var.bucket_name
}
```
- Variables are declared inside the variables block and is defined as:
```
variable "bucket_name" {
    description =  "Name of the s3 bucket"
    type = string 
}
```
## VARIABLE TYPES
1. **string** - stores text
- Example: "us-east-1"

2. **Number** - stores numeric values
- Example: instance_count

3. **Boolean** - stores eithet true or false
- Example: enable_versioning = true

4. **List** - A list is a ordered collection of values of the same type
- Example: 
``` 
availability_zones = [
  "us-east-1a",
  "us-east-1b",
  "us-east-1c"
]
Here it is Map<String, String>
```
4. **Map** - stores key-value
- Example:
```
variable "tags" {
  type = map(string)
}

tags = {
  Environment = "Dev"
  Owner       = "Puneeth"
  Team        = "Platform"
}
```
5. **Object**: - stores attributes with different datatype
- Example:
```
variable "server" {
  type = object({
    name   = string
    cpu    = number
    public = bool
  })
}

server = {
  name   = "web-server"
  cpu    = 2
  public = true
}
```

## DEFAULT VALUES
- A default value is a value that terraform automatically uses if the user does not provide one 

- Example:
```
variable "region" {
    description = "aws_region"
    type = string 
    default = "us-east-1"
}
```
## REQURIED VARIABLES
- A required variable is a variable that must be provided by the user becaue it has no default value 
- We basically **do not include default**

- Example:
```
variable "bucket_name" {
    description =  "s3"
    type = string 
}
```
## SENSITIVE VARIABLE
- A sensitive variable stores confidential information that terraform hides from its output
- Example:
```
variable "db_password" {
    description = db pwd
    type = string 
    sensitive = true
}
```
#### NOTE:
- `sensitive = true` does not encrypt the pwd instead it hides from CLI output 

## TERRAFORM.TFVARS

- `terraform.tfvars` is a file that provides value for terraform input variables
- Meaning, in the below code block example. we gave it only the variable name and we never declared it 
```
variable "bucket_name" {
    description = "bucket_name"
    type = string
}
```
- In terraform.tfvars:
```
bucket_name = "kryptxx-demo-bucket-3627"
```
and **we can have mulitple `.tfvars` file ** and we can choose one while planning
- Example:
```
dev.tfvars

test.tfvars

prod.tfvars
```
```
terraform plan -var-file="dev.tfvars"
or 
terraform applu -var-file="dev.tfvars"
```

| `variables.tf`               | `terraform.tfvars`            |
| ---------------------------- | ----------------------------- |
| Declares variables           | Assigns values                |
| Defines type and description | Stores actual values          |
| Used by developers           | Often changed per environment |
| Usually stays the same       | Changes for Dev/Test/Prod     |

## Precedence (Highest → Lowest):
```
-var
-var-file
*.auto.tfvars
terraform.tfvars
TF_VAR_* environment variables
default values
```
