package gonhl

import (
	"encoding/json"
	"strconv"
	"time"
)

func (c *Client) getGame(id int, appendedURL string) (Game, int, error) {
	var game Game
	data, statusCode, err := c.get(STATS_BASE_URL + "/game/" + strconv.Itoa(id) + "/")
	if err != nil {
		return game, statusCode, err
	}
	json.Unmarshal(data, &game)
	return game, statusCode, nil
}

func (c *Client) GetLiveGameFeed(id int) (Game, int, error) {
	return c.getGame(id, "live")
}

func (c *Client) GetGameBoxscore(id int) (Game, int, error) {
	return c.getGame(id, "boxscore")
}

type Game struct {
	Copyright string `json:"copyright"`
	GamePk    int    `json:"gamePk"`
	Link      string `json:"link"`
	MetaData  struct {
		Wait      int    `json:"wait"`
		TimeStamp string `json:"timeStamp"`
	} `json:"metaData"`
	GameData struct {
		Game struct {
			Pk     int    `json:"pk"`
			Season string `json:"season"`
			Type   string `json:"type"`
		} `json:"game"`
		Datetime struct {
			DateTime    time.Time `json:"dateTime"`
			EndDateTime time.Time `json:"endDateTime"`
		} `json:"datetime"`
		Status struct {
			AbstractGameState string `json:"abstractGameState"`
			CodedGameState    string `json:"codedGameState"`
			DetailedState     string `json:"detailedState"`
			StatusCode        string `json:"statusCode"`
			StartTimeTBD      bool   `json:"startTimeTBD"`
		} `json:"status"`
		Teams struct {
			Away struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Link  string `json:"link"`
				Venue struct {
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
				TriCode         string `json:"triCode"`
				TeamName        string `json:"teamName"`
				LocationName    string `json:"locationName"`
				FirstYearOfPlay string `json:"firstYearOfPlay"`
				Division        struct {
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
				Franchise struct {
					FranchiseID int    `json:"franchiseId"`
					TeamName    string `json:"teamName"`
					Link        string `json:"link"`
				} `json:"franchise"`
				ShortName       string `json:"shortName"`
				OfficialSiteURL string `json:"officialSiteUrl"`
				FranchiseID     int    `json:"franchiseId"`
				Active          bool   `json:"active"`
			} `json:"away"`
			Home struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Link  string `json:"link"`
				Venue struct {
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
				TriCode         string `json:"triCode"`
				TeamName        string `json:"teamName"`
				LocationName    string `json:"locationName"`
				FirstYearOfPlay string `json:"firstYearOfPlay"`
				Division        struct {
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
				Franchise struct {
					FranchiseID int    `json:"franchiseId"`
					TeamName    string `json:"teamName"`
					Link        string `json:"link"`
				} `json:"franchise"`
				ShortName       string `json:"shortName"`
				OfficialSiteURL string `json:"officialSiteUrl"`
				FranchiseID     int    `json:"franchiseId"`
				Active          bool   `json:"active"`
			} `json:"home"`
		} `json:"teams"`
		Players struct {
			Player struct {
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
					ID      int    `json:"id"`
					Name    string `json:"name"`
					Link    string `json:"link"`
					TriCode string `json:"triCode"`
				} `json:"currentTeam"`
				PrimaryPosition struct {
					Code         string `json:"code"`
					Name         string `json:"name"`
					Type         string `json:"type"`
					Abbreviation string `json:"abbreviation"`
				} `json:"primaryPosition"`
			}
		} `json:"players"`
		Venue struct {
			Name string `json:"name"`
			Link string `json:"link"`
		} `json:"venue"`
	} `json:"gameData"`
	LiveData struct {
		Plays struct {
			AllPlays []struct {
				Players []struct {
					Player struct {
						ID       int    `json:"id"`
						FullName string `json:"fullName"`
						Link     string `json:"link"`
					} `json:"player"`
					PlayerType string `json:"playerType"`
				} `json:"players"`
				Result struct {
					Event       string `json:"event"`
					EventCode   string `json:"eventCode"`
					EventTypeID string `json:"eventTypeId"`
					Description string `json:"description"`
				} `json:"result"`
				About struct {
					EventIdx            int       `json:"eventIdx"`
					EventID             int       `json:"eventId"`
					Period              int       `json:"period"`
					PeriodType          string    `json:"periodType"`
					OrdinalNum          string    `json:"ordinalNum"`
					PeriodTime          string    `json:"periodTime"`
					PeriodTimeRemaining string    `json:"periodTimeRemaining"`
					DateTime            time.Time `json:"dateTime"`
					Goals               struct {
						Away int `json:"away"`
						Home int `json:"home"`
					} `json:"goals"`
				} `json:"about"`
				Coordinates struct {
					X float64 `json:"x"`
					Y int     `json:"y"`
				} `json:"coordinates"`
				Team struct {
					ID      int    `json:"id"`
					Name    string `json:"name"`
					Link    string `json:"link"`
					TriCode string `json:"triCode"`
				} `json:"team"`
			} `json:"allPlays"`
			ScoringPlays  []int `json:"scoringPlays"`
			PenaltyPlays  []int `json:"penaltyPlays"`
			PlaysByPeriod []struct {
				StartIndex int   `json:"startIndex"`
				Plays      []int `json:"plays"`
				EndIndex   int   `json:"endIndex"`
			} `json:"playsByPeriod"`
			CurrentPlay struct {
				Result struct {
					Event       string `json:"event"`
					EventCode   string `json:"eventCode"`
					EventTypeID string `json:"eventTypeId"`
					Description string `json:"description"`
				} `json:"result"`
				About struct {
					EventIdx            int       `json:"eventIdx"`
					EventID             int       `json:"eventId"`
					Period              int       `json:"period"`
					PeriodType          string    `json:"periodType"`
					OrdinalNum          string    `json:"ordinalNum"`
					PeriodTime          string    `json:"periodTime"`
					PeriodTimeRemaining string    `json:"periodTimeRemaining"`
					DateTime            time.Time `json:"dateTime"`
					Goals               struct {
						Away int `json:"away"`
						Home int `json:"home"`
					} `json:"goals"`
				} `json:"about"`
				Coordinates struct {
				} `json:"coordinates"`
			} `json:"currentPlay"`
		} `json:"plays"`
		Linescore struct {
			CurrentPeriod              int    `json:"currentPeriod"`
			CurrentPeriodOrdinal       string `json:"currentPeriodOrdinal"`
			CurrentPeriodTimeRemaining string `json:"currentPeriodTimeRemaining"`
			Periods                    []struct {
				PeriodType string    `json:"periodType"`
				StartTime  time.Time `json:"startTime"`
				EndTime    time.Time `json:"endTime"`
				Num        int       `json:"num"`
				OrdinalNum string    `json:"ordinalNum"`
				Home       struct {
					Goals       int    `json:"goals"`
					ShotsOnGoal int    `json:"shotsOnGoal"`
					RinkSide    string `json:"rinkSide"`
				} `json:"home"`
				Away struct {
					Goals       int    `json:"goals"`
					ShotsOnGoal int    `json:"shotsOnGoal"`
					RinkSide    string `json:"rinkSide"`
				} `json:"away"`
			} `json:"periods"`
			ShootoutInfo struct {
				Away struct {
					Scores   int `json:"scores"`
					Attempts int `json:"attempts"`
				} `json:"away"`
				Home struct {
					Scores   int `json:"scores"`
					Attempts int `json:"attempts"`
				} `json:"home"`
			} `json:"shootoutInfo"`
			Teams struct {
				Home struct {
					Team struct {
						ID           int    `json:"id"`
						Name         string `json:"name"`
						Link         string `json:"link"`
						Abbreviation string `json:"abbreviation"`
						TriCode      string `json:"triCode"`
					} `json:"team"`
					Goals        int  `json:"goals"`
					ShotsOnGoal  int  `json:"shotsOnGoal"`
					GoaliePulled bool `json:"goaliePulled"`
					NumSkaters   int  `json:"numSkaters"`
					PowerPlay    bool `json:"powerPlay"`
				} `json:"home"`
				Away struct {
					Team struct {
						ID           int    `json:"id"`
						Name         string `json:"name"`
						Link         string `json:"link"`
						Abbreviation string `json:"abbreviation"`
						TriCode      string `json:"triCode"`
					} `json:"team"`
					Goals        int  `json:"goals"`
					ShotsOnGoal  int  `json:"shotsOnGoal"`
					GoaliePulled bool `json:"goaliePulled"`
					NumSkaters   int  `json:"numSkaters"`
					PowerPlay    bool `json:"powerPlay"`
				} `json:"away"`
			} `json:"teams"`
			PowerPlayStrength string `json:"powerPlayStrength"`
			HasShootout       bool   `json:"hasShootout"`
			IntermissionInfo  struct {
				IntermissionTimeRemaining int  `json:"intermissionTimeRemaining"`
				IntermissionTimeElapsed   int  `json:"intermissionTimeElapsed"`
				InIntermission            bool `json:"inIntermission"`
			} `json:"intermissionInfo"`
			PowerPlayInfo struct {
				SituationTimeRemaining int  `json:"situationTimeRemaining"`
				SituationTimeElapsed   int  `json:"situationTimeElapsed"`
				InSituation            bool `json:"inSituation"`
			} `json:"powerPlayInfo"`
		} `json:"linescore"`
		Boxscore struct {
			Teams struct {
				Away struct {
					Team struct {
						ID           int    `json:"id"`
						Name         string `json:"name"`
						Link         string `json:"link"`
						Abbreviation string `json:"abbreviation"`
						TriCode      string `json:"triCode"`
					} `json:"team"`
					TeamStats struct {
						TeamSkaterStats struct {
							Goals                  int     `json:"goals"`
							Pim                    int     `json:"pim"`
							Shots                  int     `json:"shots"`
							PowerPlayPercentage    string  `json:"powerPlayPercentage"`
							PowerPlayGoals         float64 `json:"powerPlayGoals"`
							PowerPlayOpportunities float64 `json:"powerPlayOpportunities"`
							FaceOffWinPercentage   string  `json:"faceOffWinPercentage"`
							Blocked                int     `json:"blocked"`
							Takeaways              int     `json:"takeaways"`
							Giveaways              int     `json:"giveaways"`
							Hits                   int     `json:"hits"`
						} `json:"teamSkaterStats"`
					} `json:"teamStats"`
					Players struct {
						Player struct {
							Person struct {
								ID            int    `json:"id"`
								FullName      string `json:"fullName"`
								Link          string `json:"link"`
								ShootsCatches string `json:"shootsCatches"`
								RosterStatus  string `json:"rosterStatus"`
							} `json:"person"`
							JerseyNumber string `json:"jerseyNumber"`
							Position     struct {
								Code         string `json:"code"`
								Name         string `json:"name"`
								Type         string `json:"type"`
								Abbreviation string `json:"abbreviation"`
							} `json:"position"`
							Stats struct {
								GoalieStats struct {
									TimeOnIce                  string  `json:"timeOnIce"`
									Assists                    int     `json:"assists"`
									Goals                      int     `json:"goals"`
									Pim                        int     `json:"pim"`
									Shots                      int     `json:"shots"`
									Saves                      int     `json:"saves"`
									PowerPlaySaves             int     `json:"powerPlaySaves"`
									ShortHandedSaves           int     `json:"shortHandedSaves"`
									EvenSaves                  int     `json:"evenSaves"`
									ShortHandedShotsAgainst    int     `json:"shortHandedShotsAgainst"`
									EvenShotsAgainst           int     `json:"evenShotsAgainst"`
									PowerPlayShotsAgainst      int     `json:"powerPlayShotsAgainst"`
									Decision                   string  `json:"decision"`
									SavePercentage             float64 `json:"savePercentage"`
									PowerPlaySavePercentage    float64 `json:"powerPlaySavePercentage"`
									EvenStrengthSavePercentage float64 `json:"evenStrengthSavePercentage"`
								} `json:"goalieStats"`
							} `json:"stats"`
						}
					} `json:"players"`
					Goalies   []int `json:"goalies"`
					Skaters   []int `json:"skaters"`
					OnIce     []int `json:"onIce"`
					OnIcePlus []struct {
						PlayerID      int `json:"playerId"`
						ShiftDuration int `json:"shiftDuration"`
						Stamina       int `json:"stamina"`
					} `json:"onIcePlus"`
					Scratches  []int         `json:"scratches"`
					PenaltyBox []interface{} `json:"penaltyBox"`
					Coaches    []struct {
						Person struct {
							FullName string `json:"fullName"`
							Link     string `json:"link"`
						} `json:"person"`
						Position struct {
							Code         string `json:"code"`
							Name         string `json:"name"`
							Type         string `json:"type"`
							Abbreviation string `json:"abbreviation"`
						} `json:"position"`
					} `json:"coaches"`
				} `json:"away"`
				Home struct {
					Team struct {
						ID           int    `json:"id"`
						Name         string `json:"name"`
						Link         string `json:"link"`
						Abbreviation string `json:"abbreviation"`
						TriCode      string `json:"triCode"`
					} `json:"team"`
					TeamStats struct {
						TeamSkaterStats struct {
							Goals                  int     `json:"goals"`
							Pim                    int     `json:"pim"`
							Shots                  int     `json:"shots"`
							PowerPlayPercentage    string  `json:"powerPlayPercentage"`
							PowerPlayGoals         float64 `json:"powerPlayGoals"`
							PowerPlayOpportunities float64 `json:"powerPlayOpportunities"`
							FaceOffWinPercentage   string  `json:"faceOffWinPercentage"`
							Blocked                int     `json:"blocked"`
							Takeaways              int     `json:"takeaways"`
							Giveaways              int     `json:"giveaways"`
							Hits                   int     `json:"hits"`
						} `json:"teamSkaterStats"`
					} `json:"teamStats"`
					Players struct {
						Player struct {
							Person struct {
								ID            int    `json:"id"`
								FullName      string `json:"fullName"`
								Link          string `json:"link"`
								ShootsCatches string `json:"shootsCatches"`
								RosterStatus  string `json:"rosterStatus"`
							} `json:"person"`
							JerseyNumber string `json:"jerseyNumber"`
							Position     struct {
								Code         string `json:"code"`
								Name         string `json:"name"`
								Type         string `json:"type"`
								Abbreviation string `json:"abbreviation"`
							} `json:"position"`
							Stats struct {
								SkaterStats struct {
									TimeOnIce            string `json:"timeOnIce"`
									Assists              int    `json:"assists"`
									Goals                int    `json:"goals"`
									Shots                int    `json:"shots"`
									Hits                 int    `json:"hits"`
									PowerPlayGoals       int    `json:"powerPlayGoals"`
									PowerPlayAssists     int    `json:"powerPlayAssists"`
									PenaltyMinutes       int    `json:"penaltyMinutes"`
									FaceOffWins          int    `json:"faceOffWins"`
									FaceoffTaken         int    `json:"faceoffTaken"`
									Takeaways            int    `json:"takeaways"`
									Giveaways            int    `json:"giveaways"`
									ShortHandedGoals     int    `json:"shortHandedGoals"`
									ShortHandedAssists   int    `json:"shortHandedAssists"`
									Blocked              int    `json:"blocked"`
									PlusMinus            int    `json:"plusMinus"`
									EvenTimeOnIce        string `json:"evenTimeOnIce"`
									PowerPlayTimeOnIce   string `json:"powerPlayTimeOnIce"`
									ShortHandedTimeOnIce string `json:"shortHandedTimeOnIce"`
								} `json:"skaterStats"`
							} `json:"stats"`
						}
					} `json:"players"`
					Goalies   []int `json:"goalies"`
					Skaters   []int `json:"skaters"`
					OnIce     []int `json:"onIce"`
					OnIcePlus []struct {
						PlayerID      int `json:"playerId"`
						ShiftDuration int `json:"shiftDuration"`
						Stamina       int `json:"stamina"`
					} `json:"onIcePlus"`
					Scratches  []int         `json:"scratches"`
					PenaltyBox []interface{} `json:"penaltyBox"`
					Coaches    []struct {
						Person struct {
							FullName string `json:"fullName"`
							Link     string `json:"link"`
						} `json:"person"`
						Position struct {
							Code         string `json:"code"`
							Name         string `json:"name"`
							Type         string `json:"type"`
							Abbreviation string `json:"abbreviation"`
						} `json:"position"`
					} `json:"coaches"`
				} `json:"home"`
			} `json:"teams"`
			Officials []struct {
				Official struct {
					ID       int    `json:"id"`
					FullName string `json:"fullName"`
					Link     string `json:"link"`
				} `json:"official"`
				OfficialType string `json:"officialType"`
			} `json:"officials"`
		} `json:"boxscore"`
		Decisions struct {
			Winner struct {
				ID       int    `json:"id"`
				FullName string `json:"fullName"`
				Link     string `json:"link"`
			} `json:"winner"`
			Loser struct {
				ID       int    `json:"id"`
				FullName string `json:"fullName"`
				Link     string `json:"link"`
			} `json:"loser"`
			FirstStar struct {
				ID       int    `json:"id"`
				FullName string `json:"fullName"`
				Link     string `json:"link"`
			} `json:"firstStar"`
			SecondStar struct {
				ID       int    `json:"id"`
				FullName string `json:"fullName"`
				Link     string `json:"link"`
			} `json:"secondStar"`
			ThirdStar struct {
				ID       int    `json:"id"`
				FullName string `json:"fullName"`
				Link     string `json:"link"`
			} `json:"thirdStar"`
		} `json:"decisions"`
	} `json:"liveData"`
}
