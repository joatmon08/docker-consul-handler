package lib

import (
	"io/ioutil"
	"testing"
)

func readFixture(filename string) ([]byte, error) {
	file, e := ioutil.ReadFile("fixtures/" + filename)
	if e != nil {
		return nil, e
	}
	return file, nil
}

func TestShouldGetNewestNetwork(t *testing.T) {
	expectedNetwork := "ee15ed2cc513bc95923b94794c14bb2bf4928a7a09bc051c15c9f785c4556068"
	data, err := readFixture("consul_output.json")
	if err != nil {
		t.Errorf("Error reading test fixture : %s", err.Error())
	}
	network := &Network{CreateIndex: 11}
	isNew, err := network.HasNewNetwork(data)
	if err != nil {
		t.Errorf("Error parsing networks : %s", err.Error())
	}
	if !isNew {
		t.Errorf("Should return true")
	}
	t.Logf("Newest network is %s", network.ID)
	if network.ID != expectedNetwork {
		t.Errorf("Expected %s, actual %s", expectedNetwork, network.ID)
	}
}

func TestShouldReturnBlankNetwork(t *testing.T) {
	expectedNetwork := ""
	data, err := readFixture("consul_output_no_networks.json")
	if err != nil {
		t.Errorf("Error reading test fixture : %s", err.Error())
	}
	network := &Network{}
	isNew, err := network.HasNewNetwork(data)
	if err != nil {
		t.Errorf("Error parsing networks : %s", err.Error())
	}
	if isNew {
		t.Errorf("Should return false, since blank network is not a new network")
	}
	t.Logf("Newest network is %s", network.ID)
	if network.ID != expectedNetwork {
		t.Errorf("Expected %s, actual %s", expectedNetwork, network.ID)
	}
}

func TestShouldIgnoreModificationAndReturnBlank(t *testing.T) {
	expectedNetwork := ""
	data, err := readFixture("consul_output_modification.json")
	if err != nil {
		t.Errorf("Error reading test fixture : %s", err.Error())
	}
	network := &Network{CreateIndex: 17}
	isNew, err := network.HasNewNetwork(data)
	if err != nil {
		t.Errorf("Error parsing networks : %s", err.Error())
	}
	if isNew {
		t.Errorf("Should return false, since this is a modification")
	}
	t.Logf("Newest network is %s", network.ID)
	if network.ID != expectedNetwork {
		t.Errorf("Expected %s, actual %s", expectedNetwork, network.ID)
	}
}

func TestShouldFailToConvertData(t *testing.T) {
	var data []byte
	network := &Network{}
	_, err := network.HasNewNetwork(data)
	if err == nil {
		t.Errorf("Should fail due to lack of data to convert : %s", err.Error())
	}
}
