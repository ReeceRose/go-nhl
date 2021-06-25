package gonhl

import (
	"encoding/json"
)

func (c *Client) GetEventTypes() ([]EventType, int, error) {
	var eventTypes []EventType
	data, statusCode, err := c.get(STATS_BASE_URL + "/eventTypes")
	if err != nil {
		return eventTypes, statusCode, err
	}
	json.Unmarshal(data, &eventTypes)
	return eventTypes, statusCode, nil
}

type EventType struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}
