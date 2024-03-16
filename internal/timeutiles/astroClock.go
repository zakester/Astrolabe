package timeutiles

import (
	"time"

	"github.com/zakester/Astrolabe/internal/julian"
)

type AstroClock struct {
	Time    time.Time
	Julian  *julian.Julian
	MeanGST float64
}

func (ac AstroClock) Init(t time.Time, utcOffset float32) *AstroClock {
	var jd = julian.Init(t, utcOffset)
	return &AstroClock{
		Time:    t,
		Julian:  jd,
		MeanGST: MeanGST(jd),
	}
}

func (ac AstroClock) Set(t time.Time, utcOffset float32) *AstroClock {
	ac = *ac.Init(t, utcOffset)
	return &ac
}
