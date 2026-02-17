package schema

import (
	"github.com/terraform-providers/terraform-provider-smc/internal/apijson"
)

func (r *SingleFirewallResource) ModelToJson(data *SingleFirewallResourceModel) ([]byte, error) {
	// the API expects an empty array instead of null for
	// WebAuthentication.EnabledInterface (error 400, parameter1 is null)
	if data.WebAuthentication != nil && data.WebAuthentication.EnabledInterface == nil {
		data.WebAuthentication.EnabledInterface = &[]EnabledInterfaceEntryResourceModel{}
	}
	return apijson.MarshalRoot(*data)
}
