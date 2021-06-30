package gonhl

import (
	"encoding/json"
)

func (c *Client) getMilestone(appendedURL string) (Milestone, int, error) {
	var response Milestone
	data, statusCode, err := c.get(RECORDS_BASE_URL + "/milestone-" + appendedURL)
	if err != nil {
		return response, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response, statusCode, nil
}

func (c *Client) Get1000PointCareerMilestone() (Milestone, int, error) {
	return c.getMilestone("1000-point-career")
}

func (c *Client) Get100PointSeasonMilestone() (Milestone, int, error) {
	return c.getMilestone("100-point-season")
}

func (c *Client) Get500GoalCareerMilestone() (Milestone, int, error) {
	return c.getMilestone("500-goal-career")
}

func (c *Client) Get50GoalSeasonMilestone() (Milestone, int, error) {
	return c.getMilestone("50-goal-season")
}

func (c *Client) Get5GoalGameMilestone() (Milestone, int, error) {
	return c.getMilestone("1000-point-career")
}

// TODO: refactor this
type Milestone struct {
	Data []struct {
		ID struct {
			GameID   int `json:"gameId"`
			SeasonID int `json:"seasonId"`
			SkaterID int `json:"skaterId"`
		} `json:"id"`
		ActivePlayer     int    `json:"activePlayer"`
		AgeInDaysForYear int    `json:"ageInDaysForYear"`
		AgeInYears       int    `json:"ageInYears"`
		EmptyNetGoals    int    `json:"emptyNetGoals"`
		GameDate         string `json:"gameDate"`
		GameID           int    `json:"gameId"`
		GameTypeID       int    `json:"gameTypeId"`
		Goalie1FirstName string `json:"goalie1FirstName"`
		Goalie1ID        int    `json:"goalie1Id"`
		Goalie1LastName  string `json:"goalie1LastName"`
		Goalie2FirstName string `json:"goalie2FirstName"`
		Goalie2ID        int    `json:"goalie2Id"`
		Goalie2LastName  string `json:"goalie2LastName"`
		Goals            int    `json:"goals"`
		HomeScore        int    `json:"homeScore"`
		HomeTeamID       int    `json:"homeTeamId"`
		HomeTriCode      string `json:"homeTriCode"`
		PenaltyShotGoals int    `json:"penaltyShotGoals"`
		SeasonID         int    `json:"seasonId"`
		SkaterFirstName  string `json:"skaterFirstName"`
		SkaterID         int    `json:"skaterId"`
		SkaterLastName   string `json:"skaterLastName"`
		SkaterTeamID     int    `json:"skaterTeamId"`
		SkaterTeamName   string `json:"skaterTeamName"`
		VisitorScore     int    `json:"visitorScore"`
		VisitorTeamID    int    `json:"visitorTeamId"`
		VisitorTriCode   string `json:"visitorTriCode"`
	} `json:"data"`
	Total int `json:"total"`
}
