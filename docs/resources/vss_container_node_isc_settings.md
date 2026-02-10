---
page_title: "vss_node_isc"
subcategory: "engines"
description: |-
  This represents the ISC Firewall Node Settings container, which includes management IP, gateway, hypervisor, and network settings.
---

# vss_node_isc

This represents the ISC Firewall Node Settings container, which includes management IP, gateway, hypervisor, and network settings.





## Simple Attributes
    
- `comment` (String) An optional comment for the element. This field is not required.
- `contact_ip` (String) The contact IP address for the ISC Firewall Node, used for communication with the SMC. If the contact IP is the same as the management IP, this field may be null.
- `isc_hypervisor` (String) The hypervisor type for the ISC Firewall Node, such as 'VMware' or 'Hyper-V'.
- `management_gateway` (String) The management gateway IP address for the ISC Firewall Node.
- `management_ip` (String) The management IP address of the ISC Firewall Node.
- `management_netmask` (Number) The netmask for the management network of the ISC Firewall Node.
- `name` (String) Name of the object.


## Readonly Attributes
    
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](zzattrs_api_link.md)) The API's links of the element, providing additional actions or resources.