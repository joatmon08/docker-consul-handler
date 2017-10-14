package tests

import (
	"testing"
	"github.com/joatmon08/docker-consul-handler/lib"
)

func TestNetworkShouldReturnName(t *testing.T) {
	networkName, err := lib.GetNetworkDetails("8aa5119445007ca6ce569b956805d4a4564cc713c8b277c1690d92d1832b7137")
	if err != nil {
		t.Error(err.Error)
	}
	if networkName != "bridge" {
		t.Errorf("Expected bridge, got %s", networkName)
	}
}
