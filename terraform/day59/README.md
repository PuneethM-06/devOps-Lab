## DAY 59 - TERRAFORM 

## IMPORT AND PRODUCTION PRACTICES
- `terraform import` gets the current state of infrastructure or services running in your cloud
- And once there are resources identified which was not created by terraform is not noted and is now managed by terraform 
- General syntax
```
terraform import <resource-address> <resource-id>
terraform import aws_s3_bucket.main my-existing-bucket
```
