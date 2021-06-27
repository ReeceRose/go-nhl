package gonhl

import (
	"encoding/json"
)

func (c *Client) GetPlatforms() ([]Platform, int, error) {
	var platforms []Platform
	data, statusCode, err := c.get(STATS_BASE_URL + "/platforms")
	if err != nil {
		return platforms, statusCode, err
	}
	json.Unmarshal(data, &platforms)
	return platforms, statusCode, nil
}

type Platform struct {
	PlatformCode        string `json:"platformCode"`
	PlatformDescription string `json:"platformDescription"`
}
