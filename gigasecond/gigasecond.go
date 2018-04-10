/*
Package calculates a time 1 billion seconds into the future.

Input is a time.Time object. AddGigasecond will output
a time object 1 billion seconds into the future.
*/
package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	// Calculate a time 1E9 seconds, 1 billion seconds, into the future.
	return t.Add(1E9 * time.Second)
}
