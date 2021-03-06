package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Event Types")
	client := nhl.NewClient()
	fmt.Println("Attempting to get event types")
	eventTypes, statusCode, err := client.GetEventTypes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetEventTypes returned a status code of %d\n", statusCode)
	for _, event := range eventTypes {
		fmt.Printf("%s:%s\n", event.ID, event.Description)
	}
}
