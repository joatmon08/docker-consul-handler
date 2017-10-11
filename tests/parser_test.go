package tests

import (
	"github.com/joatmon08/consul-handler/lib"
	"io/ioutil"
	"testing"
)

func readFixture(filename string) ([]byte, error) {
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		return nil, e
	}
	return file, nil
}

func TestShouldGetNewestModifyIndex(t *testing.T) {
	expectedNetwork := "ee15ed2cc513bc95923b94794c14bb2bf4928a7a09bc051c15c9f785c4556068"
	data, err := readFixture("fixtures/consul_output.json")
	if err != nil {
		t.Errorf("Error reading test fixture : %s", err.Error())
	}
	network, err := lib.GetNewestNetwork(data)
	if err != nil {
		t.Errorf("Error parsing networks : %s", err.Error())
	}
	t.Logf("Newest network is %s", network)
	if network != expectedNetwork {
		t.Errorf("Expected %s, actual %s", expectedNetwork, network)
	}
}

func TestShouldReturnBlankNetwork(t *testing.T) {
	expectedNetwork := ""
	data, err := readFixture("fixtures/consul_output_no_networks.json")
	if err != nil {
		t.Errorf("Error reading test fixture : %s", err.Error())
	}
	network, err := lib.GetNewestNetwork(data)
	if err != nil {
		t.Errorf("Error parsing networks : %s", err.Error())
	}
	t.Logf("Newest network is %s", network)
	if network != expectedNetwork {
		t.Errorf("Expected %s, actual %s", expectedNetwork, network)
	}
}

func TestShouldFailToConvertData(t *testing.T) {
	var data []byte
	_, err := lib.GetNewestNetwork(data)
	if err == nil {
		t.Errorf("Should fail due to lack of data to convert : %s", err.Error())
	}
}
