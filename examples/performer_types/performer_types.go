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
}
