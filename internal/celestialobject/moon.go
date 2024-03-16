package celestialobject

import (
	"github.com/zakester/Astrolabe/internal/mathutils"
	"math"
)

const RADIANS float64 = math.Pi / 180

type Moon struct {
	MeanAnomaly              float64
	MeanLongitude            float64
	ArgumentOfLatitude       float64
	LongitudeDistanceFromSun float64
	LunarAnomaly             *LunarAnomaly
	Terms                    *Terms
}

func InitMoon(time, sunMeanAnomaly float64) *Moon {
	var meanAnomaly = math.Mod(134.96298139+(477198.867398*time)+(0.0086972*time*time), 360.0)
	var meanLongitude = math.Mod(218.31617+481267.88088*time-4.06*time*time, 360.0)
	var argumentOfLatitude = math.Mod(93.27283+483202.01873*time-11.56*time*time, 360.0)
	var longitudeDistanceFromSun = math.Mod(297.85027+445267.11135*time-5.15*time*time, 360.0)

	var lunarAnomaly = initLunarAnomaly(meanAnomaly, longitudeDistanceFromSun, sunMeanAnomaly, argumentOfLatitude)
	var terms = initTerms(meanAnomaly, longitudeDistanceFromSun, sunMeanAnomaly, argumentOfLatitude, meanLongitude)

	return &Moon{
		MeanAnomaly:              meanAnomaly,
		MeanLongitude:            meanLongitude,
		ArgumentOfLatitude:       argumentOfLatitude,
		LongitudeDistanceFromSun: longitudeDistanceFromSun,
		LunarAnomaly:             lunarAnomaly,
		Terms:                    terms,
	}
}

func (m Moon) EclipticCoordinates() *EclipticCoordinates {
	var F = m.ArgumentOfLatitude + m.LunarAnomaly.Sum()

	return &EclipticCoordinates{
		Lambda: m.MeanLongitude + m.LunarAnomaly.Sum(),
		Beta:   mathutils.Asin(mathutils.Sin(MoonInclination) * mathutils.Sin(F)),
	}
}

type LunarAnomaly struct {
	MajorInequality       float64
	Evecation             float64
	Variation             float64
	AnnualInequality      float64
	EclipticReduction     float64
	ParallacticInequality float64
}

func initLunarAnomaly(meanAnomaly float64, longitudeDistanceFromSun float64, sunMeanAnomaly float64, argumentOfLatitude float64) *LunarAnomaly {

	var majorInequality = 6.288889*mathutils.Sin(meanAnomaly) + 0.213611*mathutils.Sin(meanAnomaly*2)
	var evecation = -1.273889 * mathutils.Sin(meanAnomaly-2*longitudeDistanceFromSun)
	var variation = 0.658333 * mathutils.Sin(2*longitudeDistanceFromSun*math.Pi/180)
	var annualInequality = -0.185556 * mathutils.Sin(sunMeanAnomaly*math.Pi/100)
	var eclipticReduction = -0.114444 * mathutils.Sin(2*argumentOfLatitude)
	var parallacticInequality = -0.034722 * mathutils.Sin(longitudeDistanceFromSun)

	return &LunarAnomaly{
		MajorInequality:       majorInequality,
		Evecation:             evecation,
		Variation:             variation,
		AnnualInequality:      annualInequality,
		EclipticReduction:     eclipticReduction,
		ParallacticInequality: parallacticInequality,
	}
}

// Sum of all Lunar Anomalies (Q1-Q6).
func (l LunarAnomaly) Sum() float64 {
	return l.ParallacticInequality + l.AnnualInequality + l.EclipticReduction + l.Variation + l.Evecation + l.MajorInequality
}

type Terms struct {
	Term1 float64
	Term2 float64
	Term3 float64
	Term4 float64
	Term5 float64
	Term6 float64
	Term7 float64
}

func initTerms(meanAnomaly float64, longitudeDistanceFromSun float64, sunMeanAnomaly float64, argumentOfLatitude float64, meanLongitude float64) *Terms {

	var t1 = -0.058889 * math.Sin(2*meanAnomaly-2*longitudeDistanceFromSun)
	var t2 = -0.057222 * math.Sin(meanAnomaly+sunMeanAnomaly-2*longitudeDistanceFromSun)
	var t3 = 0.053333 * math.Sin(meanAnomaly+2*longitudeDistanceFromSun)
	var t4 = -0.045833 * math.Sin(meanLongitude-2*longitudeDistanceFromSun)
	var t5 = 0.041111 * math.Sin(meanAnomaly-sunMeanAnomaly)
	var t6 = -0.030556 * math.Sin(meanAnomaly+sunMeanAnomaly)
	var t7 = -0.015278 * math.Sin(2*argumentOfLatitude-2*longitudeDistanceFromSun)

	return &Terms{
		Term1: t1,
		Term2: t2,
		Term3: t3,
		Term4: t4,
		Term5: t5,
		Term6: t6,
		Term7: t7,
	}
}

// Sum of all Terms (Term1-Term7).
func (t Terms) Sum() float64 {
	return t.Term1 + t.Term2 + t.Term3 + t.Term4 + t.Term5 + t.Term6 + t.Term7
}
