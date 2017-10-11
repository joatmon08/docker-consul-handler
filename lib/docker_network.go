package lib

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Network is a struct
type Network struct {
	ID   string
	Name string
}

func GetNetworkDetails() ([]Network, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return nil, err
	}

	var network_details []Network
	for _, network := range networks {
		network_details = append(network_details, Network{ID: network.ID, Name: network.Name})
	}

	return network_details, nil
}
