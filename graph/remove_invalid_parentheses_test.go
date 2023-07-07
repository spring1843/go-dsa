package graph

import (
	"reflect"
	"testing"
)

/*
TestRemoveInvalidParentheses tests solution(s) with the following signature and problem description:

	func RemoveInvalidParentheses(input string) []string

Removes the minimum number of invalid parentheses to make the input valid.
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
