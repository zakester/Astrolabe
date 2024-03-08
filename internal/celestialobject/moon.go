package celestialobject

import "math"

type Moon struct {
	Time float64
}

func (m Moon) MeanLongitude() float64 {
	return math.Mod(218.31617+481267.88088*m.Time-4.06*m.Time*m.Time, 360.0)
}

func (m Moon) F() float64 {
	return math.Mod(93.27283+483202.01873*m.Time-11.56*m.Time*m.Time, 360.0)
}

func (m Moon) MeanAnomaly() float64 {
	return math.Mod(134.96298139+(477198.867398*m.Time)+(0.0086972*m.Time*m.Time), 360.0)
}
