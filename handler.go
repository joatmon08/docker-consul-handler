package main

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/joatmon08/docker-consul-handler/lib"
	"github.com/joatmon08/docker-consul-handler/runner"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	ansibleRunner, ok := os.LookupEnv("CONSUL_HANDLER_RUNNER")
	if !ok {
		log.Panic("Set CONSUL_HANDLER_RUNNER as environment variable.")
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	runnerClient := runner.Client{URL: ansibleRunner}

	network := &lib.Network{}

	for scanner.Scan() {
		input := scanner.Bytes()
		log.WithFields(log.Fields{
			"networkID":   network.ID,
			"createIndex": network.CreateIndex,
		}).Info("Last network added.")

		if len(input) == 0 {
			continue
		}

		isNew, err := network.HasNewNetwork(input)
		if err != nil {
			log.WithFields(log.Fields{
				"message": err.Error(),
			}).Error("Failed to find newest network.")
			continue
		}
		if !isNew {
			log.Info("No new networks added.")
			continue
		}

		log.WithFields(log.Fields{
			"networkID":   network.ID,
			"createIndex": network.CreateIndex,
		}).Info("Found new network.")
		if len(network.ID) > 0 {
			networkName, err := lib.GetNetworkDetails(network.ID)
			if err != nil {
				log.WithFields(log.Fields{
					"networkID": network.ID,
					"message":   err.Error(),
				}).Error("Could not retrieve network name.")
				continue
			}
			if networkName == "bridge" {
				continue
			}
			log.WithFields(log.Fields{
				"networkID":   network.ID,
				"networkName": networkName,
			}).Info("Adding new network with runner.")
			playBody := []byte(fmt.Sprintf("{\"container_network\": \"%s\"}", networkName))
			if _, err = runnerClient.Play(playBody); err != nil {
				log.WithFields(log.Fields{
					"networkID":   network.ID,
					"networkName": networkName,
				}).Error("Network failed to add with runner.")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
