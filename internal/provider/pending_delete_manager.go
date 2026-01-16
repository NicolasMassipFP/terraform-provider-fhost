package provider

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-providers/terraform-provider-smc/internal/datastore"
	"github.com/terraform-providers/terraform-provider-smc/internal/smc"
	"net/http"
)

// returns the number of successfully deleted pending resources
func deletePendingResourcesInternal(ctx context.Context) (int, error) {
	pendingDeletes, err := datastore.GetPendingDeletes(ctx)
	if err != nil {
		return 0, err
	}

	deleteSuccessCount := 0
	for _, pendingDelete := range pendingDeletes {
		tflog.Debug(ctx, fmt.Sprintf("PendingDeleteManager: processing pending delete %s", pendingDelete.Href))
		smcClient := smc.GetClientManager().GetClient(
			ctx, pendingDelete.BaseUrl, false)
		if smcClient == nil {
			tflog.Debug(ctx, fmt.Sprintf("PendingDeleteManager: no client found for pending delete %s", pendingDelete.BaseUrl))
			continue
		}

		restClient := smc.GenericCRUDConfig{
			ResourceType: pendingDelete.ResourceType,
			Client:       smcClient,
		}

		readResp, err := restClient.ReadResourceByHref(pendingDelete.Href)
		if err != nil {
			var statusErr *smc.StatusError
			if errors.As(err, &statusErr) && statusErr.Code == http.StatusNotFound {
				tflog.Debug(ctx, fmt.Sprintf(
					"PendingDeleteManager: %s not found, assuming already deleted", pendingDelete.Href))
			} else {
				// actual error, we'll retry another time
				continue
			}
		} else {
			err = smc.DeleteResourceByHref(&restClient,
				pendingDelete.Href, (*readResp).ETag)
			if err != nil {
				tflog.Debug(ctx, fmt.Sprintf(
					"PendingDeleteManager: error deleting %s: %s",
					pendingDelete.Href, err.Error()))
				continue
			}
		}

		tflog.Debug(ctx, fmt.Sprintf(
			"PendingDeleteManager: %s successfully deleted", pendingDelete.Href))
		err = datastore.ManagePendingDelete(ctx,
			datastore.RemovePendingDelete, "", "", pendingDelete.Href)
		if err != nil {
			return 0, err
		}
		deleteSuccessCount++
	}

	return deleteSuccessCount, nil
}

func DeletePendingResources(ctx context.Context) {
	deleteSuccessCount := 0
	for ok := true; ok; ok = (deleteSuccessCount > 0) {
		var err error
		deleteSuccessCount, err = deletePendingResourcesInternal(ctx)
		if err != nil {
			tflog.Error(ctx, fmt.Sprintf("PendingDeleteManager: error deleting pending resources: %s", err.Error()))
			return
		}
		tflog.Debug(ctx, fmt.Sprintf("PendingDeleteManager: deleted %d pending resources in this pass", deleteSuccessCount))
	}
}
