package gonhl

import (
	"encoding/json"
)

func (c *Client) GetStandingsTypes() ([]StandingsType, int, error) {
	var standingsTypes []StandingsType
	data, statusCode, err := c.get(STATS_BASE_URL + "/standingsTypes")
	if err != nil {
		return standingsTypes, statusCode, err
	}
	json.Unmarshal(data, &standingsTypes)
	return standingsTypes, statusCode, nil
}

type StandingsType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
