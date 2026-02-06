---
page_title: "gateway_node"
subcategory: ""
description: |-
  This represents a gateway node, which is used to manage the VPN topology and its nodes, including their usage and relationships.
---

# gateway_node

This represents a gateway node, which is used to manage the VPN topology and its nodes, including their usage and relationships.


## Simple Attributes
    
- `child_node` (List of String) URI of the child Gateway Node.
- `comment` (String) An optional comment for the element. This field is not required.
- `gateway` (String) This is the base class for all storable elements.
- `name` (String) Name of the object.
- `node_usage` (String) The usage type of the gateway node, indicating its role in the VPN topology, such as 'central', 'satellite', or 'mobile'.
- `parent_node` (String) This represents a gateway node, which is used to manage the VPN topology and its nodes, including their usage and relationships.
- `vpn_key` (Number) The unique identifier for the VPN node, used to reference this node in the VPN topology.


## Readonly Attributes
    
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](../attributes/api_link.md)) The API's links of the element, providing additional actions or resources.