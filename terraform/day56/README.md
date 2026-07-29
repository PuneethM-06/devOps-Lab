## DAY 56 - ADVANCED TERRAFORM

1. ## COUNT 
- This command is used to create identical resources
- Example:
```
reource "aws_instance" "server" {
    count = 3

    ami = var.ami
    instance_type ="t2.micro" 
}
```
- It can be addressed as `aws_instance.server[0].id`

2. ## COUNT INDEX
- Useful insisde a counted instance
- Example:
```
resource "aws_instance" "server" {
    count = 3

    tags = {
        Name  = "server-${count.index}"
    }
}
```

3. ## PROBLEM WITH COUNT 
- Suppose if we redice `count = 2` then terraform destroys server[2], but if we had server like "dev, stage and prod" then it would be difficult

4. ## FOR EACH
