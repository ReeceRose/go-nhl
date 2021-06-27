package gonhl

import "encoding/json"

func (c *Client) GetTournamentTypes() ([]TournamentType, int, error) {
	var tournamentTypes []TournamentType
	data, statusCode, err := c.get(STATS_BASE_URL + "/tournamentTypes")
	if err != nil {
		return tournamentTypes, statusCode, err
	}
	json.Unmarshal(data, &tournamentTypes)
	return tournamentTypes, statusCode, nil
}

type TournamentType struct {
	Description  string `json:"description"`
	GameTypeEnum struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		Postseason  bool   `json:"postseason"`
	} `json:"gameTypeEnum"`
	Parameter string `json:"parameter"`
}
