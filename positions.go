package gonhl

import "encoding/json"

func (c *Client) GetPositions() ([]Position, int, error) {
	var positions []Position
	data, statusCode, err := c.get(STATS_BASE_URL + "/positions/")
	if err != nil {
		return positions, statusCode, err
	}
	json.Unmarshal(data, &positions)
	return positions, statusCode, nil
}

type Position struct {
	Abbrev   string `json:"abbrev"`
	Code     string `json:"code"`
	FullName string `json:"fullName"`
	Type     string `json:"type"`
}
