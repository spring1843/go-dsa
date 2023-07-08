package stack

import "testing"

/*
TestIsExpressionBalanced tests solution(s) with the following signature and problem description:

	func IsExpressionBalanced(s string) bool

Given a set of symbols including []{}(), determine if the input is is balanced or not.
*/
func TestIsExpressionBalanced(t *testing.T) {
	tests := []struct {
		expression string
		isValid    bool
	}{
		{"", true},
		{"()", true},
		{"(){", false},
		{"(){}", true},
		{"(){}]", false},
		{"(){}][", false},
		{"(){}[]", true},
		{"({}[])", true},
		{"({[]})", true},
		{"({[])", false},
		{"(({[]))", false},
		{")({[])", false},
	}

	for i, test := range tests {
		if got := IsExpressionBalanced(test.expression); got != test.isValid {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.isValid, got)
		}
	}
}
