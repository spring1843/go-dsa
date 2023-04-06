package recursion

import (
	"testing"
)

func TestReverseDigits(t *testing.T) {
	tests := []struct {
		n, reversed int
	}{
		{0, 0},
		{1, 1},
		{12, 21},
		{112, 211},
		{110, 11},
	}

	for i, test := range tests {
		if got := ReverseDigits(test.n); got != test.reversed {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.reversed, got)
		}
	}
}
