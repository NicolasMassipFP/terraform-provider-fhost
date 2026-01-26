# ERROR
# Error: Unsupported block type
#
# on main.tf line 15, in resource "smc_ospfv2_profile" "ospfv2_profile":
# 15:   redistribution_entry {
variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_ospfv2_profile" "ospfv2_profile" {
  domain_settings_ref = "http://localhost:18082/7.4/elements/ospfv2_domain_settings/2"
  external_distance   = 120
  inter_distance      = 120
  intra_distance      = 120
  name                = "tf_ospfv2_profile"
  redistribution_entry {
    enabled                   = true
    filter_type               = "access_list"
    metric                    = 120
    metric_type               = "external_1"
    redistribution_filter_ref = "http://localhost:18082/7.4/elements/ip_access_list/29"
    type                      = "kernel"
  }
  redistribution_entry {
    enabled     = false
    filter_type = "none"
    metric_type = "external_1"
    type        = "static"
  }
  redistribution_entry {
    enabled     = true
    filter_type = "none"
    metric_type = "external_1"
    type        = "connected"
  }
  redistribution_entry {
    enabled                   = true
    filter_type               = "access_list"
    metric                    = 170
    metric_type               = "external_1"
    redistribution_filter_ref = "http://localhost:18082/7.4/elements/ip_access_list/29"
    type                      = "bgp"
  }
  redistribution_entry {
    enabled     = false
    filter_type = "none"
    metric_type = "external_1"
    type        = "default_originate"
  }
  comment = var.resource_comment
}
