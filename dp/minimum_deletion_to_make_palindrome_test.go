package dp

import (
	"testing"
)

func TestMinimumDeletionsToMakePalindrome(t *testing.T) {
	tests := []struct {
		input                           string
		minChangeNeededToMakePalindrome int
	}{
		{"", 0},
		{"a", 0},
		{"ab", 1},
		{"acb", 2},
		{"abccb", 1},
		{"abcba", 0},
		{"abdcba", 1},
		{"abdecba", 2},
		{"zabdecbaz", 2},
	}

	for i, test := range tests {
		if got := MinimumDeletionsToMakePalindrome(test.input); got != test.minChangeNeededToMakePalindrome {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.minChangeNeededToMakePalindrome, got)
		}
	}
}
