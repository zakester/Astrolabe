package celestialobject

import "github.com/zakester/Astrolabe/internal/mathutils"

type CelestialCoordinates struct {
	// Declination.
	delta float64
	// Right Ascension.
	alpha float64
}

// Convert Declination (δ) and Right Ascension (α) to Ecliptic Longitude (λ) and Ecliptic Latitude (β)
func (ec CelestialCoordinates) ToEclipticCoordintaes() *EclipticCoordinates {
  var A = (mathutils.Cos(Epsilon)*mathutils.Sin(ec.delta)*mathutils.Sin(ec.alpha) + mathutils.Sin(Epsilon) * mathutils.Sin(ec.delta)) / mathutils.Cos(ec.delta) * mathutils.Cos(ec.alpha)

  var B = mathutils.Cos(Epsilon) * mathutils.Sin(ec.delta) - mathutils.Sin(Epsilon) * mathutils.Cos(ec.delta) * mathutils.Sin(ec.alpha)
	return &EclipticCoordinates{
		lambda: mathutils.Asin(A),
		beta:   mathutils.Atan(B),
	}
}
