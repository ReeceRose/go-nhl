package gonhl

import (
	"encoding/json"
)

func (c *Client) GetAwards() ([]Award, int, error) {
	var response awardsResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/awards")
	if err != nil {
		return response.Awards, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Awards, statusCode, nil
}

type awardsResponse struct {
	Copyright string  `json:"copyright"`
	Awards    []Award `json:"awards"`
}

type Award struct {
	Name          string `json:"name"`
	ShortName     string `json:"shortName"`
	Description   string `json:"description"`
	RecipientType string `json:"recipientType"`
	History       string `json:"history,omitempty"`
	ImageURL      string `json:"imageUrl"`
	HomePageURL   string `json:"homePageUrl"`
	Link          string `json:"link"`
}
