package gonhl

import "encoding/json"

func (c *Client) getSeasons(appendedURL string) ([]Season, int, error) {
	var response seasonsResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/seasons/" + appendedURL)
	if err != nil {
		return response.Seasons, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Seasons, statusCode, nil
}

func (c *Client) GetSeasons() ([]Season, int, error) {
	return c.getSeasons("")
}

func (c *Client) GetCurrentSeason() (Season, int, error) {
	seasons, responseCode, error := c.getSeasons("current")
	return seasons[0], responseCode, error
}

func (c *Client) GetSeasonByYear(year string) (Season, int, error) {
	seasons, responseCode, error := c.getSeasons(year)
	return seasons[0], responseCode, error
}

type seasonsResponse struct {
	Copyright string   `json:"copyright"`
	Seasons   []Season `json:"seasons"`
}

type Season struct {
	SeasonID               string `json:"seasonId"`
	RegularSeasonStartDate string `json:"regularSeasonStartDate"`
	RegularSeasonEndDate   string `json:"regularSeasonEndDate"`
	SeasonEndDate          string `json:"seasonEndDate"`
	NumberOfGames          int    `json:"numberOfGames"`
	TiesInUse              bool   `json:"tiesInUse"`
	OlympicsParticipation  bool   `json:"olympicsParticipation"`
	ConferencesInUse       bool   `json:"conferencesInUse"`
	DivisionsInUse         bool   `json:"divisionsInUse"`
	WildCardInUse          bool   `json:"wildCardInUse"`
}
