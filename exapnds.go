package gonhl

import (
	"encoding/json"
)

func (c *Client) GetExpands() ([]Expand, int, error) {
	var expands []Expand
	data, statusCode, err := c.get(STATS_BASE_URL + "/expands")
	if err != nil {
		return expands, statusCode, err
	}
	json.Unmarshal(data, &expands)
	return expands, statusCode, nil
}

type Expand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
