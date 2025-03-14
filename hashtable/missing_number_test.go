package hashtable

import "testing"

/*
TestMissingNumber tests solution(s) with the following signature and problem description:

	func MissingNumber(numbers []int) int

Given an unsorted slice of numbers like return the missing integer.

For example given {7,5,3,4,1,2,0,-1}, return the missing integer like 6.
*/
func TestMissingNumber(t *testing.T) {
	tests := []struct {
		numbers []int
		missing int
	}{
		{[]int{}, -1},
		{[]int{2}, -1},
		{[]int{1}, -1},
		{[]int{1, 2}, -1},
		{[]int{1, 2, 3, 4, 6}, 5},
		{[]int{1, 2, 3, 4, 5, 6}, -1},
		{[]int{6, 3, 2, 1, 4}, 5},
		{[]int{-3, -1, 0, 1, 2}, -2},
		{[]int{-3, -2, -1, 1, 2}, 0},
	}

	for i, test := range tests {
		got := MissingNumber(test.numbers)
		if got != test.missing {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.missing, got)
		}
	}
}
