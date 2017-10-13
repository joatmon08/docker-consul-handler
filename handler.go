package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/joatmon08/docker-consul-handler/lib"
	"github.com/joatmon08/docker-consul-handler/runner"
)

func main() {
	filename := flag.String("filename", "/scripts/consul_watch.log", "file path")
	ansible_runner := flag.String("runner", "127.0.0.1", "URL of Ansible Playbook Runner")

	flag.Parse()

	f, err := os.OpenFile(*filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	runner_client := runner.Client{URL: *ansible_runner}

	for scanner.Scan() {
		input := scanner.Bytes()
		if len(input) > 1 {
			network, err := lib.GetNewestNetwork(input)
			if err != nil {
				panic(err)
			}

			if _, err = f.WriteString(network + "\n"); err != nil {
				panic(err)
			}
			string_play_body := fmt.Sprintf("{\"container_network\": \"%s\"}", network)
			fmt.Println(string_play_body)
			play_body := []byte(string_play_body)

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
