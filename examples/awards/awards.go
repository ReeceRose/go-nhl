package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Awards")
	client := nhl.NewClient()
	fmt.Println("Attempting to get awards")
	awards, statusCode, err := client.GetAwards()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetAwards returned a status code of %d\n", statusCode)
	fmt.Printf("Got a total of %d awards\n", len(awards))
}
