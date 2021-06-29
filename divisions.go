package gonhl

import (
	"encoding/json"
	"strconv"
)

func (c *Client) getDivisions(appendedURL string) ([]Division, int, error) {
	var response divisionsResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/divisions/" + appendedURL)
	if err != nil {
		return response.Divisions, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Divisions, statusCode, nil
}

func (c *Client) GetDivisions() ([]Division, int, error) {
	return c.getDivisions("")
}

func (c *Client) GetDivisionById(id int) (Division, int, error) {
	divisions, statusCode, err := c.getDivisions(strconv.Itoa(id))
	return divisions[0], statusCode, err
}

type divisionsResponse struct {
	Copyright string     `json:"copyright"`
	Divisions []Division `json:"divisions"`
}
type Division struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Link         string `json:"link"`
	Abbreviation string `json:"abbreviation"`
	Conference   struct {
		Link string `json:"link"`
	} `json:"conference"`
	Active bool `json:"active"`
}
