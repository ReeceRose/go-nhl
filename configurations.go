package gonhl

import "encoding/json"

func (c *Client) GetConfigurations() ([]Configuration, int, error) {
	var configurations []Configuration
	data, statusCode, err := c.get(STATS_BASE_URL + "/configurations")
	if err != nil {
		return configurations, statusCode, err
	}
	json.Unmarshal(data, &configurations)
	return configurations, statusCode, nil
}

type Configuration struct {
	Description string `json:"description"`
	Endpoint    string `json:"endpoint"`
}
