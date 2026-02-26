# 52. Terraform state

```bash
terraform import aws_vpc.main-ua6-52 vpc-0c6cb5e969056d21e
terraform import aws_subnet.main-ua6-52 subnet-0feeadbcc562c35a2
terraform import aws_instance.main-ua6-52 i-03888f4859d1147a1

terraform graph -type=plan | dot -Tpng >graph.png
```