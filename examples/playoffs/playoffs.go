package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Playoff Series")
	client := nhl.NewClient()
	fmt.Println("Attempting to get 2018/2019 Stanley Cup playoff games")
	series, statusCode, err := client.GetPlayoffSeriesBySeriesAndYear("Stanley Cup Final", "20182019")
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetPlayoffSeriesBySeriesAndYear returned a status code of %d\n", statusCode)
	fmt.Printf("2018/2019 Stanley Cup playoff series had %d games\n", series.Total)
}
