---
page_title: "smc_vpn"
subcategory: "vpn"
description: |-
  This represents a Policy Based Virtual Private Network (VPN), which is used to establish secure connections over unsecured networks. It includes various configurations such as NAT rules, mobile VPN topology modes, and associated profiles.
---

# smc_vpn (Resource)

This represents a Policy Based Virtual Private Network (VPN), which is used to establish secure connections over unsecured networks. It includes various configurations such as NAT rules, mobile VPN topology modes, and associated profiles.

## Examples

- [two_gateway_in_one_engine_disabled/main.tf](https://github.com/Forcepoint/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/sdwan/policy_based_vpn/two_gateway_in_one_engine_disabled/main.tf): Defines a VPN using a custom VPN profile without NAT or mobile VPN topology.

This snippet shows basic usage of the `smc_vpn` resource in a policy-based VPN context, including essential attributes such as profile reference, topology mode, and NAT configuration. The VPN resource is used as a parent for gateway nodes in SD-WAN deployments.

```hcl
resource "smc_vpn" "tf_sample_vpn" {
  mobile_vpn_topology_mode = "None"
  name                     = "tf_sample_vpn"
  nat                      = false
  vpn_profile              = data.smc_href.suite_b_gcm_128.href
  comment                  = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `dscp_qos_policy` (String) This represents a QoS Policy, which is used for Bandwidth Management and Traffic Prioritization based on QoS Classes or DSCP Matches.
- `link_usage_profile` (String) This represents a Link Usage Profile, which is used to manage link usage settings in a multi-link environment. It includes configurations for link balancing, packet duplication, and forward erasure correction.
- `mobile_vpn_topology_mode` (String) Indicates the mobile VPN topology mode, which defines the type of gateway communicating with the VPN client.
- `name` (String) Name of the object.
- `nat` (Boolean) Indicates whether NAT rules are applied for this VPN.
- `vpn_profile` (String) This represents a VPN Profile. It contains settings for IKE and IPsec lifetimes, keep-alive options, certificate authorities, and authentication methods.


## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
