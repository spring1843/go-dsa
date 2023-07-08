package recursion

import (
	"testing"
)

/*
TestClimbingStairs tests solution(s) with the following signature and problem description:

	func ClimbingStairs(n int) int {

Given n the number of steps, return in how many ways you can climb these stairs if you are
only able to climb 1 or 2 steps at a time.
*/
func TestClimbingStairs(t *testing.T) {
	tests := []struct {
		n, ways int
	}{
		{0, 1},
		{2, 2},
		{3, 3},
		{5, 8},
		{7, 21},
		{10, 89},
	}

	for i, test := range tests {
		if got := ClimbingStairs(test.n); got != test.ways {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.ways, got)
		}
	}
}
