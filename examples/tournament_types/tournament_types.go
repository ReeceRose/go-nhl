package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Tournament Types")
	client := nhl.NewClient()
	fmt.Println("Attempting to get tournament types")
	tournamentTypes, statusCode, err := client.GetTournamentTypes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetTournamentTypes returned a status code of %d\n", statusCode)
	fmt.Printf("Got a total of %d tournament types\n", len(tournamentTypes))
}
