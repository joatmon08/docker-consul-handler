package runner

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const PATH_PLAY = "/runner/play"

type Client struct {
	URL string
}

func (r Client) Play(extraVars []byte) ([]byte, error) {
	url := r.URL + PATH_PLAY
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
		err_message := fmt.Sprintf("Request failed with status code %s and body %s", resp.StatusCode, body)
		return nil, errors.New(err_message)
	}
	return body, nil
}
