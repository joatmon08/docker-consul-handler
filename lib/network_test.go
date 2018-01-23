package lib

import (
	"strings"
	"testing"
)

func TestShouldReturnNetworkNotFound(t *testing.T) {
	testNetwork := "notanetwork"
	expectedErrorMessage := "Error: No such network:"
	_, err := GetNetworkDetails(testNetwork)
	if err == nil {
		t.Error("Test should have thrown an error")
	}
	if !strings.Contains(err.Error(), expectedErrorMessage) {
		t.Errorf("Expected to get %s, got %s", expectedErrorMessage, err.Error())
	}
}
