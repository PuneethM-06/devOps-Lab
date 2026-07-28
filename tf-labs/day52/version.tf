terraform {
    required_version = ">=1.10"
}
required_providers {
    aws = {
        source = "hashicrop/aws"
        version = "~> 6.56.0"
    }
}