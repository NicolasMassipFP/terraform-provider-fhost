variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

resource "smc_extended_community_access_list" "extended_community_access_list" {
  comment = ""
  entries {
    extended_community_access_list_entry {
      action = "permit"
      community = ".*1545"
    }
  }
  name = "tf_extended_community_access_list"
  type = "expanded"
  comment = var.resource_comment
}