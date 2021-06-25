package gonhl

import (
	"encoding/json"
	"fmt"
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

func (c *Client) GetEventTypeById(id string) (EventType, int, error) {
	eventTypes, statusCode, err := c.GetEventTypes()
	if err != nil {
		return EventType{}, statusCode, err
	}
	for _, eventType := range eventTypes {
		if eventType.ID == id {
			return eventType, statusCode, nil
		}
	}
	return EventType{}, statusCode, fmt.Errorf("cannot find event type by id of %s", id)
}

type EventType struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}
