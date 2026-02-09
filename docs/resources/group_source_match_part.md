---
page_title: "sources"
subcategory: ""
description: |-
  This represents the Source matching criteria for a Policy Rule, which can be 'any', 'none', or a list of specific source entries.
---

# sources

This represents the Source matching criteria for a Policy Rule, which can be 'any', 'none', or a list of specific source entries.


## Simple Attributes
    
- `any` (Boolean) Indicates if any source matches the criteria. If true, no specific elements are required.
- `networkDetails` (List of String) 
- `none` (Boolean) Indicates if no source matches the criteria. If true, no specific elements are required.
- `src` (List of String) URI of the source or match expression.

