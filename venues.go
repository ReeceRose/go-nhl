package gonhl

import (
	"encoding/json"
	"strconv"
)

func (c *Client) GetVenues() ([]Venue, int, error) {
	var response Response
	data, statusCode, err := c.get(STATS_BASE_URL + "/venues")
	if err != nil {
		return response.Venues, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Venues, statusCode, nil
}

func (c *Client) GetVenueById(id int) (Venue, int, error) {
	var response Response
	data, statusCode, err := c.get(STATS_BASE_URL + "/venues/" + strconv.Itoa(id))
	if err != nil {
		return response.Venues[0], statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Venues[0], statusCode, nil
}

type Response struct {
	Copyright string  `json:"copyright"`
	Venues    []Venue `json:"venues"`
}

type Venue struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Link       string `json:"link"`
	AppEnabled string `json:"appEnabled"`
}
