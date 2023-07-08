package graph

import (
	"reflect"
	"testing"
)

/*
TestRemoveInvalidParentheses tests solution(s) with the following signature and problem description:

	func RemoveInvalidParentheses(input string) []string

Given a string containing parentheses and other alphabet letters like `(z)())()`, remove the minimum
amount of parentheses to make the string valid like `(z())()` and `(z)()()`.
*/
func TestRemoveInvalidParentheses(t *testing.T) {
	tests := []struct {
		input   string
		outputs []string
	}{
		{"", []string{}},
		{")(", []string{""}},
		{")v(", []string{"v"}},
		{"(v)", []string{"(v)"}},
		{"(z)())()", []string{"(z())()", "(z)()()"}},
		{"()()z)()", []string{"(()z)()", "()(z)()", "()()z()"}},
		{"()())()", []string{"(())()", "()()()"}},
	}

	for i, test := range tests {
		if got := RemoveInvalidParentheses(test.input); !reflect.DeepEqual(got, test.outputs) {
			t.Errorf("Failed test case #%d. Want %#v got %#v", i, test.outputs, got)
		}
	}
}
