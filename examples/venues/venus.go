package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Venues")
	client := nhl.NewClient()
	fmt.Println("Attempting to get venues")
	venues, statusCode, err := client.GetVenues()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetVenues returned a status code of %d\n", statusCode)
	fmt.Printf("Got a total of %d venues\n", len(venues))
}
