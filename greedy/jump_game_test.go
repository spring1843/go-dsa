package greedy

import (
	"testing"
)

/*
TestJumpGame tests solution(s) with the following signature and problem description:

	func JumpGame(jumps []int) bool

Given a list of integers that representing the maximum posit ions one can jump at any
given position like {1,2,4,2,1} (at position 1 we can jump up to 1, at position 2 we can
jump up to 2), return true if one can reach the last position of the list and false
otherwise. The response is true for this example.
*/
func TestJumpGame(t *testing.T) {
	tests := []struct {
		jumps          []int
		canReachTheEnd bool
	}{
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 0, 1}, false},
		{[]int{3, 1, 0, 0, 1}, false},
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{1, 2, 4, 2, 1}, true},
		{[]int{2, 3, 1, 1, 2, 0, 1}, true},
	}

	for i, test := range tests {
		if got := JumpGame(test.jumps); got != test.canReachTheEnd {
			t.Fatalf("Failed test case #%d. Want %t got %t", i, test.canReachTheEnd, got)
		}
	}
}
