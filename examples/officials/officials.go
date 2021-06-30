package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Officials")
	client := nhl.NewClient()
	fmt.Println("Attempting to get officials")
	officials, statusCode, err := client.GetOfficials()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetOfficials returned a status code of %d\n", statusCode)
	fmt.Printf("There's been a total of %d officials\n", officials.Total)
}
