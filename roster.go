package gonhl

import (
	"encoding/json"
	"strconv"
)

func (c *Client) GetRoster(id int) ([]RosterInformation, int, error) {
	var response rosterResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/teams/" + strconv.Itoa(id) + "/roster")
	if err != nil {
		return response.Roster, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Roster, statusCode, nil
}

type rosterResponse struct {
	Copyright string              `json:"copyright"`
	Roster    []RosterInformation `json:"roster"`
	Link      string              `json:"link"`
}

type RosterInformation struct {
	Person struct {
		ID       int    `json:"id"`
		FullName string `json:"fullName"`
		Link     string `json:"link"`
	} `json:"person"`
	JerseyNumber string `json:"jerseyNumber"`
	Position     struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Abbreviation string `json:"abbreviation"`
	} `json:"position"`
}
