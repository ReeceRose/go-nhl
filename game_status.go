package gonhl

import (
	"encoding/json"
	"fmt"
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

func (c *Client) GetGameStatusBycode(code string) (GameStatus, int, error) {
	gameStatuses, statusCode, err := c.GetGameStatuses()
	if err != nil {
		return GameStatus{}, statusCode, err
	}
	for _, gameStatus := range gameStatuses {
		if gameStatus.Code == code {
			return gameStatus, statusCode, nil
		}
	}
	return GameStatus{}, statusCode, fmt.Errorf("cannot find game status by code of %s", code)
}

type GameStatus struct {
	Code              string `json:"code"`
	AbstractGameState string `json:"abstractGameState"`
	DetailedState     string `json:"detailedState"`
	BaseballCode      string `json:"baseballCode"`
	StartTimeTBD      bool   `json:"startTimeTBD"`
}
