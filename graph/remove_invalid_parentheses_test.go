package graph

import (
	"reflect"
	"testing"
)

func TestRemoveInvalidParentheses(t *testing.T) {
	var tests = []struct {
		input   string
		outputs []string
	}{
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
