---
page_title: "smc_ospfv2_profile"
subcategory: ""
description: |-
  This represents the OSPFv2 Profile for Dynamic Routing Firewall functionality. It is used to configure OSPFv2 settings in the firewall's dynamic routing capabilities.
---

# smc_ospfv2_profile

This represents the OSPFv2 Profile for Dynamic Routing Firewall functionality. It is used to configure OSPFv2 settings in the firewall's dynamic routing capabilities.


## Simple Attributes
    
- `comment` (String) An optional comment for the element. This field is not required.
- `default_metric` (Number) The default metric for the OSPF Profile. This value is used when no specific metric is defined for a route.
- `domain_settings_ref` (String) This represents the abstract OSPF Domain Settings used as Dynamic Routing element. It contains settings related to OSPF domain configuration.
- `external_distance` (Number) The external distance for the OSPF Profile. This value determines the cost of external routes.
- `inter_distance` (Number) The inter-area distance for the OSPF Profile. This value determines the cost of inter-area routes.
- `intra_distance` (Number) The intra-area distance for the OSPF Profile. This value determines the cost of intra-area routes.
- `name` (String) Name of the object.

## Nested Attributes
    
- `redistribution_entry` (Block) This represents a Dynamic Routing Redistribution Entry, which defines how routes are redistributed between different routing protocols or from the kernel, static, or connected routes.

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