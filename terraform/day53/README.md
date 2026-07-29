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

