package greedy

import (
	"cmp"
	"slices"
)

// Event represents an event to be scheduled.
type Event struct {
	Name      string
	StartTime int
	EndTime   int
}

// ScheduleEvents Solves the problem in O(n*Log n) time and O(1) space.
func ScheduleEvents(events []Event) []Event {
	slices.SortFunc(events, func(a, b Event) int {
		return cmp.Compare(a.EndTime, b.EndTime)
	})

	var schedule []Event
	for _, event := range events {
		if len(schedule) == 0 || event.StartTime >= schedule[len(schedule)-1].EndTime {
			schedule = append(schedule, event)
		}
	}
	return schedule
}
