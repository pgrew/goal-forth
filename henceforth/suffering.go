package henceforth

import "time"

type Suffering struct {
	Name        string
	TimeCreated time.Time
	TimeEnded   time.Time
	Cause       CauseAffect
	Affect      CauseAffect
}
