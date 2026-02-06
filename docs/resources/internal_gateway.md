---
page_title: "internal_gateway"
subcategory: ""
description: |-
  This represents the Internal Gateway, which is used for managing VPN connections and related settings.
---

# internal_gateway

This represents the Internal Gateway, which is used for managing VPN connections and related settings.


## Simple Attributes
    
- `antivirus` (Boolean) Indicates whether the AntiVirus is enabled for this Internal Gateway.
- `auto_certificate` (Boolean) Indicates whether the Internal Gateway automatically generates and manages RSA certificates.
- `auto_site_content` (Boolean) Indicates whether the site content is automatically generated from the routing view. This is applicable only for Internal Gateways.
- `cluster_ref` (String) This represents a group of devices, or nodes, that share a given work load. You can cluster Firewalls to share the load and provide redundancy, allowing, for example, scheduled maintenance that takes one node out of service without interrupting services to the users.
- `comment` (String) An optional comment for the element. This field is not required.
- `dtls` (Boolean) 
- `end_point` (List of String) URI of the internal end-point.
- `firewall` (Boolean) Indicates whether the Firewall is enabled for this Internal Gateway.
- `gateway_profile` (String) Gateway Profiles describe the capabilities of a Gateway, i.e. supported cipher, hash, etc. Gateway Profiles of Internal Gateways are read-only and computed from Firewall version and FIPS mode. Gateway Profiles of External Gateways are user-defined.
- `name` (String) Name of the object.
- `site` (List of String) URI of the site.
- `trust_all_cas` (Boolean) Indicates whether the EndPoint trusts all VPN Certificate Authorities. If true, it trusts all CAs; if false, it requires specific trusted CAs.
- `trusted_certificate_authorities` (List of String) URI of the trusted VPN Certificate Authority.
- `vpn_client_mode` (String) The VPN Client Mode for this Internal Gateway, which can be 'no', 'ipsec', 'ssl', or 'both'.
- `windows-update` (Boolean) Indicates whether Windows Update is enabled for this Internal Gateway.

## Nested Attributes
    
- `dhcp_relay` (Single Block, see [here](../attributes/dhcp_client_configuration.md)) 
- `ssl_vpn_portal_setting` (List of Blocks, see [here](../attributes/ssl_vpn_portal_setting.md)) The Application Access Portal Settings for this Internal Gateway.
- `ssl_vpn_proxy` (Single Block, see [here](../attributes/ssl_vpn_setting.md)) 
- `ssl_vpn_tunneling` (Single Block, see [here](../attributes/ssl_vpn_setting.md)) 

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