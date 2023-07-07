package recursion

import (
	"testing"
)

/*
TestPowerOf tests solution(s) with the following signature and problem description:

	func PowerOf(x, n int) int

Raises x to the power of n.
*/
func TestPowerOf(t *testing.T) {
	tests := []struct {
		x, n, result int
	}{
		{1, 0, 1},
		{1, 1, 1},
		{1, 2, 1},
		{2, 1, 2},
		{2, 2, 4},
		{2, 3, 8},
		{3, 3, 27},
		{5, 2, 25},
		{5, 3, 125},
		{5, 4, 625},
	}

	for i, test := range tests {
		if got := PowerOf(test.x, test.n); got != test.result {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.result, got)
		}
	}
}
