package gonhl

import (
	"encoding/json"
	"strconv"
)

func (c *Client) GetDivisions() ([]Division, int, error) {
	var response divisionsResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/divisions")
	if err != nil {
		return response.Divisions, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Divisions, statusCode, nil
}

func (c *Client) GetDivisionById(id int) (Division, int, error) {
	var response divisionsResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/divisions/" + strconv.Itoa(id))
	if err != nil {
		return response.Divisions[0], statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Divisions[0], statusCode, nil
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
