package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
)

// close and save the vpn before deleting it
func (r *PolicyBasedVpnResource) beforeDelete(
	ctx context.Context, _ *PolicyBasedVpnResourceModel, id string) error {

	mgr := smc.GetVpnEditStateManager()
	state, ok := mgr.VpnEditStates.Load(id)
	if !ok {
		tflog.Debug(ctx, fmt.Sprintf("vpn %s not open for edit, no need to close", id))
		return nil
	}
	tflog.Debug(ctx, fmt.Sprintf("Need to close vpn %s", id))

	// best effort to save and close. Ignore errors
	err := smc.SendVpnCommand(state.(*smc.VpnEditState).SmcClient, id, "save")
	if err != nil {
		tflog.Warn(ctx,
			fmt.Sprintf("Failed to save vpn %s: %s", id, err.Error()))
	}
	err = smc.SendVpnCommand(state.(*smc.VpnEditState).SmcClient, id, "close")
	if err != nil {
		tflog.Error(ctx,
			fmt.Sprintf("Failed to close vpn %s: %s", id, err.Error()))
	}
	mgr.VpnEditStates.Delete(id)
	return nil
}
