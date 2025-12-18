terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "4.56.0"
    }
  }
}
provider "azurerm" {
  features {}
  subscription_id                 = "ca5f5073-c2fa-490f-aba8-c38e102a423c"
  resource_provider_registrations = "none"
}


locals {
  location = "polandcentral"
}

variable "name" {
  type    = string
  default = "ca-devopsua6-01"
}

resource "azurerm_resource_group" "ca-devopsua6-01" {
  name     = var.name
  location = local.location
}

resource "azurerm_virtual_network" "ca-devopsua6-01" {
  name                = "${var.name}-vnet"
  address_space       = ["10.0.0.0/16"]
  location            = local.location
  resource_group_name = azurerm_resource_group.ca-devopsua6-01.name
}

resource "azurerm_subnet" "ca-devopsua6-01" {
  name                                          = "${var.name}-subnet"
  resource_group_name                           = azurerm_resource_group.ca-devopsua6-01.name
  virtual_network_name                          = azurerm_virtual_network.ca-devopsua6-01.name
  address_prefixes                              = ["10.0.1.0/24"]
  private_link_service_network_policies_enabled = false
}

