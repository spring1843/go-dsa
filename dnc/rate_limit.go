package dnc

import (
	"sort"
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

// removeOldRateLimitEvents uses binary search to remove events older than 1 second.
func removeOldRateLimitEvents(now int64) {
	if len(rateLimitEvents) == 0 {
		return
	}

	cutoff := now - nanoSecond
	idx := sort.Search(len(rateLimitEvents), func(i int) bool {
		return rateLimitEvents[i] >= cutoff
	})

	rateLimitEvents = rateLimitEvents[idx:]
}
