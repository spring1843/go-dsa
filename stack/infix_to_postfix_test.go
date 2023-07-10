package stack

import (
	"reflect"
	"testing"
)

/*
TestInfixToPostfix tests solution(s) with the following signature and problem description:

	func InfixToPostfix(infix []string) []string

Given an infix expression e.g. 1*2+3+4*5, convert it to a postfix expression
like 1 2 * 3 + 4 5 * supporting the four basic arithmetic operations and parentheses.
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
		if got := InfixToPostfix(test.infix); !reflect.DeepEqual(got, test.postfix) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.postfix, got)
		}
	}
}
