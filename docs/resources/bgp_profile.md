---
page_title: "smc_bgp_profile"
subcategory: "routing"
description: |-
  This represents the BGP Profile for Dynamic Routing Firewall functionality, including port settings, distances, and BGP entries.
---

# smc_bgp_profile (Resource)

This represents the BGP Profile for Dynamic Routing Firewall functionality, including port settings, distances, and BGP entries.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `external` (Number) The External distance value, which should be between 1 and 255.
- `internal` (Number) The Internal distance value, which should be between 1 and 255.
- `local` (Number) The Local distance value, which should be between 1 and 255.
- `name` (String) Name of the object.
- `port` (Number) The BGP port value.

## Nested Attributes
- `aggregation_entry` (List of Blocks, see [here](attr_bgp_aggregation_entry.md)) A list of BGP aggregation entries, which define the aggregation of BGP routes.
- `bmp_entry` (List of Blocks, see [here](attr_bgp_bmp_entry.md)) A list of BGP BMP entries, which define the BGP Monitoring Protocol entries.
- `distance_entry` (List of Blocks, see [here](attr_bgp_distance_entry.md)) A list of BGP distance entries, which define the distance for specific BGP routes.
- `redistribution_entry` (List of Blocks, see [here](attr_dynamic_routing_redistribution_entry.md)) A list of BGP Redistribution entries, which define how routes are redistributed in BGP.

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
