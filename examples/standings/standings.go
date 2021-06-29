package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Standings")
	client := nhl.NewClient()
	fmt.Println("Attempting to get standings")
	standings, statusCode, err := client.GetStandings()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetStandings returned a status code of %d\n", statusCode)
	for _, standing := range standings {
		fmt.Printf("%s:%s\n", standing.StandingsType, standing.Division.Name)
	}
}
