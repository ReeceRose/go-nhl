package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Conferences")
	client := nhl.NewClient()
	fmt.Println("Attempting to get conferences")
	conferences, statusCode, err := client.GetConferences()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetConferences returned a status code of %d\n", statusCode)
	fmt.Printf("Got a total of %d conferences\n", len(conferences))

	fmt.Println("Attempting to get conference by ID")
	conference, statusCode, err := client.GetConferenceById(5) // Western
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetConferenceById returned a status code of %d\n", statusCode)
	fmt.Printf("Got conference with the name: %v\n", conference.Name)
}
