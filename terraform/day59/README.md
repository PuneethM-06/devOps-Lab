## DAY 59 - TERRAFORM 

## IMPORT AND PRODUCTION PRACTICES
- `terraform import` gets the current state of infrastructure or services running in your cloud
- And once there are resources identified which was not created by terraform is not noted and is now managed by terraform 
- General syntax
```
terraform import <resource-address> <resource-id>
terraform import aws_s3_bucket.main my-existing-bucket
```
1. **STEP 1 - RESOURCE EXISTS OUTSIDE TERRAFORM** 
2. **STEP 2 - WRITE THE TERRAFORM CODE**
3. **STEP 3 - IMPORT THE RESOURCE** 
- After this terraform starts tracking that resource

1. ### terraform state list
- List all resources that are tracked in the terraform state
- It does not query AWS, it reads from the state list 

2. ### terraform state show 
- Displays all the attributes of one resources stored in terraform state 
- Unlike, `terraform state list` it displays all the needed information about a specific resource

3. ### terraform state mv
- It moves the resources within the terraform state without making changes in the infrastructure 
- It does not:
    1. Create resources
    2. Delete resources
    3. Modify AWS
- Example:
```
Suppose I initially wrote 
resource "aws_s3_bucket" "bucket" {
  bucket = "company-logs"
}

Later, I want to rename it to
resource "aws_s3_bucket" "logs" {
  bucket = "company-logs"
}

Instead of deleting and creating the same bucket again, I can simply do:
terraform state mv aws_s3_bucket.bucket aws_s3_bucket.logs
```