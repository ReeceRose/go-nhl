package gonhl

import (
	"encoding/json"
	"strconv"
)

func (c *Client) getConferences(appendedURL string) ([]Conference, int, error) {
	var response conferencesResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/conferences/" + appendedURL)
	if err != nil {
		return response.Conferences, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Conferences, statusCode, nil
}

func (c *Client) GetConferences() ([]Conference, int, error) {
	return c.getConferences("")
}

func (c *Client) GetConferenceById(id int) (Conference, int, error) {
	conferences, statusCode, err := c.getConferences(strconv.Itoa(id))
	return conferences[0], statusCode, err
}

type conferencesResponse struct {
	Copyright   string       `json:"copyright"`
	Conferences []Conference `json:"conferences"`
}
type Conference struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Link         string `json:"link"`
	Abbreviation string `json:"abbreviation"`
	ShortName    string `json:"shortName"`
	Active       bool   `json:"active"`
}
