package gonhl

import (
	"encoding/json"
	"fmt"
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

func (c *Client) GetGameTypeById(id string) (GameType, int, error) {
	gameTypes, statusCode, err := c.GetGameTypes()
	if err != nil {
		return GameType{}, statusCode, err
	}
	for _, gameType := range gameTypes {
		if gameType.ID == id {
			return gameType, statusCode, nil
		}
	}
	return GameType{}, statusCode, fmt.Errorf("cannot find game type by id of %s", id)
}

type GameType struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Postseason  bool   `json:"postseason"`
}
