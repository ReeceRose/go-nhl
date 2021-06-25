package gonhl

import "encoding/json"

func (c *Client) GetStandingsTypes() ([]StandingsTypes, int, error) {
	var standingsTypes []StandingsTypes
	data, statusCode, err := c.get(STATS_BASE_URL + "/standingsTypes")
	if err != nil {
		return standingsTypes, statusCode, err
	}
	json.Unmarshal(data, &standingsTypes)
	return standingsTypes, statusCode, nil
}

type StandingsTypes struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
