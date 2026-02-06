---
page_title: "web_site_access_situation_parameter_value"
subcategory: ""
description: |-
  This represents a web site access parameter value within a situation, allowing for the application of specific web site access rules to the situation's parameters.
---

# web_site_access_situation_parameter_value

This represents a web site access parameter value within a situation, allowing for the application of specific web site access rules to the situation's parameters.


## Simple Attributes
    
- `comment` (String) An optional comment for the element. This field is not required.
- `dfas` (List of String) URI of the associated DFA file.
- `name` (String) Name of the object.
- `order` (Number) The order of this parameter value within the situation, determining its sequence.
- `parameter_ref` (String) This represents a parameter within a situation, allowing for the configuration of specific parameters that can be used in the context of the situation.
- `regularExpression` (String) The regular expression used in this parameter value, which defines the pattern to match.


## Readonly Attributes
    
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](../attributes/api_link.md)) The API's links of the element, providing additional actions or resources.