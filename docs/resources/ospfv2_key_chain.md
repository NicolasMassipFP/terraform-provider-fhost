---
page_title: "smc_ospfv2_key_chain"
subcategory: ""
description: |-
  This represents the OSPFv2 Key Chain element used as Message Digest authentication method for OSPFv2 Interface Settings for Dynamic Routing Firewall functionality.
---

# smc_ospfv2_key_chain

This represents the OSPFv2 Key Chain element used as Message Digest authentication method for OSPFv2 Interface Settings for Dynamic Routing Firewall functionality.


## Simple Attributes
    
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.

## Nested Attributes
    
- `ospfv2_key_chain_entry` (List of Blocks, see [here](../attributes/ospf_key_chain_entry.md)) The OSPFv2 Key Chain Entries, which define the keys used for authentication in the OSPFv2 Key Chain.

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