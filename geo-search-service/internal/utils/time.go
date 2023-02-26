package utils

import "time"

func NowUTC() time.Time {
	return time.Now().UTC()
}

func TimeSubMinutes(t1, t2 time.Time) float64 {
	diff := t1.Sub(t2)
	return diff.Minutes()
}
