---
page_title: "smc_route_map"
subcategory: "routing"
description: |-
  This represents a Route Map Policy for the Dynamic Routing Firewall settings, which is used to control the routing behavior based on specific rules.
---

# smc_route_map (Resource)

This represents a Route Map Policy for the Dynamic Routing Firewall settings, which is used to control the routing behavior based on specific rules.

## Examples

- [Route Map Example](https://github.com/Forcepoint/terraform-provider-fp-ngfw-smc/blob/release/0.0.1/examples/engines/dynamic_routing/routing_node_bgp/main.tf)

A `smc_route_map` provides route filtering or modification capabilities for imported and exported routes.

```hcl
resource "smc_route_map" "route_map" {
  name    = "tf_route_map"
  comment = var.resource_comment
}
```


## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.


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
