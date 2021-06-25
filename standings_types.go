package gonhl

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetStandingsTypes() ([]StandingsType, int, error) {
	var standingsTypes []StandingsType
	data, statusCode, err := c.get(STATS_BASE_URL + "/standingsTypes")
	if err != nil {
		return standingsTypes, statusCode, err
	}
	json.Unmarshal(data, &standingsTypes)
	return standingsTypes, statusCode, nil
}

func (c *Client) GetStandingsTypeByName(name string) (StandingsType, int, error) {
	standingsTypes, statusCode, err := c.GetStandingsTypes()
	if err != nil {
		return StandingsType{}, statusCode, err
	}
	for _, standingsType := range standingsTypes {
		if standingsType.Name == name {
			return standingsType, statusCode, nil
		}
	}
	return StandingsType{}, statusCode, fmt.Errorf("cannot find standings type with name of %s", name)
}

type StandingsType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
