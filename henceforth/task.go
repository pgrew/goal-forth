package henceforth

import "time"

type Task struct {
	Name          string
	TimeCreated   time.Time
	TimeCompleted time.Time
}
