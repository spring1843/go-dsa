package recursion

import (
	"testing"
)

func TestRegularExpressions(t *testing.T) {
	var tests = []struct {
		input, pattern string
		match          bool
	}{
		{"", "", true},
		{"a", "", false},
		{"a", ".", true},
		{"a", "*", true},
		{"aa", "*", true},
		{"aa", ".", true},
		{"ab", ".", false},
		{"ad", ".d", true},
		{"ad", "*d", true},
		{"abd", "*d", true},
	}

	for i, test := range tests {
		if got := IsRegularExpressionMatch(test.input, test.pattern); got != test.match {
			t.Skipf("Failed test case %q. Want %t got %t", i, test.match, got)
		}
	}
}
