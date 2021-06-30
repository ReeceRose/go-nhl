package gonhl

import (
	"encoding/json"
)

func (c *Client) GetAttendance() (Attendance, int, error) {
	var attendance Attendance
	data, statusCode, err := c.get(RECORDS_BASE_URL + "/attendance/")
	if err != nil {
		return attendance, statusCode, err
	}
	json.Unmarshal(data, &attendance)
	return attendance, statusCode, nil
}

type Attendance struct {
	Data []struct {
		ID                int `json:"id"`
		PlayoffAttendance int `json:"playoffAttendance"`
		RegularAttendance int `json:"regularAttendance"`
		SeasonID          int `json:"seasonId"`
		TotalAttendance   int `json:"totalAttendance"`
	} `json:"data"`
	Total int `json:"total"`
}
