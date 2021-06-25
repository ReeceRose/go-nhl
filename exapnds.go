package gonhl

import (
	"encoding/json"
	"fmt"
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

func (c *Client) GetExpandByName(name string) (Expand, int, error) {
	expands, statusCode, err := c.GetExpands()
	if err != nil {
		return Expand{}, statusCode, err
	}
	for _, expand := range expands {
		if expand.Name == name {
			return expand, statusCode, nil
		}
	}
	return Expand{}, statusCode, fmt.Errorf("cannot find expand with name of %s", name)
}

type Expand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
