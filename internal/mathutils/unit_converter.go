package mathutils

import (
	"fmt"
	"math"
)

type DMS struct {
	Deg int
	Min int
	Sec int
	DMS string
}

type HMS struct {
	Hour int
	Min  int
	Sec  int
	HMS  string
}

func DegreesToHMS(x float64) *HMS {
	var totalSeconds = int(x * 3600)
	var hour = totalSeconds / 3600
	var min = (totalSeconds % 3600) / 60
	var sec = totalSeconds % 60

	return &HMS{
		Hour: hour,
		Min:  min,
		Sec:  sec,
		HMS:  fmt.Sprintf("%02d:%02d:%02d", hour, min, sec),
	}
}

func HMStoDegrees(hms HMS) float64 {
	var totalSeconds = hms.Hour*3600 + hms.Min*60 + hms.Sec
	var degrees = float64(totalSeconds) / 3600.0
	return degrees
}

func Deg2DMS(x float64) *DMS {
	var deg = int(x)
	var min = int((x - float64(deg)) * 60)
	var sec = int(x-float64(deg)-float64(min/60)) * 3600

	return &DMS{
		Deg: deg,
		Min: min,
		Sec: sec,
		DMS: fmt.Sprintf("%d° %d' %d\"", deg, min, sec),
	}
}

func DMS2Deg(x DMS) float64 {
	return float64(x.Deg) + float64(x.Min/60) + float64(x.Sec/3600)
}

func HourAngle2Deg(x float64) float64 {
	return x / 0.066666667
}

func Deg2HourAngle(x float64) float64 {
	return x * 0.066666667
}

func Deg2Rad(x float64) float64 {
	return x * math.Pi / 180
}

func Rad2Deg(x float64) float64 {
	return x * 180 / math.Pi
}
