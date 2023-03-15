package dnc

import (
	"time"
)

var rateLimitEvents []int64

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

	lastIndex := indexOfFirstOldEvent(0, len(rateLimitEvents)-1, now-1)
	if lastIndex > -1 {
		rateLimitEvents = rateLimitEvents[:lastIndex]
	}
}

func indexOfFirstOldEvent(start, end int, search int64) int {
	if rateLimitEvents[start] == search {
		return start
	}

	if start == end {
		return -1
	}

	mid := start + (end-start)/2
	if rateLimitEvents[mid] == search {
		for i := mid; i > start; i-- {
			if rateLimitEvents[i] != search {
				return i + 1
			}
		}
		return 0
	}

	if search < rateLimitEvents[mid] {
		return indexOfFirstOldEvent(start, mid, search)
	}

	return indexOfFirstOldEvent(mid+1, end, search)
}
