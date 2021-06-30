package main

import (
	"fmt"

	nhl "github.com/reecerose/go-nhl"
)

func main() {
	fmt.Println("Go NHL Wrapper Example: Attendance")
	client := nhl.NewClient()
	fmt.Println("Attempting to get attendance")
	attendance, statusCode, err := client.GetAttendance()
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetAttendance returned a status code of %d\n", statusCode)
	for _, year := range attendance.Data {
		fmt.Printf("Year: %d Total Attendance: %d\n", year.SeasonID, year.TotalAttendance)
	}
}
