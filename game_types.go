package gonhl

import (
	"encoding/json"
)

func (c *Client) GetGameTypes() ([]GameType, int, error) {
	var gametypes []GameType
	data, statusCode, err := c.get(STATS_BASE_URL + "/gameTypes")
	if err != nil {
		return gametypes, statusCode, err
	}
	json.Unmarshal(data, &gametypes)
	return gametypes, statusCode, nil
}

type GameType struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Postseason  bool   `json:"postseason"`
}
