package greedy

import (
	"reflect"
	"testing"
)

/*
TestScheduleEvents tests solution(s) with the following signature and problem description:

	func ScheduleEvents(events []Event) []Event

Given a list of events with start and end times, find the maximum number of events that
can be scheduled without overlap.
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
		if got := ScheduleEvents(test.events); !reflect.DeepEqual(got, test.schedule) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.schedule, got)
		}
	}
}
