---
page_title: "url_list_situation_parameter_value"
subcategory: ""
description: |-
  This represents a URL list parameter value within a situation, allowing for the application of specific URL lists to the situation's parameters.
---

# url_list_situation_parameter_value

This represents a URL list parameter value within a situation, allowing for the application of specific URL lists to the situation's parameters.


## Simple Attributes
    
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.
- `order` (Number) The order of this parameter value within the situation, determining its sequence.
- `parameter_ref` (String) This represents a parameter within a situation, allowing for the configuration of specific parameters that can be used in the context of the situation.
- `urls` (List of String) The list of URLs associated with this parameter value.


## Readonly Attributes
    
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](../attributes/api_link.md)) The API's links of the element, providing additional actions or resources.