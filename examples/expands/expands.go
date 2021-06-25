package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Expands")
	client := nhl.NewClient()
	fmt.Println("Attempting to get expands")
	expands, statusCode, err := client.GetExpands()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetExpands returned a status code of %d\n", statusCode)
	for _, expand := range expands {
		fmt.Printf("%s:%s\n", expand.Name, expand.Description)
	}

	fmt.Println("Attempting to get expand by name")
	expand, statusCode, err := client.GetExpandByName("person.currentTeam")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetExpandByName returned a status code of %d\n", statusCode)
	fmt.Printf("Got expand of %s\n", expand.Description)
}
