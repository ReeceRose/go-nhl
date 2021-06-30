package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Trophies")
	client := nhl.NewClient()
	fmt.Println("Attempting to get trophies")
	trophies, statusCode, err := client.GetTrophies()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetTrophies returned a status code of %d\n", statusCode)
	fmt.Printf("Got a total of %d trophies\n", len(trophies.Data))
}
