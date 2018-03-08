package runner

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PathPlay denotes the path to run the playbook on the runner.
const PathPlay = "/runner/play"

// Client is the runner's client interface.
type Client struct {
	URL string
}

// Play runs the playbook.
func (r Client) Play(extraVars []byte) ([]byte, error) {
	url := r.URL + PathPlay
	request := bytes.NewReader(extraVars)
	req, err := http.NewRequest("POST", url, request)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 399 {
		errMessage := fmt.Sprintf("Request failed with status code %d and body %s", resp.StatusCode, body)
		return nil, errors.New(errMessage)
	}
	return body, nil
}
