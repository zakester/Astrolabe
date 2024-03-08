package main

import (
	"fmt"
	"time"

	celestialobject "github.com/zakester/Astrolabe/internal/celestialobject"
	julian "github.com/zakester/Astrolabe/internal/julian"
)

func main() {
	var jd = julian.Init(time.Now())
  var moon = celestialobject.Moon {
    Time: jd.Century(),
  }

	fmt.Printf("Julian Day: %f\n", jd.JulianDay)
	fmt.Printf("Since J2000: %f\n", jd.SinceJ2000())
	fmt.Printf("Century: %f\n", jd.Century())

	fmt.Printf("Moon's Mean Anomaly: %f\n", moon.MeanAnomaly())
	fmt.Printf("F: %f\n", moon.F())
	fmt.Printf("Mean Longitude L0: %f\n", moon.MeanLongitude())
}
