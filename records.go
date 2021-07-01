package gonhl

import (
	"encoding/json"
	"strconv"
)

func (c *Client) getAllTimeRecordVsFranchises(appendedURL string) (RecordVsFranchises, int, error) {
	var response RecordVsFranchises
	data, statusCode, err := c.get(RECORDS_BASE_URL + "/all-time-record-vs-franchise/" + appendedURL)
	if err != nil {
		return response, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response, statusCode, nil
}

func (c *Client) GetAllTimeRecordVsFranchises() (RecordVsFranchises, int, error) {
	return c.getAllTimeRecordVsFranchises("")
}

func (c *Client) GetAllTimeRecordVsFranchiseById(id int) (RecordVsFranchises, int, error) {
	return c.getAllTimeRecordVsFranchises("?cayenneExp=teamFranchiseId=" + strconv.Itoa(id))
}

func (c *Client) getPlayoffRecordVsFranchises(appendedURL string) (PlayoffRecordVsFranchises, int, error) {
	var response PlayoffRecordVsFranchises
	data, statusCode, err := c.get(RECORDS_BASE_URL + "/platoff-franchise-vs-franchise/" + appendedURL)
	if err != nil {
		return response, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response, statusCode, nil
}

func (c *Client) GetPlayoffRecordVsFranchises() (PlayoffRecordVsFranchises, int, error) {
	return c.getPlayoffRecordVsFranchises("")
}

func (c *Client) GetPlayoffRecordVsFranchiseById(id int) (PlayoffRecordVsFranchises, int, error) {
	return c.getPlayoffRecordVsFranchises("?cayenneExp=teamFranchiseId=" + strconv.Itoa(id))
}

type RecordVsFranchises struct {
	Data []struct {
		ID                       int    `json:"id"`
		ActiveFranchise          int    `json:"activeFranchise"`
		ActiveOpponentFranchise  int    `json:"activeOpponentFranchise"`
		FranchiseName            string `json:"franchiseName"`
		GameTypeID               int    `json:"gameTypeId"`
		HomeGamesPlayed          int    `json:"homeGamesPlayed"`
		HomeGoalsAgainst         int    `json:"homeGoalsAgainst"`
		HomeGoalsFor             int    `json:"homeGoalsFor"`
		HomeLastMeetingSeasonID  int    `json:"homeLastMeetingSeasonId"`
		HomeLosses               int    `json:"homeLosses"`
		HomeOtLosses             int    `json:"homeOtLosses"`
		HomePoints               int    `json:"homePoints"`
		HomeTies                 int    `json:"homeTies"`
		HomeWins                 int    `json:"homeWins"`
		OpponentFranchiseID      int    `json:"opponentFranchiseId"`
		OpponentFranchiseName    string `json:"opponentFranchiseName"`
		OpponentTeamID           int    `json:"opponentTeamId"`
		RoadGamesPlayed          int    `json:"roadGamesPlayed"`
		RoadGoalsAgainst         int    `json:"roadGoalsAgainst"`
		RoadGoalsFor             int    `json:"roadGoalsFor"`
		RoadLastMeetingSeasonID  int    `json:"roadLastMeetingSeasonId"`
		RoadLosses               int    `json:"roadLosses"`
		RoadOtLosses             int    `json:"roadOtLosses"`
		RoadPoints               int    `json:"roadPoints"`
		RoadTies                 int    `json:"roadTies"`
		RoadWins                 int    `json:"roadWins"`
		TeamFranchiseID          int    `json:"teamFranchiseId"`
		TeamID                   int    `json:"teamId"`
		TotalGamesPlayed         int    `json:"totalGamesPlayed"`
		TotalGoalsAgainst        int    `json:"totalGoalsAgainst"`
		TotalGoalsFor            int    `json:"totalGoalsFor"`
		TotalLastMeetingSeasonID int    `json:"totalLastMeetingSeasonId"`
		TotalLosses              int    `json:"totalLosses"`
		TotalOtLosses            int    `json:"totalOtLosses"`
		TotalPoints              int    `json:"totalPoints"`
		TotalTies                int    `json:"totalTies"`
		TotalWins                int    `json:"totalWins"`
	} `json:"data"`
	Total int `json:"total"`
}

type PlayoffRecordVsFranchises struct {
	Data []struct {
		ID                      int    `json:"id"`
		ActiveFranchise         int    `json:"activeFranchise"`
		ActiveOpponentFranchise int    `json:"activeOpponentFranchise"`
		GamesPlayed             int    `json:"gamesPlayed"`
		GoalsAgainst            int    `json:"goalsAgainst"`
		GoalsFor                int    `json:"goalsFor"`
		LastMeetingLosses       int    `json:"lastMeetingLosses"`
		LastMeetingResult       string `json:"lastMeetingResult"`
		LastMeetingSeasonID     int    `json:"lastMeetingSeasonId"`
		LastMeetingWins         int    `json:"lastMeetingWins"`
		Losses                  int    `json:"losses"`
		OpponentFranchiseID     int    `json:"opponentFranchiseId"`
		OpponentTeamID          int    `json:"opponentTeamId"`
		OpponentTeamName        string `json:"opponentTeamName"`
		OvertimeLosses          int    `json:"overtimeLosses"`
		SeriesAbbrev            string `json:"seriesAbbrev"`
		SeriesLosses            int    `json:"seriesLosses"`
		SeriesTitle             string `json:"seriesTitle"`
		SeriesVsOpponent        int    `json:"seriesVsOpponent"`
		SeriesWins              int    `json:"seriesWins"`
		TeamFranchiseID         int    `json:"teamFranchiseId"`
		TeamID                  int    `json:"teamId"`
		TeamName                string `json:"teamName"`
		Ties                    int    `json:"ties"`
		Wins                    int    `json:"wins"`
	} `json:"data"`
	Total int `json:"total"`
}
