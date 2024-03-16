package timeutiles

import (
	"math"

	"github.com/zakester/Astrolabe/internal/julian"
	"github.com/zakester/Astrolabe/internal/mathutils"
)

func MeanGST(jd *julian.Julian) float64 {
  var era = earthRotationAngle(jd)
  var t = jd.Century()
  var gmst = math.Mod(era + (0.014506 + 4612.15739966*t + 1.39667721*t*t + -0.00009344*t*t*t + 0.00001882*t*t*t*t)/60/60*math.Pi/180, math.Pi*2)
  
  if gmst < 0 {
    gmst += math.Pi*2
  }

  //var gmstHMS = mathutils.Deg2HMS(mathutils.Deg2HourAngle(mathutils.Rad2Deg(gmst)))
  //fmt.Printf("GMST: %s\n", gmstHMS.ToString())

	return mathutils.Rad2Deg(gmst)
}

func earthRotationAngle(jd *julian.Julian) float64 {
	var sinceJ2000 = jd.SinceJ2000()
	var frac = math.Mod(sinceJ2000, 1)

	var era = math.Mod(math.Pi * 2.0 * (0.7790572732640 + 0.00273781191135448*sinceJ2000 + frac), math.Pi * 2)
	if era < 0 {
		era += math.Pi * 2
	}

	return era
}
