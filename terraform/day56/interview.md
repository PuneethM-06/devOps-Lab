# Terraform Interview Questions — Day 56
## Advanced Terraform (count, for_each, dynamic blocks, lifecycle, depends_on)

---

## 1. What are Terraform meta-arguments?

Meta-arguments are special arguments that control how Terraform manages resources rather than defining the resource itself.

Common meta-arguments include:

- count
- for_each
- depends_on
- lifecycle
- provider

These can be used with most Terraform resources.

---

## 2. What is `count`?

`count` creates multiple instances of the same resource.

Example:

```hcl
resource "aws_instance" "web" {
  count = 3

  ami           = var.ami
  instance_type = "t2.micro"
}
```

Terraform creates:

- aws_instance.web[0]
- aws_instance.web[1]
- aws_instance.web[2]

Each instance has its own unique AWS resource ID, IP addresses, and lifecycle.

---

## 3. What is `count.index`?

`count.index` returns the current iteration index while using `count`.

Example:

```hcl
tags = {
  Name = "server-${count.index}"
}
```

Result:

- server-0
- server-1
- server-2

---

## 4. What are the limitations of `count`?

`count` identifies resources by numeric indexes.

If the list order changes or an item is removed, Terraform may destroy and recreate resources because their indexes shift.

This makes `count` less suitable for resources with stable identities.

---

## 5. What is `for_each`?

`for_each` creates one resource for each element in a map or set.

Example:

```hcl
resource "aws_instance" "web" {
  for_each = toset(["dev", "staging", "prod"])

  tags = {
    Name = each.key
  }
}
```

Terraform creates:

- web["dev"]
- web["staging"]
- web["prod"]

---

## 6. When should you use `count` vs `for_each`?

Use `count` when:

- Creating identical resources
- Only the number of resources matters

Use `for_each` when:

- Resources have unique names or identities
- Individual resources should not be recreated due to index changes

---

## 7. What are `each.key` and `each.value`?

When iterating over a map:

```hcl
{
  web = "t2.micro"
  api = "t3.micro"
}
```

- `each.key` → web, api
- `each.value` → t2.micro, t3.micro

For a list/set of ports:

```hcl
[80, 443, 8000]
```

- `ingress.key` → 0, 1, 2
- `ingress.value` → 80, 443, 8000

---

## 8. What are dynamic blocks?

Dynamic blocks generate nested configuration blocks programmatically.

Example:

```hcl
dynamic "ingress" {
  for_each = var.ports

  content {
    from_port = ingress.value
    to_port   = ingress.value
    protocol  = "tcp"
  }
}
```

Instead of manually writing multiple ingress blocks, Terraform generates them automatically.

---

## 9. When are dynamic blocks useful?

They are useful when:

- Creating multiple ingress/egress rules
- Multiple EBS volumes
- Nested IAM policies
- Repeated nested configuration blocks

They help reduce duplicated code.

---

## 10. What is the `lifecycle` block?

The `lifecycle` block controls how Terraform manages a resource.

Common options include:

- prevent_destroy
- create_before_destroy
- ignore_changes

---

## 11. What does `prevent_destroy` do?

It prevents Terraform from deleting a resource.

Example:

```hcl
lifecycle {
  prevent_destroy = true
}
```

Useful for:

- Production databases
- Terraform state buckets
- Critical S3 buckets

---

## 12. What does `create_before_destroy` do?

Terraform creates the replacement resource before destroying the old one.

This helps minimize downtime during resource replacement.

---

## 13. What does `ignore_changes` do?

Terraform ignores changes to specific resource attributes after creation.

Example:

```hcl
lifecycle {
  ignore_changes = [
    tags
  ]
}
```

Useful when external systems modify resources.

---

## 14. What is `depends_on`?

`depends_on` creates an explicit dependency between resources when Terraform cannot infer one automatically.

Example:

```hcl
depends_on = [
  aws_iam_role_policy_attachment.attach
]
```

Terraform ensures the dependency is completed before creating the dependent resource.

---

## 15. When should `depends_on` be used?

Use it only when there is a real dependency that Terraform cannot detect through attribute references.

Avoid using it unnecessarily because Terraform already builds a dependency graph automatically.

---

## 16. How does Terraform determine resource creation order?

Terraform analyzes references between resources.

Example:

```hcl
vpc_id = aws_vpc.main.id
```

Terraform automatically knows the VPC must exist before creating the subnet.

This is called an implicit dependency.

---

## 17. Explain implicit vs explicit dependencies.

Implicit dependency:

Terraform detects it automatically through resource references.

Example:

```hcl
subnet_id = aws_subnet.public.id
```

Explicit dependency:

Defined manually using:

```hcl
depends_on = [...]
```

---

## 18. What are some best practices for advanced Terraform?

- Prefer `for_each` over `count` for named resources.
- Keep dynamic blocks readable.
- Use `prevent_destroy` for critical infrastructure.
- Use `create_before_destroy` for zero or low-downtime replacements.
- Avoid unnecessary `depends_on`.
- Let Terraform infer dependencies whenever possible.