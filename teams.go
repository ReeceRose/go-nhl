package gonhl

import (
	"encoding/json"
	"strconv"
)

func (c *Client) GetTeams() ([]Team, int, error) {
	var response teamsResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/teams")
	if err != nil {
		return response.Teams, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Teams, statusCode, nil
}

func (c *Client) GetTeamById(id int) (Team, int, error) {
	var response teamsResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/teams/" + strconv.Itoa(id))
	if err != nil {
		return response.Teams[0], statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Teams[0], statusCode, nil
}

type teamsResponse struct {
	Copyright string `json:"copyright"`
	Teams     []Team `json:"teams"`
}

type Team struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Link  string `json:"link"`
	Venue struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Link     string `json:"link"`
		City     string `json:"city"`
		TimeZone struct {
			ID     string `json:"id"`
			Offset int    `json:"offset"`
			Tz     string `json:"tz"`
		} `json:"timeZone"`
	} `json:"venue"`
	Abbreviation    string `json:"abbreviation"`
	TeamName        string `json:"teamName"`
	LocationName    string `json:"locationName"`
	FirstYearOfPlay string `json:"firstYearOfPlay"`
	Division        struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"division"`
	Conference struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"conference"`
	Franchise struct {
		FranchiseID int    `json:"franchiseId"`
		TeamName    string `json:"teamName"`
		Link        string `json:"link"`
	} `json:"franchise"`
	ShortName       string `json:"shortName"`
	OfficialSiteURL string `json:"officialSiteUrl"`
	FranchiseID     int    `json:"franchiseId"`
	Active          bool   `json:"active"`
}
