package celestialobject

import "math"

type Sun struct {
	MeanAnomaly float64
}

func InitSun(time float64) *Sun {
	return &Sun{
    MeanAnomaly: math.Mod((357.52543 + 35999.04944*time - 0.58*time*time), 360.0),
  }
}

//func (s Sun) MeanAnomaly() float64 {
//	return math.Mod((357.52543 + 35999.04944*s.Time - 0.58*s.Time*s.Time), 360.0)
//}

//func (s Sun) MeanLongitude() float64 {
//	return math.Mod((516.16644 + 926534.9922*s.Time - 9.21*s.Time*s.Time), 360.0)
//}
