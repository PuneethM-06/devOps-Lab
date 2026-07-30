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
- Use for_each when each resource has its own identity (names, configurations, environments).

5. ## DYNAMIC BLOCKS
- Normally
```
resource "aws_security_group" "web" {
    ingress {
        from_port = 80
    }
    ingress {
        from_port = 443
    }
}
```
- this is too repetitive

```
variable "ports" {
    default = [80,443,8000]
}

resource "aws_security_group" "web" {

    dynamic "ingress" {
        for_each = var.ports

        content {
            from_port = ingress.value
            to_port = ingress.value
            protocol = "tcp"
            cidr_blocks = ["0.0.0.0/0"]
        }
    }
}
```
- In a dynamic block, Terraform iterates over the collection and gives you an iterator named after the block (ingress in this case).
| Iteration | `ingress.key` | `ingress.value` |
| --------- | ------------- | --------------- |
| 1         | `0`           | `80`            |
| 2         | `1`           | `443`           |
| 3         | `2`           | `8000`          |
- hence ingress.value gives the port numbers

## LIFECYCLE
- Lifecycle tells how to manage resources and not what resources to create 
```
resources "aws_s3_bucket" "logs" {
    bucket = var.bucket_name

    lifecycle {

    }
}
```

1. ### prevent_destroy
- This ensures that `terraform destroy` isn't applied
- Example:
```
resources "aws_s3_bucket" "logs" {
    bucket = var.bucket_name

    lifecycle {
        prevent_destroy = true
    }
}
```

2. ### create_before_detroy
- This is **Extremely Useful**
- Lets say we make a change that cannot make updates in an running Ec2 instances. Instead it needs to create a new EC2 instance
- Normally what Terraform does
```
Can't modify AMI

↓

Must replace instance

Destroy old EC2

↓

Create new EC2
```
- this means, downtime and server lost
- Instead
```
lifecycle {
    create_before_destroy = true
}
```
```
Create new server

↓

Wait until ready

↓

Delete old server
``` 

3. ### ignore_changes
- this means ignore changes of future modification to a specified attribute
```
lifecycle {

    ignore_changes = [
        desired_capacity
    ]

}
```

4. ### depends_on
- terraform builds a dependency graph 
```
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_subnet" "public" {
  vpc_id = aws_vpc.main.id
}
```
- Terraform knows `subnet` uses `vpc` and hence it waits until VPC is created
- This is called **implicit dependency**

- But we can also explicitly define it as well
```
resource "aws_iam_role_policy_attachment" "attach" {
  role       = aws_iam_role.role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
}

resource "aws_instance" "web" {
  ami           = var.ami
  instance_type = "t2.micro"

  depends_on = [
    aws_iam_role_policy_attachment.attach
  ]
}
```
- Here we are telling, do not create EC2 instances unless the policy is created 
