package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Languages")
	client := nhl.NewClient()
	fmt.Println("Attempting to get languages")
	languages, statusCode, err := client.GetLanguages()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetLanguages returned a status code of %d\n", statusCode)
	for _, language := range languages {
		fmt.Printf("%s:%s\n", language.LanguageCode, language.Locale)
	}
}
