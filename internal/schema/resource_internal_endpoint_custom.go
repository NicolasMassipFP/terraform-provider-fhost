package schema

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func (r *InternalEndpointResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.AtLeastOneOf(
			path.MatchRoot("ssl_vpn_portal"),
			path.MatchRoot("ssl_vpn_tunnel"),
			path.MatchRoot("ipsec_vpn"),
		),
	}
}
