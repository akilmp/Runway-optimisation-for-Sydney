package perf

import "time"

func AdjustOAT(oat float64, t time.Time) float64 {
	month := t.Month()
	hour := t.Hour()
	if (month == time.December || month == time.January || month == time.February) &&
		hour >= 11 && hour <= 16 {
		return oat + 7
	}
	return oat
}
