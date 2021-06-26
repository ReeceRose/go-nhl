package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Game Types")
	client := nhl.NewClient()
	fmt.Println("Attempting to get game types")
	gameTypes, statusCode, err := client.GetGameTypes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetGameTypes returned a status code of %d\n", statusCode)
	for _, gameType := range gameTypes {
		fmt.Printf("%s:%s:%t\n", gameType.ID, gameType.Description, gameType.Postseason)
	}

	fmt.Println("Attempting to get game type by ID")
	gameType, statusCode, err := client.GetGameTypeById("O")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetGameTypeById returned a status code of %d\n", statusCode)
	fmt.Printf("Got game type of %s\n", gameType.Description)
}
