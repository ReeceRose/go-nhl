package gonhl

import (
	"encoding/json"
)

func (c *Client) GetVenues() ([]Venue, int, error) {
	var response venueResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/venues")
	if err != nil {
		return response.Venues, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Venues, statusCode, nil
}

type venueResponse struct {
	Copyright string  `json:"copyright"`
	Venues    []Venue `json:"venues"`
}

type Venue struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Link       string `json:"link"`
	AppEnabled string `json:"appEnabled"`
}
