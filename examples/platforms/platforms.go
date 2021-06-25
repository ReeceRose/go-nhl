package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Platforms")
	client := nhl.NewClient()
	fmt.Println("Attempting to get platforms")
	platforms, statusCode, err := client.GetPlatforms()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetPlatforms returned a status code of %d\n", statusCode)
	for _, platform := range platforms {
		fmt.Printf("%s:%s\n", platform.PlatformCode, platform.PlatformDescription)
	}

	fmt.Println("Attempting to get platform by code")
	platform, statusCode, err := client.GetPlatformByCode("web")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetPlatformByCode returned a status code of %d\n", statusCode)
	fmt.Printf("Got platform of %s\n", platform.PlatformCode)
}
