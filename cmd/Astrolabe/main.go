package main

import (
	"time"
  "fmt"
	celestialobject "github.com/zakester/Astrolabe/internal/celestialobject"
	"github.com/zakester/Astrolabe/internal/mathutils"
)


func main() {
	//var date = time.Date(2024, time.March, 16, 12, 57, 00, 0, time.Local)
	var date = time.Now()
	var observer = celestialobject.InitObserver(36.5166667, 2.88333, date, 1.0)
	var jd = observer.Julian

	var sun = celestialobject.InitSun(observer.Julian.Century())
	var moon = celestialobject.InitMoon(observer.Julian.Century(), sun.MeanAnomaly)

	fmt.Printf("Julian Day:  %f\n", jd.JulianDay)
	fmt.Printf("Since J2000: %f\n", jd.SinceJ2000())
	fmt.Printf("Century:     %f\n", jd.Century())

	fmt.Println("----Sun----")
	fmt.Printf("Sun's Mean Anomaly: %f\n", sun.MeanAnomaly)
	fmt.Printf("Mean Longitude L0:  %f\n", sun.MeanLongitude)
	fmt.Printf("Longitude:          %f\n", sun.EclipticCoordinates.Lambda)

	fmt.Println("")

	var sunCC = sun.EclipticCoordinates.ToCelestialCoordinates()
	fmt.Printf("Sun's Declination:     %f°     %f ha      %s\n", sunCC.Delta, mathutils.Deg2HourAngle(sunCC.Delta), mathutils.Deg2HMS(sunCC.Delta).ToString())
	fmt.Printf("Sun's Right Ascension: %f°     %f ha      %s\n", sunCC.Alpha, mathutils.Deg2HourAngle(sunCC.Alpha), mathutils.Deg360HMS(sunCC.Alpha).ToString())

	fmt.Println("")

	fmt.Printf("Sun's Altitude: %f°\n", sun.EclipticCoordinates.ToCelestialCoordinates().ToHorizonzalCoordinates(observer).Altitude)
	fmt.Printf("Sun's Azimuth:  %f°\n", sun.EclipticCoordinates.ToCelestialCoordinates().ToHorizonzalCoordinates(observer).Azimuth)

	fmt.Println("")

	fmt.Println("----Moon----")
	fmt.Printf("Mean Longitude L0:            %f\n", moon.MeanLongitude)
	fmt.Printf("Moon's Mean Anomaly:          %f\n", moon.MeanAnomaly)
	fmt.Printf("Argument Of Latitude:         %f\n", moon.ArgumentOfLatitude)
	fmt.Printf("Major Inequality (q1):        %f\n", moon.LunarAnomaly.MajorInequality)
	fmt.Printf("Evecation (q2):               %f\n", moon.LunarAnomaly.Evecation)
	fmt.Printf("Variation (q3):               %f\n", moon.LunarAnomaly.Variation)
	fmt.Printf("Annual Inequality (q4):       %f\n", moon.LunarAnomaly.AnnualInequality)
	fmt.Printf("Ecliptic Reduction (q5):      %f\n", moon.LunarAnomaly.EclipticReduction)
	fmt.Printf("Parallactic Inequality (q6):  %f\n", moon.LunarAnomaly.ParallacticInequality)

	//kar D = sun.MeanLongitude()
	//fmt.Printf("D: %f\n", moon.MeanLongitude() - D)
	fmt.Printf("D: %f\n", moon.LongitudeDistanceFromSun)

	var ec = moon.EclipticCoordinates()
	var cc = ec.ToCelestialCoordinates()
	var hc = cc.ToHorizonzalCoordinates(observer)

	fmt.Printf("Moon's Longitude: %f\n", ec.Lambda)
	fmt.Printf("Moon's Latitude: %f\n", ec.Beta)

	fmt.Println("")

	fmt.Printf("Moon's Declination: %f°\n", cc.Delta)
	fmt.Printf("Moon's Right Ascension: %f°     %f ha\n", cc.Alpha, mathutils.Deg2HourAngle(cc.Alpha))

	fmt.Println("")

	fmt.Printf("Moon's Altitude: %f°\n", hc.Altitude)
	fmt.Printf("Moon's Azimuth:  %f°\n", hc.Azimuth)
}
