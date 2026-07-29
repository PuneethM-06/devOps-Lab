## TERRAFORM - DAY53

## TERRAFORM VARIABLES
- Terraform variables is a nmaed input that allow you to pass values into our terraform configuration files instead of hardcoding them
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
