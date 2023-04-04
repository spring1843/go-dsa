package stack

import "testing"

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
