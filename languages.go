package gonhl

import "encoding/json"

func (c *Client) GetLanguages() ([]Language, int, error) {
	var languages []Language
	data, statusCode, err := c.get(STATS_BASE_URL + "/languages")
	if err != nil {
		return languages, statusCode, err
	}
	json.Unmarshal(data, &languages)
	return languages, statusCode, nil
}

type Language struct {
	LanguageCode string `json:"languageCode"`
	Locale       string `json:"locale"`
}
