---
page_title: "vpn_site"
subcategory: ""
description: |-
  This represents a VPN Site, which is a collection of network elements that are part of a VPN. It can be associated with a gateway and has references to VPNs.
---

# vpn_site

This represents a VPN Site, which is a collection of network elements that are part of a VPN. It can be associated with a gateway and has references to VPNs.


## Simple Attributes
    
- `automatic` (Boolean) Indicates whether this site is automatically created or manually added. If true, the site is automatically created and cannot have VPN references set.
- `comment` (String) An optional comment for the element. This field is not required.
- `gateway` (String) This represents a VPN Gateway. A VPN Gateway is a network element that can be used to establish a VPN connection with other gateways.
- `name` (String) Name of the object.
- `site_element` (List of String) The URIs of network element that are part of this VPN Site. These elements can be routers, switches, or other network devices.

## Nested Attributes
    
- `vpn_references` (List of Blocks, see [here](../attributes/site_vpn_reference.md)) The VPN references associated with this site, which define the VPNs that this site is part of. These references link the site to specific VPN configurations.

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