terraform {
  backend "s3" {
    bucket = "ca-devops-ua6-terraform-state"
    key    = "47/terraform.tfstate"
    region = "eu-central-1"
  }
}
