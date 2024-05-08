package dates

import "time"

const dateLayout = "2006-01-02T15:04:05Z"

func GetTime() time.Time {
	return time.Now().UTC()
}

func GetFormattedTime() string {
	return GetTime().Format(dateLayout)
}
