package gonhl

import (
	"encoding/json"
	"time"
)

func (c *Client) getStandings(appendedURL string) ([]Record, int, error) {
	var response standingsReponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/standings/" + appendedURL)
	if err != nil {
		return response.Records, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Records, statusCode, nil
}

func (c *Client) GetStandings() ([]Record, int, error) {
	return c.getStandings("")
}

func (c *Client) GetStandingsByYear(year string) (Record, int, error) {
	records, responseCode, error := c.getStandings("?season=" + year)
	return records[0], responseCode, error
}

func (c *Client) GetDetailedStandingsByYear(year string) (Record, int, error) {
	records, responseCode, error := c.getStandings("?season=" + year + "&expand=standings.record")
	return records[0], responseCode, error
}

func (c *Client) GetDetailedStandingsByYar(year string) (Record, int, error) {
	records, responseCode, error := c.getStandings("?expand=standings.record")
	return records[0], responseCode, error
}

type standingsReponse struct {
	Copyright string   `json:"copyright"`
	Records   []Record `json:"records"`
}

type Record struct {
	StandingsType string `json:"standingsType"`
	League        struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"league"`
	Division struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		NameShort    string `json:"nameShort"`
		Link         string `json:"link"`
		Abbreviation string `json:"abbreviation"`
	} `json:"division"`
	Conference struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"conference"`
	Season      string `json:"season"`
	TeamRecords []struct {
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"team"`
		LeagueRecord struct {
			Wins   int    `json:"wins"`
			Losses int    `json:"losses"`
			Ot     int    `json:"ot"`
			Type   string `json:"type"`
		} `json:"leagueRecord"`
		RegulationWins     int    `json:"regulationWins"`
		GoalsAgainst       int    `json:"goalsAgainst"`
		GoalsScored        int    `json:"goalsScored"`
		Points             int    `json:"points"`
		DivisionRank       string `json:"divisionRank"`
		DivisionL10Rank    string `json:"divisionL10Rank"`
		DivisionRoadRank   string `json:"divisionRoadRank"`
		DivisionHomeRank   string `json:"divisionHomeRank"`
		ConferenceRank     string `json:"conferenceRank"`
		ConferenceL10Rank  string `json:"conferenceL10Rank"`
		ConferenceRoadRank string `json:"conferenceRoadRank"`
		ConferenceHomeRank string `json:"conferenceHomeRank"`
		LeagueRank         string `json:"leagueRank"`
		LeagueL10Rank      string `json:"leagueL10Rank"`
		LeagueRoadRank     string `json:"leagueRoadRank"`
		LeagueHomeRank     string `json:"leagueHomeRank"`
		WildCardRank       string `json:"wildCardRank"`
		Row                int    `json:"row"`
		GamesPlayed        int    `json:"gamesPlayed"`
		Streak             struct {
			StreakType   string `json:"streakType"`
			StreakNumber int    `json:"streakNumber"`
			StreakCode   string `json:"streakCode"`
		} `json:"streak"`
		PointsPercentage float64 `json:"pointsPercentage"`
		PpDivisionRank   string  `json:"ppDivisionRank"`
		PpConferenceRank string  `json:"ppConferenceRank"`
		PpLeagueRank     string  `json:"ppLeagueRank"`
		Records          struct {
			DivisionRecords []struct {
				Wins   int    `json:"wins"`
				Losses int    `json:"losses"`
				Ot     int    `json:"ot"`
				Type   string `json:"type"`
			} `json:"divisionRecords"`
			OverallRecords []struct {
				Wins   int    `json:"wins"`
				Losses int    `json:"losses"`
				Ot     int    `json:"ot,omitempty"`
				Type   string `json:"type"`
			} `json:"overallRecords"`
			ConferenceRecords []struct {
				Wins   int    `json:"wins"`
				Losses int    `json:"losses"`
				Ot     int    `json:"ot"`
				Type   string `json:"type"`
			} `json:"conferenceRecords"`
		} `json:"records"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"teamRecords"`
}
