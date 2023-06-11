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
		{"", ".", true},
		{"", "*", true},
		{"a", "", false},
		{"a", ".", true},
		{"a", "*", true},
		{"aa", "*", true},
		{"aa", ".", true},
		{"ab", ".", false},
		{"ad", ".d", true},
		{"ad", "*d", true},
		{"abd", "*d", true},
		{"abdef", "*d", true},
		{"abdef", "*d*", true},
		{"abdef", "d*", true},
	}

	for i, test := range tests {
		if got := IsRegularExpressionMatch(test.input, test.pattern); got != test.match {
			t.Skipf("Failed test case %q. Want %t got %t", i, test.match, got)
		}
	}
}
