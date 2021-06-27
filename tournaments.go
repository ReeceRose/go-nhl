package gonhl

import "encoding/json"

func (c *Client) GetTournamentByName(name string) (Tournament, int, error) {
	var tournament Tournament
	data, statusCode, err := c.get(STATS_BASE_URL + "/tournaments/" + name)
	if err != nil {
		return tournament, statusCode, err
	}
	json.Unmarshal(data, &tournament)
	return tournament, statusCode, nil
}

type Tournament struct {
	Copyright    string `json:"copyright"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Season       string `json:"season"`
	DefaultRound int    `json:"defaultRound"`
	Rounds       []struct {
		Number int `json:"number"`
		Code   int `json:"code"`
		Names  struct {
			Name      string `json:"name"`
			ShortName string `json:"shortName"`
		} `json:"names"`
		Format struct {
			Name          string `json:"name"`
			Description   string `json:"description"`
			NumberOfGames int    `json:"numberOfGames"`
			NumberOfWins  int    `json:"numberOfWins"`
		} `json:"format"`
	} `json:"rounds"`
}
