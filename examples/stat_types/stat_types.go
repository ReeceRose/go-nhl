package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Stat Types")
	client := nhl.NewClient()
	fmt.Println("Attempting to get stat types")
	statTypes, statusCode, err := client.GetStatTypes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetStatTypes returned a status code of %d\n", statusCode)
	for _, statType := range statTypes {
		fmt.Printf("%s\n", statType.DisplayName)
	}
}
