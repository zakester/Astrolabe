package mathutils

import "math"

const RADIANS float64 = math.Pi / 180

func Sin(x float64) float64 {
	return math.Sin(x * RADIANS)
}

func Cos(x float64) float64 {
  return math.Cos(x * RADIANS)
}

func Tan(x float64) float64 {
  return math.Tan(x * RADIANS)
}

func Asin(x float64) float64 {
	return math.Asin(x) * 180 / math.Pi
}

func Atan(x float64) float64 {
  return math.Atan(x) * 180 / math.Pi
}

func Atan2(y float64, x float64) float64 {
  return math.Atan2(y, x) * 180 / math.Pi
}
