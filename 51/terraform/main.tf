locals {
  short_env_name = {
    development = "dev"
    staging     = "stage"
    production  = "prod"
  }
  instance_type = {
    development = "t3.micro"
    staging     = "t3.small"
    production  = "t3.medium"
  }
  short_env    = lookup(local.short_env_name, var.environment)
  intance_size = lookup(local.instance_type, var.environment)
  # if production 3 instances else 1 instance
  intance_count  = var.environment == "production" ? 3 : 1
  private_subnet = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
}




resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "main"
  }
}


resource "aws_subnet" "main" {
  for_each   = toset(local.private_subnet)
  vpc_id     = aws_vpc.main.id
  cidr_block = each.value
  tags = {
    Name = "private-${each.key}"
  }
}

resource "aws_instance" "main" {
  count         = local.intance_count
  ami           = "ami-0bae57ee7c4478e01"
  instance_type = local.intance_size
  tags = {
    Name        = "ca-devops-ua6-51-${local.short_env}"
    Environment = var.environment
    Env         = local.short_env
  }
  subnet_id = aws_subnet.main["10.0.1.0/24"].id

  lifecycle {
    create_before_destroy = true
    prevent_destroy       = false
    ignore_changes        = [tags]
  }
}

resource "aws_s3_bucket" "main" {
  bucket = "ca-devops-ua6-51-${local.short_env}"
  tags = {
    Name        = "ca-devops-ua6-51-${local.short_env}"
    Environment = var.environment
    Env         = local.short_env
  }
  depends_on = [aws_instance.main[0]]
}
