package gonhl

import (
	"encoding/json"
)

func (c *Client) GetStatTypes() ([]StatType, int, error) {
	var statTypes []StatType
	data, statusCode, err := c.get(STATS_BASE_URL + "/statTypes")
	if err != nil {
		return statTypes, statusCode, err
	}
	json.Unmarshal(data, &statTypes)
	return statTypes, statusCode, nil
}

type StatType struct {
	DisplayName string `json:"displayName"`
	GameType    struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		Postseason  bool   `json:"postseason"`
	} `json:"gameType"`
}
