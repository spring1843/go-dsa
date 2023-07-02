package dp

import (
	"testing"
)

func TestWordDistance(t *testing.T) {
	tests := []struct {
		input1   string
		input2   string
		distance int
	}{
		{"", "", 0},
		{"", "abcde", 5},
		{"a", "a", 0},
		{"a", "b", 1},
		{"ab", "ac", 1},
		{"ab", "dc", 2},
		{"ab", "dcd", 3},
		{"abcdef", "abcde", 1},
		{"gabcdef", "abcde", 2},
	}

	for i, test := range tests {
		if got := WordDistance(test.input1, test.input2); got != test.distance {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.distance, got)
		}
	}
}
