package celestialobject

import (
	"math"

	"github.com/zakester/Astrolabe/internal/mathutils"
)

type Sun struct {
	MeanAnomaly         float64
	MeanLongitude       float64
	EclipticCoordinates *EclipticCoordinates
}

func InitSun(time float64) *Sun {
	var meanAnomaly = math.Mod((357.52543 + 35999.04944*time - 0.58*time*time), 360.0)
	var meanLongitude = math.Mod((280.46646 + 36000.76983*time - 0.0003032*time*time), 360.0)
	// Sun's equation of center
	var equationOfCenter = +(1.914602 - (0.004817 * time) - (0.000014 * math.Pow(time, 2))) * mathutils.Sin(meanAnomaly)+(0.019993 - (0.000101 * time)) * mathutils.Sin(meanAnomaly*2)+(0.000289 * mathutils.Sin(meanAnomaly*3))

	return &Sun{
		MeanAnomaly:   meanAnomaly,
		MeanLongitude: meanLongitude,
		EclipticCoordinates: &EclipticCoordinates{
			Lambda: meanLongitude + equationOfCenter,
			Beta:   0.0,
		},
	}
}

func earthEccentricity(time float64) float64 {
	var earthEccentricity = 0.016708634 - 0.000042037*time - 0.0000001267*math.Pow(time, 2)
	return earthEccentricity
}

//func (s Sun) MeanAnomaly() float64 {
//	return math.Mod((357.52543 + 35999.04944*s.Time - 0.58*s.Time*s.Time), 360.0)
//}

//func (s Sun) MeanLongitude() float64 {
//	return math.Mod((516.16644 + 926534.9922*s.Time - 9.21*s.Time*s.Time), 360.0)
//}
