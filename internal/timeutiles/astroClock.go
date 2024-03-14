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

func (ac AstroClock) Init(t time.Time) *AstroClock {
	var jd = julian.Init(t)
	return &AstroClock{
		Time:    t,
		Julian:  jd,
		MeanGST: MeanGST(jd),
	}
}

func (ac AstroClock) Set(t time.Time) *AstroClock {
  ac = *ac.Init(t)
  return &ac
}
