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