package util

import "time"

// GetDayStartTime returns the current day start time in UTC for a given time
func GetDayStartTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, time.UTC)
}

// GetWeekStartTime returns the current week start time in UTC for a given time
func GetWeekStartTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day()+int(time.Monday-d.Weekday()), 0, 0, 0, 0, time.UTC)
}

// SameDay checks if to dates are equal, omitting time part
func SameDay(d1 time.Time, d2 time.Time) bool {
	return GetDayStartTime(d1) == GetDayStartTime(d2)
}

// SameWeek checks if to dates are in same week, omitting time part
func SameWeek(d1 time.Time, d2 time.Time) bool {
	return GetWeekStartTime(d1) == GetWeekStartTime(d2)
}
