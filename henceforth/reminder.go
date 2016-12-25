package henceforth

import "time"

type Reminder struct {
	Name         string
	Message      string
	TimeCreated  time.Time
	TimeToRemind time.Time
}
