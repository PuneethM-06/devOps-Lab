## TERRAFORM DAY 54

## WHAT IS TERRAFORM STATE 
- Terraform state is a file that records what infrastructure has been created and what is the curret state of the infrastructure 

### WHY DO WE NEED STATE?
- As we have known before, terraform doesnt execute a series of commands instead it compares with the current state of the infrastructure to the desired state of the infrastructure.

### cat terraform.tfstate gives idea n version, resources and output etc.

> what happens to the state when we do `terraform destroy`
- Answer: Terraform removes the destroyed resources from the state. The state is always kept in sync with the infrastructure it manages

### TERRAFORM.TFSTATE
- `terraform.tfstate` are not configuration file that we write
- It's actually a JSON file 
- Example:
```
{
  "version": 4,
  "terraform_version": "1.13.0",
  "serial": 8,
  "lineage": "xxxx",
  "outputs": {},
  "resources": []
}
```

1. **version** - This is the version of the state file format
2. **terraform_version** - which terraform version last updated this state 
3. **serial** - Everytime terraform changes infrastructure, serial increments by 1
4. **lineage** - This is the unique id of our terraform project
5. **outputs** - it stores the output generated after infratructure change
6. **resources**:
- This is the **heart of the tf state**
- It stores:
    - reource type
    - resource name 
    - ID
    - ARN
    - Metadata

### WHAT IS INFRASTRUCTURE DRIFT
- IT is a situation where he cloud infrastructure differs from the desired state defined in terraform, usually due to manual changes or changes made outside of terraform 

## Terraform Drift Detection

Terraform detects **Infrastructure Drift** when the actual infrastructure differs from the Terraform configuration.

### Example

Terraform configuration:

```hcl
tags = {
  Environment = "dev"
}
```

Someone manually changes the tag in the AWS Console:

```
Environment = production
```

Run:

```bash
terraform plan
```

Terraform detects the difference:

```
Desired : dev
Actual  : production
```

Run:

```bash
terraform apply
```

Terraform changes AWS back to:

```
Environment = dev
```

> **Infrastructure Drift** is the difference between the desired state (Terraform configuration) and the actual state (cloud infrastructure). Terraform detects drift with `terraform plan` and fixes it with `terraform apply`.

## REMOTE STATE BACKEND S3
- We know that currently `terraform.tfstate` is stored locally in our project folder called as **local backend**
- But this can be an issue when we are working in a team. Since each indiviual have their own .tfstate files 

### SOLUTION 
- To overcome this situation we can store the state in a **REMOTE BACKEND AND NOT LOCAL BACKEND**
- We can store it in a **AWS S3**
- By this way, everyone reads and writes to the same file 
### HOW DO WE SAY IT TO TF?
- By code:
```
terraform {
  backend "s3" {
    bucket = "my-terraform-state"
    key    = "terraform-lab/terraform.tfstate"
    region = "us-east-1"
  }
}
```
- And the next time we run terraform init, terraform now knows where to point to 
