---
page_title: "smc_ospfv2_area"
subcategory: ""
description: |-
  This represents the OSPFv2 Area. It is a logical grouping of OSPFv2 routers that share the same area ID and routing information.
---

# smc_ospfv2_area

This represents the OSPFv2 Area. It is a logical grouping of OSPFv2 routers that share the same area ID and routing information.


## Simple Attributes
    
- `area_default_cost` (Number) The default cost for this OSPF area. This is used to determine the cost of routes in this area. If the area type is 'normal', this value is not applicable and should be null.
- `area_id` (Number) The unique identifier for the OSPF area.
- `area_type` (String) The type of the OSPF area.
- `comment` (String) An optional comment for the element. This field is not required.
- `inbound_filters_ref` (List of String) URI of the inbound filter access list.
- `interface_settings_ref` (String) This represents the abstract OSPF Interface Settings used as Dynamic Routing element. It contains settings related to OSPF interface configuration.
- `name` (String) Name of the object.
- `outbound_filters_ref` (List of String) URI of the outbound filter access list.
- `shortcut_capable_area` (Boolean) Indicates whether this OSPF area is capable of shortcut routing.

## Nested Attributes
    
- `ospf_abr_substitute_container` (List of Blocks, see [here](../attributes/ospf_abr_substitute_container.md)) The ABR substitute networks for this OSPF area. These are used to substitute routes in ABR areas.
- `ospfv2_virtual_links_endpoints_container` (List of Blocks, see [here](../attributes/ospfv2_virtual_links_end_points_container.md)) The OSPFv2 Virtual Links Endpoints Container belonging to the OSPFv2 Area.

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