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

func (c *Client) GetTeamStatsById(id int) ([]Stat, int, error) {
	var response teamStatsResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/teams/" + strconv.Itoa(id) + "/stats")
	if err != nil {
		return response.Stats, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Stats, statusCode, nil
}

type teamStatsResponse struct {
	Copyright string `json:"copyright"`
	Stats     []Stat `json:"stats"`
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

type Stat struct {
	Type struct {
		DisplayName string `json:"displayName"`
		GameType    struct {
			ID          string `json:"id"`
			Description string `json:"description"`
			Postseason  bool   `json:"postseason"`
		} `json:"gameType"`
	} `json:"type"`
	Splits []struct {
		Stat struct {
			GamesPlayed            int     `json:"gamesPlayed"`
			Wins                   int     `json:"wins"`
			Losses                 int     `json:"losses"`
			Ot                     int     `json:"ot"`
			Pts                    int     `json:"pts"`
			PtPctg                 string  `json:"ptPctg"`
			GoalsPerGame           float64 `json:"goalsPerGame"`
			GoalsAgainstPerGame    float64 `json:"goalsAgainstPerGame"`
			EvGGARatio             float64 `json:"evGGARatio"`
			PowerPlayPercentage    string  `json:"powerPlayPercentage"`
			PowerPlayGoals         float64 `json:"powerPlayGoals"`
			PowerPlayGoalsAgainst  float64 `json:"powerPlayGoalsAgainst"`
			PowerPlayOpportunities float64 `json:"powerPlayOpportunities"`
			PenaltyKillPercentage  string  `json:"penaltyKillPercentage"`
			ShotsPerGame           float64 `json:"shotsPerGame"`
			ShotsAllowed           float64 `json:"shotsAllowed"`
			WinScoreFirst          float64 `json:"winScoreFirst"`
			WinOppScoreFirst       float64 `json:"winOppScoreFirst"`
			WinLeadFirstPer        float64 `json:"winLeadFirstPer"`
			WinLeadSecondPer       float64 `json:"winLeadSecondPer"`
			WinOutshootOpp         float64 `json:"winOutshootOpp"`
			WinOutshotByOpp        float64 `json:"winOutshotByOpp"`
			FaceOffsTaken          float64 `json:"faceOffsTaken"`
			FaceOffsWon            float64 `json:"faceOffsWon"`
			FaceOffsLost           float64 `json:"faceOffsLost"`
			FaceOffWinPercentage   string  `json:"faceOffWinPercentage"`
			ShootingPctg           float64 `json:"shootingPctg"`
			SavePctg               float64 `json:"savePctg"`
		} `json:"stat"`
		Team struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"team"`
	} `json:"splits"`
}
