---
page_title: "inspected_service"
subcategory: ""
description: |-
  This represents a sub part of a Proxy Server, defining the protocol to be inspected and the port number used for inspection.
---

# inspected_service

This represents a sub part of a Proxy Server, defining the protocol to be inspected and the port number used for inspection.


## Simple Attributes
    
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `port` (Number) The port number used for the inspected service.
- `service_type` (String) The type of service being inspected, which can be 'FTP', 'HTTP', 'HTTPS', or 'SMTP'.


## Readonly Attributes
    
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](../attributes/api_link.md)) The API's links of the element, providing additional actions or resources.