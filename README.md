# SMC Terraform Provider

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://www.datocms-assets.com/2885/1629941242-logo-terraform-main.svg" width="600px">

## Which provider version to use ?

Provider is not available currently for version prior to 7.3.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x +
- [Go](https://golang.org/doc/install) 1.21.x (to build the provider plugin)
- The provider can cover SMC >= 7.3.X versions, the configuration of all parameters should be based on the relevant SMC documentation. For SMC, the support on this provider is deprecated, please use [SMC Terraform provider](https://registry.terraform.io/providers/smcdev/smc/latest) instead.

## Building the Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using "make":

## Using the Provider

Internal version for testing purpose only.

## Developing the Provider

This provider is managed by Forcepoint R&D internally. The public repositoryis updated by Forcepoint when releasing new SMC Version or on provider update.

You can still compile locally the provider using ``make build`` command.
The compilation rely on docker container that will be retrieved and generate locally.

For your information, the provider implementation is generated from the Open API specification of SMC API.
