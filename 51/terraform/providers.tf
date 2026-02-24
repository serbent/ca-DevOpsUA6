terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.28.0"
    }
  }
  backend "s3" {
    bucket = "ca-devops-ua6-terraform-state"
    key    = "51/terraform.tfstate"
    region = "eu-central-1"
  }
}


provider "aws" {
  region = var.region
}
