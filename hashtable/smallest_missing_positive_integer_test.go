package hashtable

import "testing"

/*
TestSmallestMissingPositiveInteger tests solution(s) with the following signature and problem description:

	func SmallestMissingPositiveInteger(input []int) int

Given a slice of integers, return the smallest missing positive integer.

For example given {-9, 0, 1, 2, 3, 4, 5, 6} return 7 because -9 is not a positive integer and
all numbers from 0 to 6 are present in the list.
*/
func TestSmallestMissingPositiveInteger(t *testing.T) {
	tests := []struct {
		numbers []int
		missing int
	}{
		{[]int{}, 1},
		{[]int{0}, 1},
		{[]int{-1}, 1},
		{[]int{2}, 1},
		{[]int{1, 2, 3}, 4},
		{[]int{1, 3}, 2},
		{[]int{-9, 0, 1, 2, 3, 4, 5, 6}, 7},
	}
	for i, test := range tests {
		if got := SmallestMissingPositiveInteger(test.numbers); got != test.missing {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.missing, got)
		}
	}
}
