package celestialobject

import (
	"github.com/zakester/Astrolabe/internal/mathutils"
)

type EclipticCoordinates struct {
	// Ecliptic Longitude
	Lambda float64
	// Ecliptic Latitude
	Beta float64
}

// Ecliptic Longitude (λ) and Ecliptic Latitude (β) to Convert Declination (δ) and Right Ascension (α).
func (ec EclipticCoordinates) ToCelestialCoordinates() *CelestialCoordinates {
	var A = mathutils.Cos(Epsilon)*mathutils.Sin(ec.Beta) + mathutils.Sin(Epsilon)*mathutils.Cos(ec.Beta)*mathutils.Sin(ec.Lambda)

	var B = (mathutils.Cos(Epsilon)*mathutils.Cos(ec.Beta)*mathutils.Sin(ec.Lambda) - mathutils.Sin(Epsilon)*mathutils.Sin(ec.Beta)) // mathutils.Cos(ec.Beta) * mathutils.Cos(ec.Lambda)

	var alpha = mathutils.Atan2(B, mathutils.Cos(ec.Lambda)*mathutils.Cos(ec.Beta))
	if alpha < 0 {
    alpha += 360.0
	}

	//var B = mathutils.Cos(C)*mathutils.Sin(ec.Lambda)*mathutils.Cos(Epsilon) - mathutils.Sin(Epsilon)*mathutils.Sin(C)
	//var B = mathutils.Cos(Epsilon) * mathutils.Sin(ec.Lambda) - mathutils.Sin(Epsilon) * mathutils.Tan(ec.Beta)
	return &CelestialCoordinates{
		Delta: mathutils.Asin(A),
		Alpha: alpha,
	}
}
