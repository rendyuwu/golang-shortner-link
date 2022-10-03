package helper

import "time"

// get currect timestamp and add duration, "ms", "s", "m", "h"
// example 2h45m
func AddTimestamp(s string) int64 {
	duration, err := time.ParseDuration(s)
	PanicIfError(err)
	now := time.Now().Add(duration)

	return now.Unix()
}
