package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Roster")
	client := nhl.NewClient()
	fmt.Println("Attempting to get St. Louis Blues roster")
	roster, statusCode, err := client.GetRoster(19)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetRoster returned a status code of %d\n", statusCode)
	fmt.Printf("Got a total of %d players on the roster\n", len(roster))
}
