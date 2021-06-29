package gonhl

import "encoding/json"

func (c *Client) getDraft(appendURL string) (Draft, int, error) {
	var response draftResponse
	data, statusCode, err := c.get(STATS_BASE_URL + "/draft/" + appendURL)
	if err != nil {
		return response.Drafts[0], statusCode, err
	}
	json.Unmarshal(data, &response)
	return response.Drafts[0], statusCode, nil
}

func (c *Client) GetDraft() (Draft, int, error) {
	return c.getDraft("")
}

func (c *Client) GetDraftByYear(year string) (Draft, int, error) {
	return c.getDraft(year)
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
