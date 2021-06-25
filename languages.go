package gonhl

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetLanguages() ([]Language, int, error) {
	var languages []Language
	data, statusCode, err := c.get(STATS_BASE_URL + "/languages")
	if err != nil {
		return languages, statusCode, err
	}
	json.Unmarshal(data, &languages)
	return languages, statusCode, nil
}

func (c *Client) GetLanguageByCode(languageCode string) (Language, int, error) {
	languages, statusCode, err := c.GetLanguages()
	if err != nil {
		return Language{}, statusCode, err
	}
	for _, language := range languages {
		if language.LanguageCode == languageCode {
			return language, statusCode, nil
		}
	}
	return Language{}, statusCode, fmt.Errorf("cannot find language with code of %s", languageCode)
}

type Language struct {
	LanguageCode string `json:"languageCode"`
	Locale       string `json:"locale"`
}
