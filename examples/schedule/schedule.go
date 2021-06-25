package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Schedule")
	client := nhl.NewClient()
	fmt.Println("Attempting to get schedule")
	schedule, statusCode, err := client.GetSchedule()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetSchedule returned a status code of %d\n", statusCode)
	fmt.Printf("Tonight there are %d game(s) on\n", schedule.TotalGames)
}
