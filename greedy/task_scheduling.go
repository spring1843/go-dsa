package greedy

import "sort"

type Event struct {
	Name      string
	StartTime int
	EndTime   int
}

func ScheduleEvents(events []Event) []Event {
	sort.Slice(events, func(i, j int) bool {
		return events[i].EndTime < events[j].EndTime
	})

	var schedule []Event
	for _, event := range events {
		if len(schedule) == 0 || event.StartTime >= schedule[len(schedule)-1].EndTime {
			schedule = append(schedule, event)
		}
	}
	return schedule
}
