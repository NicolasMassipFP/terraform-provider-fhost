variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

variable "shared_host_config" {
  description = "Shared host configuration for both SMC instances"
  type = object({
    name        = string
    description = string
    address     = string
    secondary   = optional(list(string))
  })

  default = {
    name        = "shared-host"
    description = "Host used by both SMC instances"
    address     = "192.168.1.100"
    secondary   = ["212.20.1.1", "123.6.5.22"]
    tags = {
      environment = "production"
    }
  }
}

module "smc1_module" {
  source = "./modules/smc1"
  providers = {
    smc = smc.smc1
  }
  host_config = var.shared_host_config
}

module "smc2_module" {
  source = "./modules/smc2"
  providers = {
    smc = smc.smc2
  }
  host_config = var.shared_host_config
}
