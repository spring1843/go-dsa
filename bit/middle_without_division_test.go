package bit

import "testing"

func TestMiddleWithoutDivision(t *testing.T) {
	var tests = []struct {
		a, b, mid int
	}{
		{0, 0, 0},
		{0, 1, 0},
		{0, 5, 2},
		{1, 5, 3},
		{0, 6, 3},
		{1, 5, 3},
		{2, 20, 11},
		{2, -2, 0},
		{-3, -2, -3},
	}

	for i, test := range tests {
		if got := MiddleWithoutDivision(test.a, test.b); got != test.mid {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.mid, got)
		}
	}
}
