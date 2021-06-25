package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Performer Types")
	client := nhl.NewClient()
	fmt.Println("Attempting to get performer types")
	performerTypes, statusCode, err := client.GetPerformerTypes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetPerformerTypes returned a status code of %d\n", statusCode)
	for _, performerType := range performerTypes {
		fmt.Printf("%s\n", performerType.Name)
	}

	fmt.Println("Attempting to get performer type by name")
	performerType, statusCode, err := client.GetPerformerTypeByName("person")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetPlatformByCode returned a status code of %d\n", statusCode)
	fmt.Printf("Got performer type of %s\n", performerType.Name)
}
