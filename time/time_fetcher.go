package time

import "time"

type TimeFetcher interface {
	GetNow() time.Time
}

type LocalTimeFetcher struct{}

func (tf LocalTimeFetcher) GetNow() time.Time {
	return time.Now()
}
