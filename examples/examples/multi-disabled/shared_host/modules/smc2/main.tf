variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

variable "host_config" {
  description = "Shared host configuration"
  type        = any
}

resource "smc_host" "tf_shared_host_example" {
  name      = "${var.host_config.name} SMC2"
  address   = var.host_config.address
  comment   = var.resource_comment
  secondary = var.host_config.secondary
}

resource "smc_host" "tf_shared_host_example_smc2" {
  name      = "a_tf_host_example_defined_by_multiple_smc_instances"
  address   = "192.168.1.48"
  comment   = var.resource_comment
  secondary = ["212.24.1.1", "123.12.5.22"]
}
