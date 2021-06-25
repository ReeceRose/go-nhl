package gonhl

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (c *Client) GetConferences() ([]Conference, int, error) {
	var response conferencesResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/conferences")
	if err != nil {
		return response.Conferences, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Conferences, statusCode, nil
}

func (c *Client) GetConferenceById(id int) (Conference, int, error) {
	var response conferencesResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/conferences/" + strconv.Itoa(id))
	if err != nil {
		return response.Conferences[0], statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Conferences[0], statusCode, nil
}

func (c *Client) GetConferenceByName(name string) (Conference, int, error) {
	conferences, statusCode, err := c.GetConferences()
	if err != nil {
		return Conference{}, statusCode, err
	}
	for _, conference := range conferences {
		if conference.Name == name {
			return conference, statusCode, nil
		}
	}
	return Conference{}, statusCode, fmt.Errorf("cannot find conference by name of %s", name)
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
