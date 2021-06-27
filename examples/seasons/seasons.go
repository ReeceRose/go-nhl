package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Seasons")
	client := nhl.NewClient()
	fmt.Println("Attempting to get current season")
	season, statusCode, err := client.GetCurrentSeason()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetCurrentSeason returned a status code of %d\n", statusCode)
	fmt.Printf("Current season ends on %s\n", season.SeasonEndDate)
}
