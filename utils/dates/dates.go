package dates

import "time"

// dateLayout is a string representing the layout used for formatting date and time.
// It follows the Go time package's layout format, which is "2006-01-02T15:04:05Z" for ISO 8601 format.
const dateLayout = "2006-01-02T15:04:05Z"

// GetTime returns the current time in UTC.
func GetTime() time.Time {
	return time.Now().UTC()
}

// GetFormattedTime returns the current time in UTC formatted as a string according to the dateLayout constant.
func GetFormattedTime() string {
	return GetTime().Format(dateLayout)
}
