---
page_title: "configuration_situation_parameter_value"
subcategory: "situations"
description: |-
  This represents a configuration parameter value within a situation, allowing for the application of specific configuration values to the situation's parameters.
---

# configuration_situation_parameter_value

This represents a configuration parameter value within a situation, allowing for the application of specific configuration values to the situation's parameters.





## Simple Attributes
    
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `order` (Number) The order of this parameter value within the situation, determining its sequence.
- `parameter_ref` (String) This represents a parameter within a situation, allowing for the configuration of specific parameters that can be used in the context of the situation.


## Readonly Attributes
    
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](zzattrs_api_link.md)) The API's links of the element, providing additional actions or resources.