variable "location" {
  type    = string
  default = "polandcentral"
}

variable "resource_group_name" {
  type    = string
  default = "rg-container-app"
}

variable "container_app_name" {
  type    = string
  default = "hello-container-app"
}

variable "container_app_image" {
  type    = string
  default = "serbentas/ca-devopsua6:0.0.1"
}

variable "container_app_env_name" {
  type    = string
  default = "cae-example"
}

variable "log_analytics_workspace_name" {
  type    = string
  default = "law-container-app"
}
