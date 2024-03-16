package julian

import "time"

const J2000 float64 = 2451545.0

type Julian struct {
	JulianDay float64
}

func (j Julian) SinceJ2000() float64 {
	return j.JulianDay - J2000
}

func (j Julian) Century() float64 {
	return (j.JulianDay - J2000) / 36525
}

func Init(t time.Time) *Julian {
	var m = t.Month()
	var y = t.Year()

	if m <= 2 {
		m = m + 12
		y = y - 1
	}

	var A = uint8(y / 100)
	var B = int8(2 + (A / 4) - A)
	var timezone = 1.0

	var year = uint32(365.25 * float64(y))
	var month = uint64(30.6001 * float64(m+1))
	var time = float64(t.Hour()*3600.0 + t.Minute()*60.0 + t.Second())

	var jd = 1720994.5 + float64(year) + float64(month) + float64(B) + float64(t.Day()) + (time / 86400.0) - (timezone / 24.0)

	return &Julian{JulianDay: jd}
}
