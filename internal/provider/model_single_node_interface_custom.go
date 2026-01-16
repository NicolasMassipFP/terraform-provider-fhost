package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func (r *SingleNodeInterfaceResourceModel) GetSliceIds(ctx context.Context) []string {
	if !r.Dynamic.IsNull() && !r.Dynamic.IsUnknown() && r.Dynamic.ValueBool() {
		if !r.DynamicIndex.IsNull() && !r.DynamicIndex.IsUnknown() {
			return []string{fmt.Sprintf("%d", r.DynamicIndex.ValueInt64())}
		}

		if !r.DynamicIpv6Index.IsNull() && !r.DynamicIpv6Index.IsUnknown() {
			return []string{fmt.Sprintf("%d", r.DynamicIpv6Index.ValueInt64())}
		}
		tflog.Error(ctx, "Dynamic is set but neither DynamicIndex nor DynamicIpv6Index is set")
		return nil
	}
	if !r.Address.IsNull() && !r.Address.IsUnknown() {
		return []string{r.Address.ValueString()}
	}

	return nil
}
