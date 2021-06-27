package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Standings Types")
	client := nhl.NewClient()
	fmt.Println("Attempting to get standings types")
	standingsTypes, statusCode, err := client.GetStandingsTypes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetStandingsTypes returned a status code of %d\n", statusCode)
	for _, standingType := range standingsTypes {
		fmt.Printf("%s:%s\n", standingType.Name, standingType.Description)
	}
}
