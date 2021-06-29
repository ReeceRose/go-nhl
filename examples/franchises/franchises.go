package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Franchises")
	client := nhl.NewClient()
	fmt.Println("Attempting to get franchises")
	franchises, statusCode, err := client.GetFranchises()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetFranchises returned a status code of %d\n", statusCode)
	fmt.Printf("Got a total of %d franchises\n", len(franchises))
}
