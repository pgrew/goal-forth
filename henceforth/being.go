package henceforth

type Being struct {
	Goals      []Goal
	Ignorances []Ignorance
	Sufferings []Suffering
}

func (b Being) IsEnlightened() bool {
	if l := len(b.Ignorances); l == 0 {
		return true
	}
	return false
}
