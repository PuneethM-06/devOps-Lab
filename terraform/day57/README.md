## DAY 57 - TERRAFORM 

## WHY DO WE NEED MULTIPLE ENVIRONMENT
1. To ensure each env is isolated in time of a incident 
2. dev -> developer and for testing 
3. stage -> near to prod 
4. prod -> user facing 

### PROBLEM
- Lets say `dev` needs `instance_type` of `t2.micro,` `stage` needs `t2.small` and `prod` needs `t2.large.`
- Having a seperate code for each of these are repetative and hence we make use of **workspace, variable files(.tfvars), backend seperation and resuable modules**

## WHAT IS A WORKSPACE
- A workspace lets you use the same **terraform configuration file while maintaining seperaye state files**

## CREATING AND MANAGING TERRAFORM WORKSPACE
- `terraform workspace list` indicates the currently active workspace

1. ### CREATE A NEW WORSPACE
- `terraform workspace new dev`
- It creates a new workspace called as dev and is switched to it as well 

2. ### SWITCH WORKSPACES
- `terraform workspace select prod`

3. ### DELETE A WORKSPACE
- `terraform workspace delete prod`
