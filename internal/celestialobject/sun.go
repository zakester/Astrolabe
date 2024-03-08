package celestialobject

import "math"

type Sun struct {
	t float64
}

func (s Sun) MeanAnomaly() float64 {
	return math.Mod((357.52543 + 35999.04944*s.t - 0.58*s.t*s.t), 360.0)
}
