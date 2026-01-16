package smc

import (
	"testing"
)

func TestIsVpnUrl(t *testing.T) {
	url := "http://localhost:8082/7.4/elements/vpn/268438672/gateway_tree_nodes/central"
	prefix := isVpnUrl(url)
	if prefix != "http://localhost:8082/7.4/elements/vpn/268438672" {
		t.Errorf("expected vpn url prefix, got: %s", prefix)
	}

	prefix = isVpnUrl("http://localhost:8082/7.4/elements/single_fw/268571118")
	if prefix != "" {
		t.Errorf("expected empty string for non-vpn url, got: %s", prefix)
	}
}
