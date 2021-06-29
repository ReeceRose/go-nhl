package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Draft")
	client := nhl.NewClient()
	fmt.Println("Attempting to get 2020 draft")
	draft, statusCode, err := client.GetDraftByYear("2020")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetDraftByYear returned a status code of %d\n", statusCode)
	fmt.Printf("Got the %d draft\n", draft.DraftYear)
}
