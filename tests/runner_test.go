package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joatmon08/docker-consul-handler/runner"
)

func TestShouldReturnSuccess(t *testing.T) {
	response := string("{\"playbook_return_code\": 0}")
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, response)
	}))
	defer ts.Close()

	client := runner.Client{
		URL: ts.URL,
	}

	resp, err := client.Play([]byte("{\"container_network\": \"test\"}"))

	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if strings.TrimRight(string(resp), "\n") != response {
		t.Errorf("expected %s, got %s", response, resp)
	}

}

func TestShouldFail(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"message\": \"ansible playbook returned error code 2\"}"))
	}))
	defer ts.Close()

	client := runner.Client{
		URL: ts.URL,
	}

	_, err := client.Play([]byte("{\"container_network\": \"test\"}"))

	if err == nil {
		t.Error("Should have thrown error")
	}
}
