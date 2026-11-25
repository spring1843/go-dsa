package dnc

import (
	"slices"
	"time"
)

var rateLimitEvents []int64

// nanoSecond represents one second in nanoseconds.
const nanoSecond = 1e9

// IsAllowed solves the problem in O(log n) time and O(n) space.
func IsAllowed(limitPerSecond int) bool {
	now := time.Now().UnixNano()
	removeOldRateLimitEvents(now)
	if len(rateLimitEvents) >= limitPerSecond {
		return false
	}

	rateLimitEvents = append(rateLimitEvents, now)
	return true
}

// removeOldRateLimitEvents uses binary search to remove events that are at least 1 second old.
func removeOldRateLimitEvents(now int64) {
	if len(rateLimitEvents) == 0 {
		return
	}

	cutoff := now - nanoSecond
	// An event that is exactly one second old no longer counts.
	idx, _ := slices.BinarySearch(rateLimitEvents, cutoff+1)

	rateLimitEvents = rateLimitEvents[idx:]
}
