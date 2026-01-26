package provider

import (
	"context"
)

func UseHCL2(ctx context.Context) bool {
	useHcl2_val := ctx.Value("use_hcl2")
	if useHcl2_val == nil {
		return false
	}
	useHcl2, ok := useHcl2_val.(bool)
	if !ok {
		return false
	}
	return useHcl2
}
