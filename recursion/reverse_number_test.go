package recursion

import "testing"

/*
TestReverseDigits tests solution(s) with the following signature and problem description:

	func ReverseDigits(n int) int

Given an integer reverse the order of the digits.

For example given 123 return 321.
*/
func TestReverseDigits(t *testing.T) {
	tests := []struct {
		n, reversed int
	}{
		{0, 0},
		{1, 1},
		{12, 21},
		{112, 211},
		{110, 11},
		{123, 321},
	}

	for i, test := range tests {
		if got := ReverseDigits(test.n); got != test.reversed {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.reversed, got)
		}
	}
}
