package gonhl

import "encoding/json"

func (c *Client) GetPerformerTypes() ([]PerformerType, int, error) {
	var performerType []PerformerType
	data, statusCode, err := c.get(STATS_BASE_URL + "/performerTypes")
	if err != nil {
		return performerType, statusCode, err
	}
	json.Unmarshal(data, &performerType)
	return performerType, statusCode, nil
}

type PerformerType struct {
	Name string `json:"name"`
}
