package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/joatmon08/consul-handler/lib"
	"os"
	"strings"
)

func main() {
	filename := flag.String("filename", "/scripts/consul_watch.log", "file path")
	flag.Parse()

	f, err := os.OpenFile(*filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	existingNetworkIDs := []string{}

	for scanner.Scan() {
		input := scanner.Bytes()
		networkIDs, _ := lib.ParseNetworkIDs(input)
		newNetworks := lib.CompareNetworks(existingNetworkIDs, networkIDs)
		stringNetworkIDs := strings.Join(newNetworks, " ")
		if _, err = f.WriteString(stringNetworkIDs + "\n"); err != nil {
			panic(err)
		}
		existingNetworkIDs = networkIDs
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
