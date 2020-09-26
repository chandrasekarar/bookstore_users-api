package dateutils

import "time"

const (
	apiDateFormat = "2006-01-02T15:04:05Z"
	apiDBLayout   = "2006-01-02 15:04:05"
)

// GetNow UTC Time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString UTC Time in string
func GetNowString() string {
	return GetNow().Format(apiDateFormat)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}
