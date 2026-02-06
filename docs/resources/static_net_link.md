---
page_title: "smc_netlink"
subcategory: ""
description: |-
  This represents a Static NetLink, which is a type of NetLink used for routing in Multi-Link features. It includes attributes for gateway, networks, DNS elements, and outbound IP addresses.
---

# smc_netlink

This represents a Static NetLink, which is a type of NetLink used for routing in Multi-Link features. It includes attributes for gateway, networks, DNS elements, and outbound IP addresses.


## Simple Attributes
    
- `active_mode_period` (Number) The active period in milliseconds for the NetLink, which defines how often the link is probed when it is in Active or Standby mode. Leave the setting for Standby Mode as 0 if you prefer not to test this link when it is on standby.
- `active_mode_timeout` (Number) The active timeout in milliseconds for the NetLink, which defines how long the firewall waits before it considers the probe failed. Change the setting for Standby Mode to 0 if you prefer not to test this link when it is on standby.
- `comment` (String) An optional comment for the element. This field is not required.
- `connection_type_ref` (String) This represents a connection type used in endpoints, which defines the connectivity group, mode, and link type for VPN connections.
- `gateway_ref` (String) This represents a network element, which is a component that has an IP address and can be part of a network. It includes a location reference.
- `input_speed` (Number) The Input Speed in bits per second for the NetLink, which defines the real-life bandwidth this network connection provides. It is used to calculate how much traffic each link receives in relation to the other links.
- `ipv4_outbound` (String) The IP address used to NAT the IPv4 outbound traffic. Defaults to the CVI address defined in the routing view.
- `ipv6_outbound` (String) The IP address used to NAT the IPv6 outbound traffic. Defaults to the CVI address defined in the routing view.
- `location_ref` (String) This represents the definition of a Location, which keeps a list of Network Elements belonging to the same location.
- `name` (String) Name of the object.
- `network_ref` (List of String) URI of the Network to define the address space.
- `nsp_name` (String) The NSP Name for the NetLink, which is the provider name of your ISP.
- `output_speed` (Number) The Output Speed in bits per second for the NetLink, which defines the real-life bandwidth this network connection provides. It is used to calculate how much traffic each link receives in relation to the other links.
- `probe_address` (List of String) The IP address that is probed with ICMP echo requests (ping) to determine if the link is up. Repeat this for each IP address you want to add. We recommend entering more than one address to avoid excluding the link in case the host that is probed goes down.
- `standby_mode_period` (Number) The standby period in seconds for the NetLink, which defines how often the link is probed when it is in Standby mode. Leave the setting for Standby Mode as 0 if you prefer not to test this link when it is on standby.
- `standby_mode_timeout` (Number) The standby timeout in seconds for the NetLink, which defines how long the firewall waits before it considers the probe failed in Standby mode. Change the setting for Standby Mode to 0 if you prefer not to test this link when it is on standby.

## Nested Attributes
    
- `domain_server_address` (List of Blocks, see [here](../attributes/dns_element.md)) The list of DNS elements that define the DNS servers provided by the ISP. Each element contains the DNS server address and other related information.

## Readonly Attributes
    
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](../attributes/api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.