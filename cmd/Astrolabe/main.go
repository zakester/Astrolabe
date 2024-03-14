package main

import (
	"fmt"
	"time"

	celestialobject "github.com/zakester/Astrolabe/internal/celestialobject"
	julian "github.com/zakester/Astrolabe/internal/julian"
	"github.com/zakester/Astrolabe/internal/mathutils"
	"github.com/zakester/Astrolabe/internal/timeutiles"
)

func main() {
	//var date = time.Date(2024, time.March, 14, 15, 36, 54, 0, time.Local)
  var date = time.Now()
	var jd = julian.Init(date)

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

	//kar D = sun.MeanLongitude()
	//fmt.Printf("D: %f\n", moon.MeanLongitude() - D)
	fmt.Printf("D: %f\n", moon.LongitudeDistanceFromSun())

	var ec = moon.EclipticCoordinates()
	var cc = ec.ToCelestialCoordinates()
	var hc = cc.ToHorizonzalCoordinates(timeutiles.MeanGST(jd))

	fmt.Printf("Moon's Longitude: %f\n", ec.Lambda)
	fmt.Printf("Moon's Latitude: %f\n", ec.Beta)

  fmt.Println("")

	fmt.Printf("Moon's Declination: %f°\n", cc.Delta)
	fmt.Printf("Moon's Right Ascension: %f°     %f ha\n", cc.Alpha, mathutils.Deg2HourAngle(cc.Alpha))

  fmt.Println("")

	fmt.Printf("Moon's Altitude: %f°\n", hc.Altitude)
	fmt.Printf("Moon's Azimuth:  %f°\n", hc.Azimuth)
}
