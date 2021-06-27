package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	// Note, this is a little misleading. It doesn't give the status of a game, rather the status
	// any game can be in.
	fmt.Println("Go NHL Wrapper Example: Game Statuses")
	client := nhl.NewClient()
	fmt.Println("Attempting to get game statuses")
	gameStatuses, statusCode, err := client.GetGameStatuses()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetGameStatuses returned a status code of %d\n", statusCode)
	for _, gameStatus := range gameStatuses {
		fmt.Printf("%s:%s:%s\n", gameStatus.Code, gameStatus.AbstractGameState, gameStatus.DetailedState)
	}
}
