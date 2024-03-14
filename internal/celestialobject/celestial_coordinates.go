package celestialobject

import (
	"github.com/zakester/Astrolabe/internal/mathutils"
)

type CelestialCoordinates struct {
	// Declination. In Degrees
	Delta float64
	// Right Ascension. In Degrees
	Alpha float64
}

// Convert Declination (δ) and Right Ascension (α) to Ecliptic Longitude (λ) and Ecliptic Latitude (β)
func (cc CelestialCoordinates) ToEclipticCoordinates() *EclipticCoordinates {
	var A = (mathutils.Cos(Epsilon)*mathutils.Sin(cc.Delta)*mathutils.Sin(cc.Alpha) + mathutils.Sin(Epsilon)*mathutils.Sin(cc.Delta)) / mathutils.Cos(cc.Delta) * mathutils.Cos(cc.Alpha)

	var B = mathutils.Cos(Epsilon)*mathutils.Sin(cc.Delta) - mathutils.Sin(Epsilon)*mathutils.Cos(cc.Delta)*mathutils.Sin(cc.Alpha)
	return &EclipticCoordinates{
		Lambda: mathutils.Asin(A),
		Beta:   mathutils.Atan(B),
	}
}

func (cc CelestialCoordinates) ToHorizonzalCoordinates(meanGST float64) *HorizontalCoordinates {
	var lat = 36.5166667
	var h = meanGST + 2.88333 - cc.Alpha
	var sinAltitude = mathutils.Sin(lat)*mathutils.Sin(cc.Delta) + mathutils.Cos(lat)*mathutils.Cos(cc.Delta)*mathutils.Cos(h)
	var x = -mathutils.Sin(lat)*mathutils.Cos(cc.Delta)*mathutils.Cos(h) + mathutils.Cos(lat)*mathutils.Sin(cc.Delta)
	var y = mathutils.Cos(cc.Delta) * mathutils.Sin(h)

	return &HorizontalCoordinates{
		Altitude: mathutils.Asin(sinAltitude),
		Azimuth:  -mathutils.Atan2(y, x),
	}
}
