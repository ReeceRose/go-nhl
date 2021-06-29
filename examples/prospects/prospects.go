package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Prospects")
	client := nhl.NewClient()
	fmt.Println("Attempting to get prospects")
	prospects, statusCode, err := client.GetProspects()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetProspects returned a status code of %d\n", statusCode)
	fmt.Printf("First prospect returned is %s\n", prospects[0].FullName)
}
