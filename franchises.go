package gonhl

import (
	"encoding/json"
	"strconv"
)

func (c *Client) getFranchises(appendedURL string) ([]Franchise, int, error) {
	var response franchisesResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/franchises/" + appendedURL)
	if err != nil {
		return response.Franchises, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Franchises, statusCode, nil
}

func (c *Client) GetFranchises() ([]Franchise, int, error) {
	return c.getFranchises("")
}

func (c *Client) GetFranchiseById(id int) (Franchise, int, error) {
	franchises, statusCode, err := c.getFranchises(strconv.Itoa(id))
	return franchises[0], statusCode, err
}

type franchisesResponse struct {
	Copyright  string      `json:"copyright"`
	Franchises []Franchise `json:"franchises"`
}

type Franchise struct {
	FranchiseID      int    `json:"franchiseId"`
	FirstSeasonID    int    `json:"firstSeasonId"`
	MostRecentTeamID int    `json:"mostRecentTeamId"`
	TeamName         string `json:"teamName"`
	LocationName     string `json:"locationName"`
	Link             string `json:"link"`
	LastSeasonID     int    `json:"lastSeasonId,omitempty"`
}
