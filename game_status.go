package gonhl

import (
	"encoding/json"
)

func (c *Client) GetGameStatuses() ([]GameStatus, int, error) {
	var gameStatuses []GameStatus
	data, statusCode, err := c.get(STATS_BASE_URL + "/gameStatus")
	if err != nil {
		return gameStatuses, statusCode, err
	}
	json.Unmarshal(data, &gameStatuses)
	return gameStatuses, statusCode, nil
}

type GameStatus struct {
	Code              string `json:"code"`
	AbstractGameState string `json:"abstractGameState"`
	DetailedState     string `json:"detailedState"`
	BaseballCode      string `json:"baseballCode"`
	StartTimeTBD      bool   `json:"startTimeTBD"`
}
