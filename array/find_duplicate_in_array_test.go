package array

import (
	"testing"
)

/*
TestFindDuplicate tests solution(s) with the following signature and problem description:

	FindDuplicate(list []int) int

Finds the duplicate in a list of integers (1,n).
*/
func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		list      []int
		duplicate int
	}{
		{[]int{}, -1},
		{[]int{1, 2, 2}, 2},
		{[]int{1, 2, 3}, -1},
		{[]int{1, 1, 2, 3}, 1},
		{[]int{1, 2, 2, 3}, 2},
	}

	for i, test := range tests {
		if got := FindDuplicate(test.list); got != test.duplicate {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.duplicate, got)
		}
	}
}
