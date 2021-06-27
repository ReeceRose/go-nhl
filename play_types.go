package gonhl

import (
	"encoding/json"
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
