package array

import (
	"slices"
	"testing"
)

/*
TestReverseInPlace tests solution(s) with the following signature and problem description:

	ReverseInPlace(list []int, start, end int)

Given a slice of integers, a start index, and an end index, reverse the integers in the in-place
without using any extra memory.

For example given {1,2,3,4,5,6} and start of 0 and end of 4, return {5,4,3,2,1,6} because:

Reverse of items from index 0 to 4 is {5,4,3,2,1} and the remaining item {6} remain unchanged, so
the the resulting slice is {5,4,3,2,1,6}.
*/
func TestReverseInPlace(t *testing.T) {
	tests := []struct {
		list     []int
		start    int
		end      int
		reversed []int
	}{
		{[]int{}, 0, 0, []int{}},
		{[]int{1, 2, 3}, 1, 2, []int{1, 3, 2}},
		{[]int{1, 2, 3}, 2, 1, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 1, 2, []int{1, 3, 2, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 0, 4, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6}, 0, 4, []int{5, 4, 3, 2, 1, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, 0, 3, []int{4, 3, 2, 1, 5, 6}},
	}

	for i, test := range tests {
		ReverseInPlace(test.list, test.start, test.end)
		if !slices.Equal(test.list, test.reversed) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.reversed, test.list)
		}
	}
}
