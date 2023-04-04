package stack

import (
	"reflect"
	"testing"
)

func TestInfixToPostfix(t *testing.T) {
	var tests = []struct {
		infix     []string
		expectErr bool
		postfix   []string
	}{
		{[]string{""}, false, []string{""}},
		{[]string{"a", "+", "b"}, false, []string{"a", "b", "+"}},
		{[]string{"a", "-", "b", "+", "c"}, false, []string{"a", "b", "c", "+", "-"}},
		{[]string{"a", "-", "(", "b", "+", "c", ")"}, false, []string{"a", "b", "c", "+", "-"}},
		{[]string{"a", "+", "b", "-", "c"}, false, []string{"a", "b", "c", "-", "+"}},
		{[]string{"a", "/", "b"}, false, []string{"a", "b", "/"}},
		{[]string{"1", "*", "2", "+", "3", "+", "4", "*", "5"}, false, []string{"1", "2", "3", "4", "5", "*", "+", "+", "*"}},
		{[]string{"1", "*", "(", "2", "+", "3", ")", "+", "4", "*", "5"}, false, []string{"1", "2", "3", "+", "4", "5", "*", "+", "*"}},
	}

	for i, test := range tests {
		got, err := InfixToPostfix(test.infix)
		if err != nil && !test.expectErr {
			t.Fatalf("No error expected. Error: %s", err)
		}
		if !reflect.DeepEqual(got, test.postfix) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.postfix, got)
		}
	}
}
