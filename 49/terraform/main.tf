## Using local modules

# module "virtual_machine" {
#   source = "./vm"
#   name   = "ca-ua6-49"
# }

# module "virtual_machine_2" {
#   source = "./vm"
#   name   = "ca-ua6-45"
# }

## Using remote modules
## Terraform Registry
# module "vpc" {
#   source  = "terraform-aws-modules/vpc/aws"
#   version = "6.2.0"
#   name    = "my-vpc"
#   cidr    = "10.0.0.0/16"

#   azs             = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
#   private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
#   public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

#   enable_nat_gateway = true
#   enable_vpn_gateway = false

#   tags = {
#     Terraform   = "true"
#     Environment = "dev"
#   }
# }


# module "ec2_instance" {
#   source  = "terraform-aws-modules/ec2-instance/aws"
#   version = "6.2.0"
#   name    = "single-instance"

#   instance_type = "t3.micro"
#   key_name      = "user1"
#   monitoring    = true
#   subnet_id     = module.vpc.public_subnets[0]

#   tags = {
#     Terraform   = "true"
#     Environment = "dev"
#   }
# }

## Terraform github 

module "github_instance" {
  source = "github.com/terraform-aws-modules/terraform-aws-ec2-instance?ref=v6.2.0"
  name   = "single-instance"

  instance_type = "t3.micro"
  key_name      = "user1"
  monitoring    = true
  subnet_id     = module.vpc.public_subnets[0]

  tags = {
    Terraform   = "true"
    Environment = "dev"
  }
}
