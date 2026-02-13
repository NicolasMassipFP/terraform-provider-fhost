---
page_title: "smc_external_gateway"
subcategory: "vpn"
description: |-
  This represents an External Gateway, which is used for establishing VPN connections with external networks. It includes gateway profile, trust settings for VPN certificate authorities, and lists of end-points and sites.
---

# smc_external_gateway (Resource)

This represents an External Gateway, which is used for establishing VPN connections with external networks. It includes gateway profile, trust settings for VPN certificate authorities, and lists of end-points and sites.




## Simple Attributes
- `id` (String) this attribute is the identifier of terraform resource
- `comment` (String) An optional comment for the element. This field is not required.
- `end_point` (List of String) URI of the external end-point.
- `gateway_profile` (String) Gateway Profiles describe the capabilities of a Gateway, i.e. supported cipher, hash, etc. Gateway Profiles of Internal Gateways are read-only and computed from Firewall version and FIPS mode. Gateway Profiles of External Gateways are user-defined.
- `name` (String) Name of the object.
- `site` (List of String) URI of the site.
- `trust_all_cas` (Boolean) Indicates whether the External Gateway trusts all VPN Certificate Authorities. If true, it does not require specific trusted CAs.
- `trusted_certificate_authorities` (List of String) URI of the trusted VPN Certificate Authority.


## Readonly Attributes
- `admin_domain` (String) This represents a Domain. Domains are administrative boundaries that allow you to separate the configuration details and other information in the system for the purpose of limiting administrator access.
- `etag` (String) The ETag of the element, used for versioning. This field is not required.
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](attr_api_link.md)) The API's links of the element, providing additional actions or resources.
- `locked` (Boolean) Indicates if the element is locked. This field is not required.
- `read_only` (Boolean) Indicates if the element is read-only. This field is not required.
- `system` (Boolean) Indicates if the element is a System element. This field is not required.
- `system_key` (Number) The system key of the System element. This field is not required.
- `trashed` (Boolean) Indicates if the element is trashed. This field is not required.
