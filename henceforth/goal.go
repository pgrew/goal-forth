package henceforth

import "time"

type Goal struct {
	TimeBegan  time.Time
	TimeTarget time.Time
	Milestones []Milestone
	Motivation Motivation
	Tasks      []Task
	Complete   bool
}

func (g *Goal) AddMilestone(m Milestone) {
	g.Milestones = append(g.Milestones, m)
}

func (g Goal) IsComplete() bool {
	return g.Complete
}

func (g Goal) GetProbabilityOfSuccessfulCompletion() float32 {
	return 0.0
}
