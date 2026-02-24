output "instance_id" {
  value = aws_instance.my_vm.id
}

output "instance_public_ip" {
  value       = aws_instance.my_vm.public_ip
  description = "Public IP address of the EC2 instance"
  sensitive   = false
}
