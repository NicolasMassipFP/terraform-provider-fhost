package schema

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func (r *IcapServerResource) Metadata(
	_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_icap_server"
}
