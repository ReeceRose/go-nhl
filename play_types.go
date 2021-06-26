package gonhl

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPlayTypes() ([]PlayType, int, error) {
	var playTypes []PlayType
	data, statusCode, err := c.get(STATS_BASE_URL + "/playTypes")
	if err != nil {
		return playTypes, statusCode, err
	}
	json.Unmarshal(data, &playTypes)
	return playTypes, statusCode, nil
}

func (c *Client) GetPlayTypeById(id string) (PlayType, int, error) {
	playTypes, statusCode, err := c.GetPlayTypes()
	if err != nil {
		return PlayType{}, statusCode, err
	}
	for _, playType := range playTypes {
		if playType.ID == id {
			return playType, statusCode, nil
		}
	}
	return PlayType{}, statusCode, fmt.Errorf("cannot find play type by id of %s", id)
}

type PlayType struct {
	Name                string        `json:"name"`
	ID                  string        `json:"id"`
	CmsKey              string        `json:"cmsKey"`
	PlayerTypes         []PlayerType  `json:"playerTypes"`
	Code                string        `json:"code"`
	SecondaryEventCodes []interface{} `json:"secondaryEventCodes"`
}

type PlayerType struct {
	Type string `json:"playerType"`
}
