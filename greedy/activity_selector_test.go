package greedy

import (
	"reflect"
	"testing"
)

/*
TestActivitySelector tests solution(s) with the following signature and problem description:

	func ActivitySelector(start, finish []int) []int

Given a list of start and finish times for a few activities like {1, 3, 0} and {4, 5, 6}
(first activity starts at 1 and ends at 4, second starts at 3 and ends at 5), return a maximal
list of non-conflicting activities.
*/
func TestActivitySelector(t *testing.T) {
	tests := []struct {
		start      []int
		finish     []int
		activities []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{4}, []int{1}},
		{[]int{1, 3}, []int{4, 5}, []int{1}},
		{[]int{1, 3, 0}, []int{4, 5, 6}, []int{1}},
		{[]int{1, 3, 0, 5, 3, 5, 6, 7, 8, 2, 12}, []int{4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16}, []int{1, 4, 8, 11}},
	}
	for i, test := range tests {
		if got := ActivitySelector(test.start, test.finish); !reflect.DeepEqual(got, test.activities) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.activities, got)
		}
	}
}
