package gonhl

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPerformerTypes() ([]PerformerType, int, error) {
	var performerType []PerformerType
	data, statusCode, err := c.get(STATS_BASE_URL + "/performerTypes")
	if err != nil {
		return performerType, statusCode, err
	}
	json.Unmarshal(data, &performerType)
	return performerType, statusCode, nil
}

func (c *Client) GetPerformerTypeByName(name string) (PerformerType, int, error) {
	performerTypes, statusCode, err := c.GetPerformerTypes()
	if err != nil {
		return PerformerType{}, statusCode, err
	}
	for _, performerType := range performerTypes {
		if performerType.Name == name {
			return performerType, statusCode, nil
		}
	}
	return PerformerType{}, statusCode, fmt.Errorf("cannot find performer type with name of %s", name)
}

type PerformerType struct {
	Name string `json:"name"`
}
