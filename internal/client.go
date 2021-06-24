package internal

import (
	"io/ioutil"
	"net/http"
	"time"
)

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *Client) get(endpoint string) ([]byte, int) {
	request, err := http.NewRequest("GET", BASE_URL+endpoint, nil)
	if err != nil {
		return nil, FAILED_TO_CREATE_REQUEST
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, FAILED_TO_EXECUTE_REQUEST
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body, response.StatusCode
}
