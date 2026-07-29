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
- **It is beteter for named resources**
- This solves the above issue and here **we create resources using key and not indexed**
- Example:
```
variable "instances" {
    default = [
        "dev",
        "stage",
        "prod"
    ]
}
resource "aws_instance" "server" {
    for_each = toset(var.instances)
    ami = var.ami
    instance_type ="t2.micro"

    tags = {
        Name = each.key
    } 
}
```
- This way terraform assigns as server["dev"], server["stage"], server["prod"] and not server[0], server[1], server[2]

| count                                | for_each                       |
| ------------------------------------ | ------------------------------ |
| Uses index                           | Uses key                       |
| Better for identical resources       | Better for named resources     |
| Index changes can recreate resources | Stable resource identity       |
| Uses `count.index`                   | Uses `each.key` / `each.value` |

- **RULE OF THUMB**
- Use count when you're creating N identical resources.
-Use for_each when each resource has its own identity (names, configurations, environments).