package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"testing"
	"strings"
	"os/exec"
	"path/filepath"
	"os"
)

func readFixture(filename string) (string, error) {
	file, e := ioutil.ReadFile("fixtures/" + filename)
	if e != nil {
		return "", e
	}
	return string(file), nil
}

func setup() *exec.Cmd {
	response := string("{\"playbook_return_code\": 0}")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	absPath, _ := filepath.Abs("./handler")
	subproc := exec.Command(absPath)
	env := os.Environ()
	env = append(env, fmt.Sprintf("CONSUL_HANDLER_RUNNER=%s", ts.URL))
	subproc.Env = env
	return subproc
}

func TestShouldOutputNewestNetwork(t *testing.T) {
	expectedNewestNetwork := "ee15ed2cc513bc95923b94794c14bb2bf4928a7a09bc051c15c9f785c4556068"
	subproc := setup()
    input, _ := readFixture("networks.json")
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
