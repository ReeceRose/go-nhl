package gonhl

import (
	"encoding/json"
	"strconv"
	"strings"
)

func (c *Client) getDraftFromStatsAPI(appendURL string) (Draft, int, error) {
	var response draftResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/draft/" + appendURL)
	if err != nil {
		return response.Drafts[0], statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Drafts[0], statusCode, nil
}

func (c *Client) GetDraft() (Draft, int, error) {
	return c.getDraftFromStatsAPI("")
}

func (c *Client) GetDraftByYear(year string) (Draft, int, error) {
	return c.getDraftFromStatsAPI(year)
}

func (c *Client) getDraftFromRecordsAPI(appendURL string) (DraftPicks, int, error) {
	var response DraftPicks
	data, statusCode, err := c.get(RECORDS_BASE_URL + "/draft/" + strings.ReplaceAll(appendURL, " ", "+"))
	if err != nil {
		return response, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response, statusCode, nil
}

func (c *Client) GetAllDraftPicks() (DraftPicks, int, error) {
	return c.getDraftFromRecordsAPI("")
}

func (c *Client) GetAllDraftPicksByYear(year string) (DraftPicks, int, error) {
	return c.getDraftFromRecordsAPI("?cayenneExp=draftYear=" + year)
}

func (c *Client) GetAllDraftPicksByTeamId(id int) (DraftPicks, int, error) {
	return c.getDraftFromRecordsAPI("?cayenneExp=draftedByTeamId=" + strconv.Itoa(id))
}

func (c *Client) GetAllDraftPicksByTeamIdAndYear(id int, year string) (DraftPicks, int, error) {
	return c.getDraftFromRecordsAPI("?cayenneExp=draftedByTeamId=" + strconv.Itoa(id) + " and draftYear=" + year)
}

type draftResponse struct {
	Copyright string  `json:"copyright"`
	Drafts    []Draft `json:"drafts"`
}

type Draft struct {
	DraftYear int `json:"draftYear"`
	Rounds    []struct {
		RoundNumber int    `json:"roundNumber"`
		Round       string `json:"round"`
		Picks       []struct {
			Year        int    `json:"year"`
			Round       string `json:"round"`
			PickOverall int    `json:"pickOverall"`
			PickInRound int    `json:"pickInRound"`
			Team        struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
			Prospect struct {
				ID       int    `json:"id"`
				FullName string `json:"fullName"`
				Link     string `json:"link"`
			} `json:"prospect"`
		} `json:"picks"`
	} `json:"rounds"`
}

type DraftPicks struct {
	Data []struct {
		ID                 int    `json:"id"`
		AgeInDays          int    `json:"ageInDays"`
		AgeInDaysForYear   int    `json:"ageInDaysForYear"`
		AgeInYears         int    `json:"ageInYears"`
		AmateurClubName    string `json:"amateurClubName"`
		AmateurLeague      string `json:"amateurLeague"`
		BirthDate          string `json:"birthDate"`
		BirthPlace         string `json:"birthPlace"`
		CountryCode        string `json:"countryCode"`
		CsPlayerID         int    `json:"csPlayerId"`
		DraftDate          string `json:"draftDate"`
		DraftMasterID      int    `json:"draftMasterId"`
		DraftYear          int    `json:"draftYear"`
		DraftedByTeamID    int    `json:"draftedByTeamId"`
		FirstName          string `json:"firstName"`
		Height             int    `json:"height"`
		LastName           string `json:"lastName"`
		Notes              string `json:"notes"`
		OverallPickNumber  int    `json:"overallPickNumber"`
		PickInRound        int    `json:"pickInRound"`
		PlayerID           int    `json:"playerId"`
		PlayerName         string `json:"playerName"`
		Position           string `json:"position"`
		RemovedOutright    string `json:"removedOutright"`
		RemovedOutrightWhy string `json:"removedOutrightWhy"`
		RoundNumber        int    `json:"roundNumber"`
		ShootsCatches      string `json:"shootsCatches"`
		SupplementalDraft  string `json:"supplementalDraft"`
		TeamPickHistory    string `json:"teamPickHistory"`
		TriCode            string `json:"triCode"`
		Weight             int    `json:"weight"`
	} `json:"data"`
	Total int `json:"total"`
}
