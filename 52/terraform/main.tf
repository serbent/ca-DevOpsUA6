resource "aws_vpc" "main-ua6-52" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name        = "main-ua6-52"
    environment = "test"
  }
}

resource "aws_subnet" "main-ua6-52" {
  vpc_id     = aws_vpc.main-ua6-52.id
  cidr_block = "10.0.0.0/24"
  tags = {
    Name        = "main-ua6-52"
    environment = "test"
  }
}

resource "aws_instance" "main-ua6-52" {
  instance_type = "t2.micro"
  ami           = "ami-0bae57ee7c4478e01"
  subnet_id     = aws_subnet.main-ua6-52.id
  tags = {
    Name        = "my_-main-vm-ua6-52"
    environment = "test"
  }
}
