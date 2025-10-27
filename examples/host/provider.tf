terraform {
  required_providers {
    smc = {
      source = "training-sop-fsmc/fhost"
      version = "0.0.2.1"
    }
  }
}

variable "api_key" {
  type        = string
  default     = null
}

variable "url" {
  type        = string
  default     = "http://localhost:8082"
}

variable "api_version" {
  type        = string
  default     = "7.4"
}

provider "smc" {
  url = var.url
  api_key = var.api_key
  api_version = var.api_version
}
