provider "aws" {
  region = "eu-central-1"
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.0"
    }
  }
  backend "s3" {
    bucket       = "ca-devops-ua6-terraform-state"
    key          = "52/terraform.tfstate"
    region       = "eu-central-1"
    use_lockfile = true
  }
}
