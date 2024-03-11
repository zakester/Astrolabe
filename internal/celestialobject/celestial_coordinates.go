package celestialobject

import "github.com/zakester/Astrolabe/internal/mathutils"

type CelestialCoordinates struct {
	// Declination.
	Delta float64
	// Right Ascension.
	Alpha float64
}

// Convert Declination (δ) and Right Ascension (α) to Ecliptic Longitude (λ) and Ecliptic Latitude (β)
func (cc CelestialCoordinates) ToEclipticCoordintaes() *EclipticCoordinates {
  var A = (mathutils.Cos(Epsilon)*mathutils.Sin(cc.Delta)*mathutils.Sin(cc.Alpha) + mathutils.Sin(Epsilon) * mathutils.Sin(cc.Delta)) / mathutils.Cos(cc.Delta) * mathutils.Cos(cc.Alpha)

  var B = mathutils.Cos(Epsilon) * mathutils.Sin(cc.Delta) - mathutils.Sin(Epsilon) * mathutils.Cos(cc.Delta) * mathutils.Sin(cc.Alpha)
	return &EclipticCoordinates{
		Lambda: mathutils.Asin(A),
		Beta:   mathutils.Atan(B),
	}
}
