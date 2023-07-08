package stack

import (
	"testing"
)

/*
TestLongestValidParentheses tests solution(s) with the following signature and problem description:

	func LongestValidParentheses(input string) int

Given a string containing parentheses, find the length of the longest valid (well-formed)
parentheses substring.
*/
func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		input                   string
		longestValidParentheses int
	}{
		{"", 0},
		{"(", 0},
		{"()", 2},
		{"())", 2},
		{"(()", 2},
		{"((()", 2},
		{"((())", 4},
		{"()(()", 2},
		{"((())()", 6},
		{"((())()(((", 6},
		{"))((())()(((", 6},
	}

	for i, test := range tests {
		if got := LongestValidParentheses(test.input); got != test.longestValidParentheses {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.longestValidParentheses, got)
		}
	}
}
