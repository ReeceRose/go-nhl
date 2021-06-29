package gonhl

import (
	"encoding/json"
	"strconv"
)

func (c *Client) getVenues(appendedURL string) ([]Venue, int, error) {
	var response venueResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/venues/" + appendedURL)
	if err != nil {
		return response.Venues, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Venues, statusCode, nil
}

func (c *Client) GetVenues() ([]Venue, int, error) {
	return c.getVenues("")
}

func (c *Client) GetVenueById(id int) (Venue, int, error) {
	venues, statusCode, err := c.getVenues(strconv.Itoa(id))
	return venues[0], statusCode, err
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
