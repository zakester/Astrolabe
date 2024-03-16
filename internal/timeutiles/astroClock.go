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

func (ac AstroClock) Init(t time.Time, timezone float32) *AstroClock {
	var jd = julian.Init(t, timezone)
	return &AstroClock{
		Time:    t,
		Julian:  jd,
		MeanGST: MeanGST(jd),
	}
}

func (ac AstroClock) Set(t time.Time, timezone float32) *AstroClock {
	ac = *ac.Init(t, timezone)
	return &ac
}
