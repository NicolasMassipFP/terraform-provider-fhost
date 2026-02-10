---
page_title: "smc_internal_user"
subcategory: "admin"
description: |-
  This represents an Internal User. It contains user details such as display name, days left until expiration, user groups, password, and pre-shared key.
---

# smc_internal_user

This represents an Internal User. It contains user details such as display name, days left until expiration, user groups, password, and pre-shared key.





## Simple Attributes
    
- `activation_date` (Number) The optional activation date in milliseconds since epoch.
- `authentication_method` (List of String) URI of the authentication method.
- `comment` (String) An optional comment for the element. This field is not required.
- `days_left` (Number) The number of days left until the user expires or is activated.
- `display_name` (String) The display name of the user.
- `expiration_after` (Number) The optional expiration delay in days. If not set, the user will never expire.
- `expiration_date` (Number) The optional expiration date in milliseconds since epoch. If not set, the user will never expire.
- `name` (String) Name of the object.
- `password` (String) The password of the user in case of password authentication method.
- `pre_shared_key` (String) The pre-shared key of the user in case of pre-shared key authentication method.
- `subject_alt_names` (String) The subject alternative names required in case of IPSec certificate authentication method. This is a comma-separated list of subject alternative names.
- `unique_id` (String) The unique id of the User/User Group element.
- `user_domain` (String) This represents a User Domain, which is used to define the authentication domain for users. It can be either an authentication domain or the internal domain.
- `user_group` (List of String) URI of the user group.


## Readonly Attributes
    
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](zzattrs_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.