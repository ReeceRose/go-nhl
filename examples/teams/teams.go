package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Teams")
	client := nhl.NewClient()
	fmt.Println("Attempting to get teams")
	teams, statusCode, err := client.GetTeams()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetTeams returned a status code of %d\n", statusCode)
	fmt.Printf("Got a total of %d teams\n", len(teams))

	fmt.Println("Attempting to get team by ID")
	team, statusCode, err := client.GetTeamById(19) // St. Louis Blues
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetTeamById returned a status code of %d\n", statusCode)
	fmt.Printf("Got team with the name: %v\n", team.Name)
}
