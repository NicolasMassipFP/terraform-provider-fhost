---
page_title: "gateway_certificate_request"
subcategory: ""
description: |-
  This represents a certificate request for an internal gateway, allowing for the generation of unsigned certificates.
---

# gateway_certificate_request

This represents a certificate request for an internal gateway, allowing for the generation of unsigned certificates.


## Simple Attributes
    
- `comment` (String) An optional comment for the element. This field is not required.
- `key_length` (Number) The length of the key used in the certificate request, typically measured in bits.
- `name` (String) Name of the object.
- `public_key_algorithm` (String) The algorithm used for the public key in the certificate request, such as RSA or ECDSA.
- `request` (String) The certificate request in PEM format, which is used to request a certificate from a Certificate Authority.
- `signature_algorithm` (String) The algorithm used to sign the certificate request, which ensures the integrity and authenticity of the request.


## Readonly Attributes
    
- `key` (Number) The unique identifier for the element. This field is required for updates but not for creation.
- `link` (List of Blocks, see [here](../attributes/api_link.md)) The API's links of the element, providing additional actions or resources.