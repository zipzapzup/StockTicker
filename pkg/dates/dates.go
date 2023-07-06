package dates

import "time"

func GetTimeNow() string {
	currentTime := time.Now()
	return currentTime.Format(time.RFC1123)[:10]
}

func ListDates(NDAYS int) []string {
	// NDAYS gets the number of days in the past from today
	// Idea here is to create a reverse loop that will return all dates leading to NDAYS
	// RangeOfDates is the placeholder

	RangeOfDates := []string{}
	CounterDate := time.Now()

	for i := 0; i <= NDAYS; i++ {
		if i == 0 {
			// NDAYS 0 mean we need to get today's date
			RangeOfDates = append(RangeOfDates, CounterDate.Format(time.RFC3339)[:10])
		} else {
			CounterDate = CounterDate.AddDate(0, 0, -1)
			RangeOfDates = append(RangeOfDates, CounterDate.Format(time.RFC3339)[:10])
		}
	}
	return RangeOfDates
}
