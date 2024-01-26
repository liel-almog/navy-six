package main

type missleType int

const (
	missleTypeTorpedo missleType = iota
	missleTypeBallistic
	missleTypeCruise
)

func (m missleType) string() string {
	switch m {
	case missleTypeTorpedo:
		return "Torpedo"
	case missleTypeBallistic:
		return "Ballistic"
	case missleTypeCruise:
		return "Cruise"
	default:
		return "Unknown"
	}
}

func (m missleType) enumIndex() int {
	return int(m)
}

type missle struct {
	missleType missleType
	failed     bool
}

func newMissle(missleType missleType) *missle {
	return &missle{missleType: missleType, failed: false}
}

var missles []*missle

func initMissles() {
	missles = make([]*missle, 10)

	for i := 0; i < len(missles); i++ {
		missles[i] = newMissle(missleType(i % 3))
	}
}
