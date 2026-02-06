---
page_title: "Forcepoint NGFW SMC - Terraform Provider"
description: Automate Forcepoint NGFW configuration deployment though SMC using terraform
---

# Forcepoint NGFW SMC - Terraform Provider

> SMC Provider is not available currently for version prior to 7.3.

> The current version of the SMC Terraform Provider does not officially support all SMC resources.
Please refer to the provider documentation to verify which resources are currently supported.

**PREVIEW WARNING**
```
/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\

This version is a pre release version for test purpose
  --- DO NOT USE UNLESS YOU ARE AWARE OF THE RISK ---

/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\/!\
```

## Which provider version to use ?

## 7.3.x 
Use version for SMC 7.4.x setup with SMC API 7.3

> **SMC 7.3.x does not officially support the SMC Terraform Provider**.
The Terraform provider released for **SMC 7.4.x**, which uses SMC API version 7.3, is mostly compatible with SMC 
7.3.x. However, there are breaking changes introduced in the SMC API between versions 7.3 and 7.4, which may lead to incompatibilities.

## 7.4.x 
Use latest version 1.74x.y

> default SMC API version used by latest provider 1.74x.y is SMC API 7.4.

## 7.5.x 
Use latest version 1.75x.y  

> default SMC API version used by latest provider 1.75x.y is SMC API 7.5.

> or latest version for 7.4.1
**SMC API 7.4 running on SMC 7.5 is backward compatible**.
As a result, the SMC Terraform provider for 7.4.x remains compatible.
However, using the provider version that matches the target SMC version is recommended to avoid any additional adaptation or compatibility handling.

## Recommended reading

The **guides folder** provides documentation to help you use the SMC Terraform provider.

# Schema

## Required

- `api_key` (String, Sensitive) The API key of the SMC API
- `url` (String) The URL of the SMC API

### Optional

- `api_version` (String) The API version to use for request if not the default one
- `trusted_cert` (String) PEM-encoded certificate content to trust for HTTPS connections.
- `verify_ssl` (Boolean) Whether to verify SSL certificates
