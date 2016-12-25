package henceforth

type Causal interface {
	GetCause() []CauseAffect
	GetAffect() []CauseAffect
}

type CauseAffect struct {
	Name        string
	Description string
}
