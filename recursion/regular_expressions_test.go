package recursion

import (
	"testing"
)

func TestRegularExpressions(t *testing.T) {
	tests := []struct {
		input, pattern string
		match          bool
	}{
		{"", "", true},
		{"", "*", false},
		{"a", "", false},
		{"a", ".", true},
		{"a", "*", false},
		{"aa", "*", false},
		{"aa", "*a", false},
		{"aa", "a*", true},
		{"aa", ".", false},
		{"ab", ".", false},
		{"ad", "d", false},
		{"ad", ".d", true},
		{"da", "*d", false},
		{"da", ".*", true},
		{"ad", "d", false},
		{"abdef", "*d*", false},
	}

	for i, test := range tests {
		if got := IsRegularExpressionMatch(test.input, test.pattern); got != test.match {
			t.Errorf("Failed test case %d. Want %t got %t", i, test.match, got)
		}
	}
}
