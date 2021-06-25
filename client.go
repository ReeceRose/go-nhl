package gonhl

import (
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *Client) get(endpoint string) ([]byte, int, error) {
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, -100, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, -101, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}
	return body, response.StatusCode, nil
}
