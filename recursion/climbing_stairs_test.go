package recursion

import (
	"testing"
)

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
