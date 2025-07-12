package array

import (
	"testing"
)

/*
TestFindFirstMissingPositive tests solution(s) with the following signature:

	FindFirstMissingPositive(nums []int) int

Given an unsorted integer slice, return the smallest missing positive integer.

Examples:
Input: []int{1, 2, 0}           → Output: 3
Input: []int{3, 4, -1, 1}       → Output: 2
Input: []int{7, 8, 9, 11, 12}   → Output: 1
*/
func TestFindFirstMissingPositive(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
		{[]int{}, 1},
		{[]int{1, 2, 3}, 4},
		{[]int{0, -1, -2}, 1},
		{[]int{2, 1}, 3},
		{[]int{1}, 2},
		{[]int{2}, 1},
		{[]int{3, 4, -1, 1, 2}, 5},
	}

	for i, test := range tests {
		if got := FindFirstMissingPositive(append([]int(nil), test.nums...)); got != test.expected {
			t.Fatalf("Failed test case #%d. Input: %v. Want %d, got %d", i, test.nums, test.expected, got)
		}
	}
}

