package models

import (
	"fmt"
	"time"
)

type User struct {
	Name             string    `json:"name"`
	StartWorkingTime time.Time `json:"start_working_time"`
	EndWorkingTime   time.Time `json:"end_working_time"`
	TeamName         string
}

func (u User) Validate() error {
	if u.Name == "" || u.StartWorkingTime.IsZero() || u.EndWorkingTime.IsZero() {
		return fmt.Errorf("inadequate input parameters. Required name, email, start_working_time, end_working_time")
	}
	return nil
}

type Event struct {
	Name                    string    `json:"name"`
	Date                    time.Time `json:"date"`
	StartTime               time.Time `json:"start_time"`
	EndTime                 time.Time `json:"end_time"`
	Users                   []string  `json:"users,omitempty"`
	Teams                   []string  `json:"teams,omitempty"`
	NumberOfRepresentatives int       `json:"number_of_representatives"`
}

func (e Event) Validate() error {
	if len(e.Users) == 0 && len(e.Teams) == 0 {
		return fmt.Errorf("required participants for the event")
	}
	if e.Name == "" || e.Date.IsZero() || e.StartTime.IsZero() || e.EndTime.IsZero() {
		return fmt.Errorf("inadequate input parameters. Required name, date, start_time, end_time")
	}
	if e.StartTime.Sub(e.EndTime).Seconds() < 0 {
		return fmt.Errorf("invalid start time and end times")
	}
	return nil
}
