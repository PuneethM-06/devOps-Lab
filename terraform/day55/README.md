## DAY 55 - TERRAFORM

## DATA SOURCES IN TERRAFORM
1. **resource** - resource is used when we want to create a new resource in AWS. Not make use of exisiting one 
- Example:
```
resource "aws_s3_bucket" "logs" {
    bucket = var.bucket_name
}
```