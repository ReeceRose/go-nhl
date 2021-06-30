package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Milestones")
	client := nhl.NewClient()
	fmt.Println("Attempting to get players who have hit the 1000 point career milestone")
	milestone, statusCode, err := client.Get1000PointCareerMilestone()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetTrophies returned a status code of %d\n", statusCode)
	fmt.Printf("A total of %d players have hit the 1000 point career milestone\n", milestone.Total)
}
