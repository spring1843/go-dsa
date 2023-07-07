package dnc

import (
	"time"
)

var rateLimitEvents []int64

// IsAllowed solves the problem in O(1) time and O(n) space.
func IsAllowed(limitPerSecond int) bool {
	now := time.Now().Unix()
	removeOldERateLimitEvents(now)
	if len(rateLimitEvents) >= limitPerSecond {
		return false
	}

	rateLimitEvents = append(rateLimitEvents, now)
	return true
}

func removeOldERateLimitEvents(now int64) {
	if len(rateLimitEvents) == 0 || now <= rateLimitEvents[0] {
		return
	}

	rateLimitEvents = []int64{}
}
