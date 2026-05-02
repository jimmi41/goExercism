// Package gigasecond provides functionality to calculate the moment
// that occurs one gigasecond (1,000,000,000 seconds) after a given time.
package gigasecond

import "time"

// AddGigasecond returns the time that occurs exactly one gigasecond
// after the given time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1_000_000_000)
}