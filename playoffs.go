package gonhl

import (
	"encoding/json"
	"strings"
)

func (c *Client) getPlayoffSeries(appendedURL string) (PlayoffSeries, int, error) {
	var response PlayoffSeries
	data, statusCode, err := c.get(RECORDS_BASE_URL + "/playoff-series/" + strings.ReplaceAll(appendedURL, " ", "+"))
	if err != nil {
		return response, statusCode, err
	}
	json.Unmarshal(data, &response)
	return response, statusCode, nil
}

func (c *Client) GetAllPlayoffSeries(appendedURL string) (PlayoffSeries, int, error) {
	return c.getPlayoffSeries("") // Seems to get from 1993 onwards
}

func (c *Client) GetPlayoffSeriesBySeriesAndYear(series string, year string) (PlayoffSeries, int, error) {
	return c.getPlayoffSeries("?cayenneExp=seriesTitle=%22" + series + "%22 and seasonId=" + year)
}

type PlayoffSeries struct {
	Data []struct {
		ID                  int    `json:"id"`
		BottomSeedWins      int    `json:"bottomSeedWins"`
		GameID              int    `json:"gameId"`
		GameNumber          int    `json:"gameNumber"`
		GameTypeID          int    `json:"gameTypeId"`
		GamesNeededToWin    int    `json:"gamesNeededToWin"`
		LengthOfSeries      int    `json:"lengthOfSeries"`
		PlayoffRound        int    `json:"playoffRound"`
		PlayoffSeriesLetter string `json:"playoffSeriesLetter"`
		SeasonID            int    `json:"seasonId"`
		SeriesTitle         string `json:"seriesTitle"`
		TopSeedWins         int    `json:"topSeedWins"`
	} `json:"data"`
	Total int `json:"total"`
}
