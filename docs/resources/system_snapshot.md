---
page_title: "smc_system_snapshot"
subcategory: ""
description: |-
  This represents a System Snapshot, which is a record of system elements defined within a specific update package. It helps in tracking changes after the activation of an update package.
---

# smc_system_snapshot

This represents a System Snapshot, which is a record of system elements defined within a specific update package. It helps in tracking changes after the activation of an update package.


## Simple Attributes
    
- `comment` (String) An optional comment for the element. This field is not required.
- `name` (String) Name of the object.


## Readonly Attributes
    
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](../attributes/api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `packageId` (Number) The ID of the package associated with this system snapshot.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.