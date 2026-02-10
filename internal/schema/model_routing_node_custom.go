package schema

import (
	"context"
)

func (r *RoutingNodeResourceModel) GetSliceIds(ctx context.Context) []string {
	var ids []string
	if (!r.Href.IsNull()) && (!r.Href.IsUnknown()) {
		ids = append(ids, r.Href.ValueString())
	}
	if (!r.Name.IsNull()) && (!r.Name.IsUnknown()) {
		ids = append(ids, r.Name.ValueString())
	}
	if len(ids) != 0 {
		return ids
	}
	return nil
}
