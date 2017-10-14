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
		network, err := lib.GetNewestNetwork(input)
		if len(network) > 0 {
			if err != nil {
				panic(err)
			}
			play_body := []byte(fmt.Sprintf("{\"container_network\": \"%s\"}", network))
			fmt.Printf("Body %s", play_body)
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
