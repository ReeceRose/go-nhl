package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Divisions")
	client := nhl.NewClient()
	fmt.Println("Attempting to get divisions")
	divisions, statusCode, err := client.GetDivisions()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetDivisions returned a status code of %d\n", statusCode)
	fmt.Printf("Got a total of %d divisions\n", len(divisions))

	fmt.Println("Attempting to get division by ID")
	division, statusCode, err := client.GetDivisionById(27) // Honda West
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetDivisionById returned a status code of %d\n", statusCode)
	fmt.Printf("Got divison with the name: %v\n", division.Name)
}
