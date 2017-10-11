package main

import (
	"fmt"
	"net/http"
	"testing"
)

func addToConsulDockerNetworkKey(networkID string) error {
	url := fmt.Sprintf("http://127.0.0.1:8500/v1/kv/docker/network/v1.0/endpoint/%s/", networkID)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	if _, err = client.Do(req); err != nil {
		return err
	}
	return nil
}

func TestShouldAddTwoNetworks(t *testing.T) {
	if err := addToConsulDockerNetworkKey("newest"); err != nil {
		t.Errorf("%s", err.Error())
	}
	if err := addToConsulDockerNetworkKey("newer"); err != nil {
		t.Errorf("%s", err.Error())
	}
}
