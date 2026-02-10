terraform {
  required_providers {
    smc = {
      source  = "forcepoint/fp-ngfw-smc"
      version = "0.0.1"
    }
  }
}

variable "api_key1" {
  type    = string
  default = null
}

variable "url1" {
  type    = string
  default = "http://192.168.56.1:8082"
}

variable "api_version1" {
  type    = string
  default = "7.4"
}

variable "api_key2" {
  type    = string
  default = null
}

variable "url2" {
  type    = string
  default = "http://192.168.56.101:8082"
}

variable "api_version2" {
  type    = string
  default = "7.4"
}

provider "smc" {
  alias       = "smc1"
  url         = var.url1
  api_key     = var.api_key1
  api_version = var.api_version1
}


provider "smc" {
  alias       = "smc2"
  url         = var.url2
  api_key     = var.api_key2
  api_version = var.api_version2
}
