module "virtual_machine" {
  source = "./vm"
  name   = "ca-ua6-49"
}

module "virtual_machine_2" {
  source = "./vm"
  name   = "ca-ua6-45"
}
