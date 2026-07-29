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