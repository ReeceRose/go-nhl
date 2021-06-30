package gonhl

import "encoding/json"

func (c *Client) GetTrophies() (Trophies, int, error) {
	var trophies Trophies
	data, statusCode, err := c.get(RECORDS_BASE_URL + "/trophy")
	if err != nil {
		return trophies, statusCode, err
	}
	json.Unmarshal(data, &trophies)
	return trophies, statusCode, nil
}

type Trophies struct {
	Data []struct {
		ID               int         `json:"id"`
		BriefDescription string      `json:"briefDescription"`
		CategoryID       int         `json:"categoryId"`
		CreatedOn        string      `json:"createdOn"`
		Description      string      `json:"description"`
		Footnote         interface{} `json:"footnote"`
		HomePageURL      string      `json:"homePageUrl"`
		ImageURL         string      `json:"imageUrl"`
		Name             string      `json:"name"`
		ShortName        string      `json:"shortName"`
	} `json:"data"`
	Total int `json:"total"`
}
