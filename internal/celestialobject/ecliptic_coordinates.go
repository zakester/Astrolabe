package celestialobject

import "github.com/zakester/Astrolabe/internal/mathutils"

type EclipticCoordinates struct {
	// Ecliptic Latitude
	lambda float64
	// Ecliptic Longitude
	beta float64
}

// Ecliptic Longitude (λ) and Ecliptic Latitude (β) to Convert Declination (δ) and Right Ascension (α).
func (ec EclipticCoordinates) ToCelestialCoordinates() *CelestialCoordinates {
	var A = mathutils.Cos(Epsilon)*mathutils.Sin(ec.beta) + mathutils.Sin(Epsilon)*mathutils.Cos(ec.beta)*mathutils.Sin(ec.lambda)

	var B = (mathutils.Cos(Epsilon)*mathutils.Cos(ec.beta)*mathutils.Sin(ec.lambda) - mathutils.Sin(Epsilon)*mathutils.Sin(ec.beta)) / mathutils.Cos(ec.beta) * mathutils.Cos(ec.lambda)

	return &CelestialCoordinates{
		delta: mathutils.Asin(A),
		alpha: mathutils.Atan(B),
	}
}
