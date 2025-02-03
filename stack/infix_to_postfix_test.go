package stack

import (
	"slices"
	"testing"
)

/*
TestInfixToPostfix tests solution(s) with the following signature and problem description:

	func InfixToPostfix(infix []string) []string

Given an infix expression convert it to a postfix expression supporting the four basic arithmetic
operations and parentheses.

Infix expression is how humans typically write arithmetic expressions like 1*2+3+4*5 which
is equivalent of (1*2) + 3 + (4*5).

For example given 1*2+3+4*5, return 1 2 * 3 + 4 5 * which both evaluate to 25.
*/
func TestInfixToPostfix(t *testing.T) {
	tests := []struct {
		infix   []string
		postfix []string
	}{
		{[]string{""}, []string{""}},
		{[]string{"a", "+", "b"}, []string{"a", "b", "+"}},
		{[]string{"a", "-", "b", "+", "c"}, []string{"a", "b", "c", "+", "-"}},
		{[]string{"a", "-", "(", "b", "+", "c", ")"}, []string{"a", "b", "c", "+", "-"}},
		{[]string{"a", "+", "b", "-", "c"}, []string{"a", "b", "c", "-", "+"}},
		{[]string{"a", "/", "b"}, []string{"a", "b", "/"}},
		{[]string{"1", "*", "2", "+", "3", "+", "4", "*", "5"}, []string{"1", "2", "3", "4", "5", "*", "+", "+", "*"}},
		{[]string{"1", "*", "(", "2", "+", "3", ")", "+", "4", "*", "5"}, []string{"1", "2", "3", "+", "4", "5", "*", "+", "*"}},
	}

	for i, test := range tests {
		if got := InfixToPostfix(test.infix); !slices.Equal(got, test.postfix) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.postfix, got)
		}
	}
}
