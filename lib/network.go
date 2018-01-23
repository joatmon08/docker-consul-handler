package lib

import (
	"context"

	"github.com/moby/moby/client"
)

// GetNetworkDetails retrieves network information from Docker networks
func GetNetworkDetails(networkID string) (string, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return "", err
	}
	network, err := cli.NetworkInspect(context.Background(), networkID)
	if err != nil {
		return "", err
	}

	return network.Name, nil
}
