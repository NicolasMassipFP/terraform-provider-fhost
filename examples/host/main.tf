
resource "smc_host" "example" {
  name = "AExampleHost"
  address = "192.168.1.2"
  comment = "Created via Terraform"
  secondary = ["212.20.1.1", "123.6.5.19"]
}

output "example_admin_domain" {
  value = smc_host.example.admin_domain
}

output "example_key" {
  value = smc_host.example.key
}

output "example_link" {
  value = smc_host.example.link[0].href
}
