package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func readFixture(filename string) (string, error) {
	file, e := ioutil.ReadFile("fixtures/" + filename)
	if e != nil {
		return "", e
	}
	return string(file), nil
}

func setup() *exec.Cmd {
	responseRunner := string("{\"playbook_return_code\": 0}")
	tsRunner := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, responseRunner)
	}))
	defer tsRunner.Close()

	absPath, _ := filepath.Abs("./handler")
	subproc := exec.Command(absPath)
	env := os.Environ()
	env = append(env, fmt.Sprintf("CONSUL_HANDLER_RUNNER=%s", tsRunner.URL))
	subproc.Env = env
	return subproc
}

func TestShouldOutputNewestNetwork(t *testing.T) {
	expectedNewestNetwork := "ee15ed2cc513bc95923b94794c14bb2bf4928a7a09bc051c15c9f785c4556068"
	subproc := setup()
	input, _ := readFixture("networks.txt")
	subproc.Stdin = strings.NewReader(input)
	output, err := subproc.Output()
	handlerOut := string(output)
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	}
	if !strings.Contains(handlerOut, expectedNewestNetwork) {
		t.Errorf("Expected %s, got %s", expectedNewestNetwork, handlerOut)
	}
	subproc.Wait()
}

func TestShouldIgnoreModifiedNetwork(t *testing.T) {
	expectedOutput := "No new networks added."
	subproc := setup()
	input, _ := readFixture("networks_modification.txt")
	subproc.Stdin = strings.NewReader(input)
	output, err := subproc.Output()
	handlerOut := string(output)
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	}
	if !strings.Contains(handlerOut, expectedOutput) {
		t.Errorf("Expected %s, got %s", expectedOutput, handlerOut)
	}
	subproc.Wait()
}

func TestShouldIgnoreBlankNetwork(t *testing.T) {
	expectedOutput := ""
	subproc := setup()
	subproc.Stdin = strings.NewReader("\n")
	output, err := subproc.Output()
	handlerOut := string(output)
	if err != nil {
		t.Errorf("Got error %s", err.Error())
	}
	if !strings.Contains(handlerOut, expectedOutput) {
		t.Errorf("Expected %s, got %s", expectedOutput, handlerOut)
	}
	subproc.Wait()
}
