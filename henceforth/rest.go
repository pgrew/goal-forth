package henceforth

import "time"

type Quality string

const (
	HighQuality Quality = "High"
	MidQuality          = "Mid"
	LowQuality          = "Low"
)

// Rest should implment Completable
type Rest struct {
	Duration time.Duration
}
