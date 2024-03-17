package celestialobject

import (
	"time"

	"github.com/zakester/Astrolabe/internal/julian"
	"github.com/zakester/Astrolabe/internal/timeutiles"
)

type Observer struct {
	Latitude  float64
	Longitude float64
	Time      time.Time
	Julian    *julian.Julian
	MeanGST   float64
}

func InitObserver(latitude float64, longitude float64, t time.Time, utcOffset float32) *Observer {
	var jd = julian.Init(t, utcOffset)
	var meanGST = timeutiles.MeanGST(jd)

	return &Observer{
		Latitude:  latitude,
		Longitude: longitude,
		Time:      t,
		Julian:    jd,
		MeanGST:   meanGST,
	}
}
