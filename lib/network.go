package lib

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetNetworkDetails(networkID string) (string, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return "", err
	}
	network, err := cli.NetworkInspect(context.Background(), networkID, types.NetworkInspectOptions{})
	if err != nil {
		return "", err
	}

	return network.Name, nil
}
