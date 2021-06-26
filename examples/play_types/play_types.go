package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Play Types")
	client := nhl.NewClient()
	fmt.Println("Attempting to get play types")
	playTypes, statusCode, err := client.GetPlayTypes()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetPlayTypes returned a status code of %d\n", statusCode)
	for _, playType := range playTypes {
		fmt.Printf("%s:%s:%s\n", playType.ID, playType.Name, playType.CmsKey)
	}

	fmt.Println("Attempting to get game type by ID")
	playType, statusCode, err := client.GetPlayTypeById("SHOT")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetPlayTypeById returned a status code of %d\n", statusCode)
	fmt.Printf("Got play type of %s\n", playType.Name)
}
