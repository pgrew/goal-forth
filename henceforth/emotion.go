package henceforth

type Emotion struct {
	Name       string
	Cause      CauseAffect
	Affect     CauseAffect
	Sufferings []Suffering
}

func (e *Emotion) AddSuffering(s Suffering) {
	e.Sufferings = append(e.Sufferings, s)
}
