package greedy

import (
	"slices"
	"testing"
)

/*
TestScheduleEvents tests solution(s) with the following signature and problem description:

	func ScheduleEvents(events []Event) []Event

Given a list of named tasks with their start and end timing like `{A, 1, 3}, {B, 2, 3}, {C, 3, 4}`
(Task A starts at 1 and ends at 3). return a schedule that includes as many events as
possible like `{A 1 3}, {C 3 4}`.
*/
func TestScheduleEvents(t *testing.T) {
	tests := []struct {
		events   []Event
		schedule []Event
	}{
		{[]Event{}, nil},
		{[]Event{{Name: "A", StartTime: 1, EndTime: 2}}, []Event{{Name: "A", StartTime: 1, EndTime: 2}}},
		{[]Event{{Name: "A", StartTime: 1, EndTime: 3}, {Name: "B", StartTime: 2, EndTime: 3}, {Name: "C", StartTime: 3, EndTime: 4}}, []Event{{Name: "A", StartTime: 1, EndTime: 3}, {Name: "C", StartTime: 3, EndTime: 4}}},
		{[]Event{{Name: "A", StartTime: 1, EndTime: 2}, {Name: "B", StartTime: 1, EndTime: 2}}, []Event{{Name: "A", StartTime: 1, EndTime: 2}}},
		{[]Event{{Name: "A", StartTime: 1, EndTime: 2}, {Name: "B", StartTime: 3, EndTime: 4}}, []Event{{Name: "A", StartTime: 1, EndTime: 2}, {Name: "B", StartTime: 3, EndTime: 4}}},
		{[]Event{{Name: "A", StartTime: 1, EndTime: 2}, {Name: "B", StartTime: 1, EndTime: 3}, {Name: "C", StartTime: 5, EndTime: 6}}, []Event{{Name: "A", StartTime: 1, EndTime: 2}, {Name: "C", StartTime: 5, EndTime: 6}}},
		{[]Event{{Name: "A", StartTime: 1, EndTime: 3}, {Name: "B", StartTime: 1, EndTime: 3}, {Name: "C", StartTime: 3, EndTime: 9}, {Name: "D", StartTime: 5, EndTime: 10}}, []Event{{Name: "A", StartTime: 1, EndTime: 3}, {Name: "C", StartTime: 3, EndTime: 9}}},
	}

	for i, test := range tests {
		if got := ScheduleEvents(test.events); !slices.Equal(got, test.schedule) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.schedule, got)
		}
	}
}
