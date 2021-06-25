package gonhl

import (
	"encoding/json"
	"fmt"
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

func (c *Client) GetPlatformByCode(code string) (Platform, int, error) {
	platforms, statusCode, err := c.GetPlatforms()
	if err != nil {
		return Platform{}, statusCode, err
	}
	for _, platform := range platforms {
		if platform.PlatformCode == code {
			return platform, statusCode, nil
		}
	}
	return Platform{}, statusCode, fmt.Errorf("cannot find platform with code of %s", code)
}

type Platform struct {
	PlatformCode        string `json:"platformCode"`
	PlatformDescription string `json:"platformDescription"`
}
