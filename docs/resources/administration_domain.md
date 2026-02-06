---
page_title: "smc_admin_domain"
subcategory: ""
description: |-
  This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
---

# smc_admin_domain

This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.


## Simple Attributes
    
- `announcement_enabled` (Boolean) Indicates whether the announcement message is enabled. If true, the announcement message will be displayed before the login window.
- `announcement_message` (String) The announcement message to be displayed before the login window. This message can be used to communicate important information to users.
- `category_filter_system` (Boolean) Indicates whether system elements should be shown in the domain. If true, system elements will be included in the domain's configuration.
- `comment` (String) An optional comment for the element. This field is not required.
- `contact_email` (String) The contact email address for the domain. This email can be used for communication related to the domain.
- `contact_number` (String) The contact phone number for the domain. This number can be used for communication related to the domain.
- `logo_ref` (String) This is a Logo File. It represents a logo file used in the system, typically for branding or identification purposes.
- `name` (String) Name of the object.
- `show_not_categorized` (Boolean) Indicates whether uncategorized elements should be shown in the domain. If true, elements that are not categorized will be included in the domain's configuration.

## Nested Attributes
    
- `user_alert_check` (List of Blocks, see [here](../attributes/user_alert_check_association.md)) The list of user alert checks associated with the domain. These checks are used to monitor and manage user alerts within the domain.

## Readonly Attributes
    
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](../attributes/api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.