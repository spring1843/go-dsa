package dp

import (
	"testing"
)

/*
TestMinimumDeletionsToMakePalindrome tests solution(s) with the following signature and problem description:

	func MinimumDeletionsToMakePalindrome(input string) int

Given a string like abccb return the minimum number of character deletions that can be done on the string
to make it a palindrome like 1 (by removing a, we will have bccb).
*/
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
