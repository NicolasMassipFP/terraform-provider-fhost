variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "ssh_service" {
  name = "SSH"
  type = "tcp_service"
}

data "smc_href" "https_service" {
  name = "HTTPS"
  type = "tcp_service"
}

data "smc_href" "default_policy_template" {
  type = "fw_template_policy"
  name = "Firewall Template"
}

resource "smc_fw_policy" "example" {
  name     = "tf_example_access_rules_policy2"
  template = data.smc_href.default_policy_template.id
  comment  = var.resource_comment
}

resource "smc_fw_ipv4_access_rule" "allow_ssh" {
  from_ref = smc_fw_ipv4_access_rule.allow_https.lk.add_after
  name     = "allow-ssh-from-any-to-any"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.ssh_service.id]
  }
  action { action = ["allow"] }
}

resource "smc_fw_ipv4_access_rule" "allow_https" {
  from_ref = smc_fw_policy.example.lk.fw_ipv4_access_rules
  name     = "allow-https-from-any-to-any"
  sources { any = true }
  destinations { any = true }
  services {
    service = [data.smc_href.https_service.id]
  }
  action { action = ["allow"] }
}
