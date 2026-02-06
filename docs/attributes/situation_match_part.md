---
page_title: "situations"
subcategory: ""
description: |-
  This represents the Situation matching criteria for an IPS Policy Rule, which can specify whether any or none of the situations match, or list specific elements that match.
---

# situations

This represents the Situation matching criteria for an IPS Policy Rule, which can specify whether any or none of the situations match, or list specific elements that match.


## Simple Attributes
    
- `any` (Boolean) Indicates if any situation matches the criteria. If true, no elements are specified.
- `none` (Boolean) Indicates if no situation matches the criteria. If true, no elements are specified.
- `situation` (List of String) URI of the situation or match expression.

