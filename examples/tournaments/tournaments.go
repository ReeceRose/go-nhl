package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Tournaments")
	client := nhl.NewClient()
	fmt.Println("Attempting to get tournament (playoffs)")
	tournament, statusCode, err := client.GetTournamentByName("playoffs")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetTournamentByName returned a status code of %d\n", statusCode)
	fmt.Printf("Got the %s tournament\n", tournament.Name)
}
