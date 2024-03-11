package main

import (
	"fmt"
	"time"

	celestialobject "github.com/zakester/Astrolabe/internal/celestialobject"
	julian "github.com/zakester/Astrolabe/internal/julian"
)

func main() {
	//var date = time.Date(2005, time.May, 5, 0, 0, 0, 0, time.Local)
	var jd = julian.Init(time.Now())

	var sun = celestialobject.Sun{Time: jd.Century()}

	var moon = celestialobject.Moon{
		Time:           jd.Century(),
		SunMeanAnomaly: sun.MeanAnomaly(),
	}

	fmt.Printf("Julian Day: %f\n", jd.JulianDay)
	fmt.Printf("Since J2000: %f\n", jd.SinceJ2000())
	fmt.Printf("Century: %f\n", jd.Century())

	fmt.Println("----Sun----")
	fmt.Printf("Sun's Mean Anomaly: %f\n", sun.MeanAnomaly())
	//fmt.Printf("Mean Longitude L0: %f\n", sun.MeanLongitude())

	fmt.Println("")

	fmt.Println("----Moon----")
	fmt.Printf("Mean Longitude L0:            %f\n", moon.MeanLongitude())
	fmt.Printf("Moon's Mean Anomaly:          %f\n", moon.MeanAnomaly())
	fmt.Printf("Argument Of Latitude:         %f\n", moon.ArgumentOfLatitude())
	fmt.Printf("Major Inequality (q1):        %f\n", moon.MajorInequality())
	fmt.Printf("Evecation (q2):               %f\n", moon.Evecation())
	fmt.Printf("Variation (q3):               %f\n", moon.Variation())
	fmt.Printf("Annual Inequality (q4):       %f\n", moon.AnnualInequality())
	fmt.Printf("Ecliptic Reduction (q5):      %f\n", moon.EclipticReduction())
	fmt.Printf("Parallactic Inequality (q6):  %f\n", moon.ParallacticInequality())

	//var D = sun.MeanLongitude()
	//fmt.Printf("D: %f\n", moon.MeanLongitude() - D)
	fmt.Printf("D: %f\n", moon.LongitudeDistanceFromSun())

	var ec = moon.EclipticCoordinates()
	var cc = ec.ToCelestialCoordinates()

	fmt.Printf("Moon's Longitude: %f\n", ec.Lambda)
	fmt.Printf("Moon's Latitude: %f\n", ec.Beta)

  fmt.Println("")

	fmt.Printf("Moon's Declination: %f\n", cc.Delta)
	fmt.Printf("Moon's Right Ascension: %f\n", cc.Alpha)
}
