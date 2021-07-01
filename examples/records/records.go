package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Records")
	client := nhl.NewClient()
	fmt.Println("Attempting to get St. Louis Blues all-time records vs franchises")
	records, statusCode, err := client.GetAllTimeRecordVsFranchiseById(18)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetAllTimeRecordVsFranchiseById returned a status code of %d\n", statusCode)
	for _, record := range records.Data {
		fmt.Printf("Wins vs Losses against: %s. %dW %dL %dT\n", record.OpponentFranchiseName, record.TotalWins, record.TotalLosses, record.TotalTies)
	}
}
