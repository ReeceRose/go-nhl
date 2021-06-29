package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Positions")
	client := nhl.NewClient()
	fmt.Println("Attempting to get positions")
	positions, statusCode, err := client.GetPositions()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetPositions returned a status code of %d\n", statusCode)
	fmt.Printf("Got a total of %d positions\n", len(positions))
}
