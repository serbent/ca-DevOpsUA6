
variable "instance_type" {
  type        = string
  default     = "t2.micro"
  description = "An instance type for our Virtual Machines"
}

variable "ami" {
  type        = string
  default     = "ami-0bae57ee7c4478e01"
  description = "An AMI for our Virtual Machines"
}

variable "subnet_id" {
  type        = string
  default     = "subnet-0c55b159cbfafe1f0"
  description = "A subnet for our Virtual Machines"
}

variable "tags" {
  type = map(string)
  default = {
    Name = "my-main-vm"
  }
  description = "Tags for our Virtual Machines"
}

variable "region" {
  type    = string
  default = "eu-central-1"
}

