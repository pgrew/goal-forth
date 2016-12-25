package henceforth

import "time"

type Milestone struct {
	Name         string
	TimeCreated  time.Time
	TimeComplete time.Time
	Obstacles    []Obstacle
	Cause        CauseAffect
	Affect       CauseAffect
}

func NewMilestone(name string) Milestone {
	m := Milestone{
		Name:        name,
		TimeCreated: time.Now(),
	}
	return m
}
