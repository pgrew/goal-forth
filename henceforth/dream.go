package henceforth

import "time"

type Dream struct {
	TimeBegan   time.Time
	TimeEnd     time.Time
	Description string
	Cause       CauseAffect
	Affect      CauseAffect
}
