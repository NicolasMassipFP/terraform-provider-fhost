variable "resource_comment" {
  type    = string
  default = "Created by Terraform"
}

data "smc_href" "log_server" {
  type = "log_server"
  name = "*"
}

data "smc_href" "switch_module" {
  type = "appliance_switch_module"
  name = "115"
}

resource "smc_single_fw" "tf_single_fw" {
  name           = "tf_single_fw_switch_interface"
  log_server_ref = data.smc_href.log_server.href
  comment        = var.resource_comment
  nodes {
    firewall_node {
      name   = "myfwnode"
      nodeid = 1
    }
  }

  physical_interfaces {
    physical_interface {
      interface_id = 0
      interfaces {
        single_node_interface {
          nodeid        = 1
          nicid         = 0
          address       = "192.168.100.14"
          network_value = "192.168.100.00/24"
          primary_mgt   = true
        }
      }
    }
  }
  physical_interfaces {
    switch_interface {
      interface_id = "SWI_0"
      name         = "Switch 0"
      port_group_interface {
        interface_id = "SWI_0.1"
        interfaces {
          single_node_interface {
            address       = "fc00:100e::12"
            network_value = "fc00:100e::/64"
            nicid         = "SWI_0.1"
            nodeid        = 1
          }
        }
        name = "Port Group 0.1 (ports 0-2)"
        switch_interface_port {
          physical_switch_port_number   = "0"
          switch_interface_port_comment = ""
        }
        switch_interface_port {
          physical_switch_port_number   = "1"
          switch_interface_port_comment = ""
        }
        switch_interface_port {
          physical_switch_port_number   = "2"
          switch_interface_port_comment = ""
        }
      }
      port_group_interface {
        interface_id = "SWI_0.2"
        interfaces {
          single_node_interface {
            address       = "192.168.15.25"
            network_value = "192.168.15.0/24"
            nicid         = "SWI_0.2"
            nodeid        = 1
          }
        }
        name = "Port Group 0.2 (ports 4-6)"
        switch_interface_port {
          physical_switch_port_number   = "4"
          switch_interface_port_comment = ""
        }
        switch_interface_port {
          physical_switch_port_number   = "5"
          switch_interface_port_comment = ""
        }
        switch_interface_port {
          physical_switch_port_number   = "6"
          switch_interface_port_comment = ""
        }
      }
      switch_interface_switch_module_ref = data.smc_href.switch_module.id
    }
  }

}
