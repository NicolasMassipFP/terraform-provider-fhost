# SMC Terraform Provider - fp-ngfw-smc

- Website: https://registry.terraform.io/providers/forcepoint/fp-ngfw-smc
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

## Which provider version to use ?

> SMC Provider is not available currently for version prior to 7.3.

| SMC Version | SMC Provider recommended        |
| ----------- | ------------------------------- |
|    7.3.x    | Use latest fp-ngfw-smc terraform version 1.74x.y **setup with SMC API 7.3** (*)  |
|    7.4.x    | Use latest fp-ngfw-smc terraform version 1.74x.y      |
|    7.5.x    | Use latest fp-ngfw-smc terraform version  -- or latest fp-ngfw-smc terraform version 1.74x.y **setup with SMC API 7.5** (**) |

> (*) **SMC 7.3.x does not officially support the SMC Terraform Provider**.
The Terraform provider released for **SMC 7.4.x**, which uses SMC API version 7.3, is mostly compatible with SMC 
7.3.x. However, there are breaking changes introduced in the SMC API between versions 7.3 and 7.4, which may lead to incompatibilities.

> (**) **SMC API 7.4 running on SMC 7.5 is backward compatible**.
As a result, the SMC Terraform provider for 7.4.x remains compatible.
However, using the provider version that matches the target SMC version is recommended to avoid any additional adaptation or compatibility handling.

> The current version of the SMC Terraform Provider does not officially support all SMC resources.
Please refer to the provider documentation to verify which resources are currently supported.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.14.x +
- [Go](https://golang.org/doc/install) 1.25.x
- [Go releaser](https://goreleaser.com/install/) 2.13.x

## Building the Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using "make":

## Using the Provider

## Provider management and versioning
This provider is internally maintained by Forcepoint R&D.
The public repository is updated by Forcepoint when:

- a new SMC version is released (only if required), or
- a new provider version is published.

You can still build the provider locally using the following command:
```shell
make build
```
The build process relies on a Docker container, which will be automatically pulled and used locally during compilation.

## OpenAPI model and version mapping
The provider implementation is generated from the SMC API OpenAPI specification.
The exact SMC API model version used is reflected in the provider version number:

For provider version M.XYZ.P, the OpenAPI model comes from SMC version X.Y.Z.

Example:
- Provider version 1.741.0
- Uses the OpenAPI model from SMC version 7.4.1

> Starting with SMC 7.4.1, the SMC API OpenAPI model is included in the SMC package.
