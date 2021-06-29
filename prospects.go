package gonhl

import (
	"encoding/json"
	"strconv"
)

func (c *Client) getProspects(appendURL string) ([]Prospect, int, error) {
	var response prospectResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/draft/prospects/" + appendURL)
	if err != nil {
		return response.Prospects, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Prospects, statusCode, nil
}

func (c *Client) GetProspects() ([]Prospect, int, error) {
	return c.getProspects("")
}

func (c *Client) GetProspectById(id int) (Prospect, int, error) {
	prospects, statusCode, err := c.getProspects(strconv.Itoa(id))
	return prospects[0], statusCode, err
}

type prospectResponse struct {
	Copyright string     `json:"copyright"`
	Prospects []Prospect `json:"prospects"`
}

type Prospect struct {
	ID                 int    `json:"id"`
	FullName           string `json:"fullName"`
	Link               string `json:"link"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	BirthDate          string `json:"birthDate"`
	BirthCity          string `json:"birthCity"`
	BirthStateProvince string `json:"birthStateProvince"`
	BirthCountry       string `json:"birthCountry"`
	Height             string `json:"height"`
	Weight             int    `json:"weight"`
	ShootsCatches      string `json:"shootsCatches"`
	PrimaryPosition    struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Abbreviation string `json:"abbreviation"`
	} `json:"primaryPosition"`
	DraftStatus      string `json:"draftStatus"`
	ProspectCategory struct {
		ID        int    `json:"id"`
		ShortName string `json:"shortName"`
		Name      string `json:"name"`
	} `json:"prospectCategory"`
	AmateurTeam struct {
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"amateurTeam"`
	AmateurLeague struct {
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"amateurLeague"`
	Ranks struct {
		FinalRank int `json:"finalRank"`
		DraftYear int `json:"draftYear"`
	} `json:"ranks"`
}
