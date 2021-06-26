package gonhl

import (
	"encoding/json"
	"strconv"
)

func (c *Client) GetPeopleById(id int) (People, int, error) {
	var response peopleResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/people/" + strconv.Itoa(id))
	if err != nil {
		return response.People, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.People, statusCode, nil
}

func (c *Client) GetPlayerById(id int) (People, int, error) {
	return c.GetPeopleById(id)
}

func (c *Client) getStats(id int, statType string, season string) ([]PlayerStats, int, error) {
	var response statsResponse
	appendedURL := ""
	if season != "" {
		appendedURL = "&season=" + season
	}
	data, statusCode, err := c.get(STATS_BASE_URL + "/people/" + strconv.Itoa(id) + "/stats?stats=" + statType + appendedURL)
	if err != nil {
		return response.Stats, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Stats, statusCode, nil
}

func (c *Client) GetPeopleStatsById(id int) (PlayerStats, int, error) {
	playerStats, responseCode, error := c.getStats(id, "seasonBySeason", "")
	return playerStats[0], responseCode, error
}

func (c *Client) GetPlayerStatsById(id int) (PlayerStats, int, error) {
	return c.GetPeopleStatsById(id)
}

func (c *Client) GetPeopleStatsBySeason(id int, season string) (PlayerStats, int, error) {
	playerStats, responseCode, error := c.getStats(id, "singleSeasonStats", season)
	return playerStats[0], responseCode, error
}

func (c *Client) GetPlayerStatsBySeason(id int, season string) (PlayerStats, int, error) {
	return c.GetPeopleStatsBySeason(id, season)
}

func (c *Client) GetPeopleStatsBySeasonSplitHomeAway(id int, season string) ([]PlayerStats, int, error) {
	return c.getStats(id, "homeAndAway", season)
}

func (c *Client) GetPlayerStatsBySeasonSplitHomeAway(id int, season string) ([]PlayerStats, int, error) {
	return c.GetPeopleStatsBySeasonSplitHomeAway(id, season)
}

func (c *Client) GetPeopleStatsBySeasonSplitWinLose(id int, season string) ([]PlayerStats, int, error) {
	return c.getStats(id, "winLoss", season)
}

func (c *Client) GetPlayerStatsBySeasonSplitWinLoss(id int, season string) ([]PlayerStats, int, error) {
	return c.GetPeopleStatsBySeasonSplitWinLose(id, season)
}

func (c *Client) GetPeopleStatsBySeasonSplitByMonth(id int, season string) ([]PlayerStats, int, error) {
	return c.getStats(id, "byMonth", season)
}

func (c *Client) GetPlayerStatsBySeasonSplotByMonth(id int, season string) ([]PlayerStats, int, error) {
	return c.GetPeopleStatsBySeasonSplitByMonth(id, season)
}

func (c *Client) GetPeopleStatsBySeasonSplitByDayOfWeek(id int, season string) ([]PlayerStats, int, error) {
	return c.getStats(id, "byDayOfWeek", season)
}

func (c *Client) GetPlayerStatsBySeasonSplitByDayOfWeek(id int, season string) ([]PlayerStats, int, error) {
	return c.GetPeopleStatsBySeasonSplitByDayOfWeek(id, season)
}

func (c *Client) GetPeopleStatsBySeasonVsDivision(id int, season string) ([]PlayerStats, int, error) {
	return c.getStats(id, "vsDivision", season)
}

func (c *Client) GetPlayerStatsBySeasonVsDivision(id int, season string) ([]PlayerStats, int, error) {
	return c.GetPeopleStatsBySeasonVsDivision(id, season)
}

func (c *Client) GetPeopleStatsBySeasonVsConference(id int, season string) ([]PlayerStats, int, error) {
	return c.getStats(id, "vsConference", season)
}

func (c *Client) GetPlayerStatsBySeasonVsConference(id int, season string) ([]PlayerStats, int, error) {
	return c.GetPeopleStatsBySeasonVsConference(id, season)
}

func (c *Client) GetPeopleStatsBySeasonVsAllTeams(id int, season string) ([]PlayerStats, int, error) {
	return c.getStats(id, "vsTeam", season)
}

func (c *Client) GetPlayerStatsBySeasonVsAllTeams(id int, season string) ([]PlayerStats, int, error) {
	return c.GetPeopleStatsBySeasonVsAllTeams(id, season)
}

func (c *Client) GetPeopleStatsBySeasonGameLog(id int, season string) ([]PlayerStats, int, error) {
	return c.getStats(id, "gameLog", season)
}

func (c *Client) GetPlayerStatsBySeasonGameLog(id int, season string) ([]PlayerStats, int, error) {
	return c.GetPeopleStatsBySeasonGameLog(id, season)
}

func (c *Client) GetPeopleStatsByRanking(id int, season string) ([]PlayerStats, int, error) {
	return c.getStats(id, "regularSeasonStatRankings", season)
}

func (c *Client) GetPlayerStatsByRanking(id int, season string) ([]PlayerStats, int, error) {
	return c.GetPeopleStatsByRanking(id, season)
}

func (c *Client) GetPeopleStatsByGoalsByGameSituation(id int, season string) ([]PlayerStats, int, error) {
	return c.getStats(id, "goalsByGameSituation", season)
}

func (c *Client) GetPlayerStatsByGaolsByGameSituation(id int, season string) ([]PlayerStats, int, error) {
	return c.GetPeopleStatsByGoalsByGameSituation(id, season)
}

func (c *Client) GetPeopleStatsByGoalsByProjectedTotals(id int, season string) ([]PlayerStats, int, error) {
	return c.getStats(id, "goalsByGameSituation", season)
}

func (c *Client) GetPlayerStatsByGaolsByProjectedTotals(id int, season string) ([]PlayerStats, int, error) {
	return c.GetPeopleStatsByGoalsByGameSituation(id, season)
}

// TODO: Omit Empty. Refactor into more structs. Move common to types.go.

type peopleResponse struct {
	Copyright string `json:"copyright"`
	People    People `json:"people"`
}

type People struct {
	ID                 int    `json:"id"`
	FullName           string `json:"fullName"`
	Link               string `json:"link"`
	FirstName          string `json:"firstName"`
	LastName           string `json:"lastName"`
	PrimaryNumber      string `json:"primaryNumber"`
	BirthDate          string `json:"birthDate"`
	CurrentAge         int    `json:"currentAge"`
	BirthCity          string `json:"birthCity"`
	BirthStateProvince string `json:"birthStateProvince"`
	BirthCountry       string `json:"birthCountry"`
	Nationality        string `json:"nationality"`
	Height             string `json:"height"`
	Weight             int    `json:"weight"`
	Active             bool   `json:"active"`
	AlternateCaptain   bool   `json:"alternateCaptain"`
	Captain            bool   `json:"captain"`
	Rookie             bool   `json:"rookie"`
	ShootsCatches      string `json:"shootsCatches"`
	RosterStatus       string `json:"rosterStatus"`
	CurrentTeam        struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"currentTeam"`
	PrimaryPosition struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Abbreviation string `json:"abbreviation"`
	} `json:"primaryPosition"`
}

type statsResponse struct {
	Copyright string        `json:"copyright"`
	Stats     []PlayerStats `json:"stats"`
}

type PlayerStats struct {
	Type struct {
		DisplayName string `json:"displayName"`
		GameType    struct {
			ID          string `json:"id"`
			Description string `json:"description"`
			Postseason  bool   `json:"postseason"`
		} `json:"gameType"`
	} `json:"type"`
	Splits []PerSeasonStats `json:"splits"`
}

type PerSeasonStats struct {
	Season string `json:"season"`
	Stat   struct {
		TimeOnIce                   string  `json:"timeOnIce"`
		Assists                     int     `json:"assists"`
		Goals                       int     `json:"goals"`
		Pim                         int     `json:"pim"`
		Shots                       int     `json:"shots"`
		Games                       int     `json:"games"`
		Hits                        int     `json:"hits"`
		PowerPlayGoals              int     `json:"powerPlayGoals"`
		PowerPlayPoints             int     `json:"powerPlayPoints"`
		PowerPlayTimeOnIce          string  `json:"powerPlayTimeOnIce"`
		EvenTimeOnIce               string  `json:"evenTimeOnIce"`
		PenaltyMinutes              string  `json:"penaltyMinutes"`
		FaceOffPct                  float64 `json:"faceOffPct"`
		ShotPct                     float64 `json:"shotPct"`
		GameWinningGoals            int     `json:"gameWinningGoals"`
		OverTimeGoals               int     `json:"overTimeGoals"`
		ShortHandedGoals            int     `json:"shortHandedGoals"`
		ShortHandedPoints           int     `json:"shortHandedPoints"`
		ShortHandedTimeOnIce        string  `json:"shortHandedTimeOnIce"`
		Blocked                     int     `json:"blocked"`
		PlusMinus                   int     `json:"plusMinus"`
		Points                      int     `json:"points"`
		Shifts                      int     `json:"shifts"`
		TimeOnIcePerGame            string  `json:"timeOnIcePerGame"`
		EvenTimeOnIcePerGame        string  `json:"evenTimeOnIcePerGame"`
		ShortHandedTimeOnIcePerGame string  `json:"shortHandedTimeOnIcePerGame"`
		PowerPlayTimeOnIcePerGame   string  `json:"powerPlayTimeOnIcePerGame"`
		RankPowerPlayGoals          string  `json:"rankPowerPlayGoals"`
		RankBlockedShots            string  `json:"rankBlockedShots"`
		RankAssists                 string  `json:"rankAssists"`
		RankShotPct                 string  `json:"rankShotPct"`
		RankGoals                   string  `json:"rankGoals"`
		RankHits                    string  `json:"rankHits"`
		RankPenaltyMinutes          string  `json:"rankPenaltyMinutes"`
		RankShortHandedGoals        string  `json:"rankShortHandedGoals"`
		RankPlusMinus               string  `json:"rankPlusMinus"`
		RankShots                   string  `json:"rankShots"`
		RankPoints                  string  `json:"rankPoints"`
		RankOvertimeGoals           string  `json:"rankOvertimeGoals"`
		RankGamesPlayed             string  `json:"rankGamesPlayed"`
		GoalsInSecondPeriod         int     `json:"goalsInSecondPeriod"`
		GoalsInThirdPeriod          int     `json:"goalsInThirdPeriod"`
		GoalsTrailingByOne          int     `json:"goalsTrailingByOne"`
		GoalsTrailingByThreePlus    int     `json:"goalsTrailingByThreePlus"`
	} `json:"stat"`
	Team struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"team"`
	League struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"league"`
	Game struct {
		GamePk  int    `json:"gamePk"`
		Link    string `json:"link"`
		Content struct {
			Link string `json:"link"`
		}
	}
	SequenceNumber   int    `json:"sequenceNumber"`
	Date             string `json:"date"`
	IsHome           bool   `json:"isHome"`
	IsWin            bool   `json:"isWin"`
	IsOT             bool   `json:"isOT"`
	Month            int    `json:"month"`
	DayOfWeek        int    `json:"dayOfWeek"`
	OpponentDivision struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	}
	OpponentConference struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	}
	Opponent struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	}
}
