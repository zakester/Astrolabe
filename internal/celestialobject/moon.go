package celestialobject

import (
  "math"
	"github.com/zakester/Astrolabe/internal/mathutils"
)

const RADIANS float64 = math.Pi / 180

type Moon struct {
	Time           float64
	SunMeanAnomaly float64
}

func (m Moon) MeanLongitude() float64 {
	return math.Mod(218.31617+481267.88088*m.Time-4.06*m.Time*m.Time, 360.0)
}

func (m Moon) LongitudeDistanceFromSun() float64 {
	return math.Mod(297.85027+445267.11135*m.Time-5.15*m.Time*m.Time, 360.0)
}

func (m Moon) ArgumentOfLatitude() float64 {
	return math.Mod(93.27283+483202.01873*m.Time-11.56*m.Time*m.Time, 360.0)
}

func (m Moon) MeanAnomaly() float64 {
	return math.Mod(134.96298139+(477198.867398*m.Time)+(0.0086972*m.Time*m.Time), 360.0)
}

func (m Moon) MajorInequality() float64 {
	var A = m.MeanAnomaly() * RADIANS
	var B = m.MeanAnomaly() * 2 * RADIANS
	return 6.288889*math.Sin(A) + 0.213611*math.Sin(B)
	//return math.Mod(6.288889*math.Sin(A)+0.213611*math.Sin(B), 360.0)
}

func (m Moon) Evecation() float64 {
	var A = (m.MeanAnomaly() - 2*m.LongitudeDistanceFromSun()) * RADIANS
	return -1.273889 * math.Sin(A)
	//return math.Mod(-1.273889 * math.Sin(A), 360)
}

func (m Moon) Variation() float64 {
	var A = 2 * m.LongitudeDistanceFromSun() * math.Pi / 180
	return 0.658333 * math.Sin(A)
	//return math.Mod(0.658333*math.Sin(A), 360)
}

func (m Moon) AnnualInequality() float64 {
	var A = m.SunMeanAnomaly * math.Pi / 180
	return -0.185556 * math.Sin(A)
	//return math.Mod(-0.185556*math.Sin(A), 360)
}

func (m Moon) EclipticReduction() float64 {
	var A = 2 * m.ArgumentOfLatitude() * RADIANS
	return -0.114444 * math.Sin(A)
	//return math.Mod(-0.114444*math.Sin(A), 360)
}

func (m Moon) ParallacticInequality() float64 {
	var A = m.LongitudeDistanceFromSun() * RADIANS
	return -0.034722 * math.Sin(A)
	//return math.Mod(-0.034722*math.Sin(A), 360)
}

func (m Moon) Term1() float64 {
	var A = (2*m.MeanAnomaly() - 2*m.LongitudeDistanceFromSun()) * RADIANS
	return -0.058889 * math.Sin(A)
}

func (m Moon) Term2() float64 {
	var A = (m.MeanAnomaly() + m.SunMeanAnomaly - 2*m.LongitudeDistanceFromSun()) * RADIANS
	return -0.057222 * math.Sin(A)
}

func (m Moon) Term3() float64 {
	var A = (m.MeanAnomaly() + 2*m.LongitudeDistanceFromSun()) * RADIANS
	return 0.053333 * math.Sin(A)
}

func (m Moon) Term4() float64 {
	var A = (m.SunMeanAnomaly - 2*m.LongitudeDistanceFromSun()) * RADIANS
	return -0.045833 * math.Sin(A)
}

func (m Moon) Term5() float64 {
	var A = (m.MeanLongitude() - m.SunMeanAnomaly) * RADIANS
	return 0.041111 * math.Sin(A)
}

func (m Moon) Term6() float64 {
	var A = (m.MeanAnomaly() + m.SunMeanAnomaly) * RADIANS
	return -0.030556 * math.Sin(A)
}

func (m Moon) Term7() float64 {
	var A = (2*m.ArgumentOfLatitude() - 2*m.LongitudeDistanceFromSun()) * RADIANS
	return -0.015278 * math.Sin(A)
}

// The sum of all Terms (Term1-Term7)
func (m Moon) Terms() float64 {
	return m.Term1() + m.Term2() + m.Term3() + m.Term4() + m.Term5() + m.Term6() + m.Term7()
}

// Sum of all Moon Anomalies (Q1-Q6)
func (m Moon) Qs() float64 {
	return m.MajorInequality() + m.Evecation() + m.Variation() + m.AnnualInequality() + m.EclipticReduction() + m.ParallacticInequality()
}

func (m Moon) EclipticCoordinates() *EclipticCoordinates {
	var F = m.ArgumentOfLatitude() + m.Qs()

	return &EclipticCoordinates{
		Lambda: m.MeanLongitude() + m.Qs(),
		Beta:   mathutils.Asin(mathutils.Sin(MoonInclination) * mathutils.Sin(F)),
	}
}
