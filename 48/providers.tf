terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
  backend "s3" {
    bucket = "ca-devops-ua6-terraform-state"
    key    = "48/terraform.tfstate"
    region = "eu-central-1"
  }
}


provider "aws" {
  region = var.region
}
