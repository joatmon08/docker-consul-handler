package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joatmon08/docker-consul-handler/lib"
	"github.com/joatmon08/docker-consul-handler/runner"
)

func main() {
	ansible_runner, ok := os.LookupEnv("CONSUL_HANDLER_RUNNER")
	if !ok {
		panic("Environment variable CONSUL_HANDER_RUNNER is not set!")
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	runner_client := runner.Client{URL: ansible_runner}

	for scanner.Scan() {
		input := scanner.Bytes()
		networkID, err := lib.GetNewestNetwork(input)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("Newest Network: %s \n", networkID)
		if len(networkID) > 0 {
			network, err := lib.GetNetworkDetails(networkID)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			if network == "bridge" {
				continue
			}
			fmt.Printf("Network ID %s, Network Name: %s \n", networkID, network)
			play_body := []byte(fmt.Sprintf("{\"container_network\": \"%s\"}", network))
			if _, err = runner_client.Play(play_body); err != nil {
				fmt.Printf("Network %s not created", network)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
