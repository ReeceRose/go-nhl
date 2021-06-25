package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Configurations")
	client := nhl.NewClient()
	fmt.Println("Attempting to get configurations")
	configurations, statusCode, err := client.GetConfigurations()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetConfigurations returned a status code of %d\n", statusCode)
	for _, configuration := range configurations {
		fmt.Printf("%s:%s\n", configuration.Description, configuration.Endpoint)
	}
}
