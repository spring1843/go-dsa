package strings

import (
	"testing"
)

func TestLongestSubstrings(t *testing.T) {
	tests := []struct {
		input            string
		longestSubstring string
	}{
		{"", ""},
		{"a", ""}, // No occurrence of sequence of two characters
		{"ab", "ab"},
		{"abcc", "bcc"},
		{"ada", "ada"},
		{"dafff", "afff"},
		{"assdeeeddfffha", "deeedd"},
	}

	for i, test := range tests {
		got := LongestSubstringOfTwoChars(test.input)
		if got != test.longestSubstring {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.longestSubstring, got)
		}
	}
}
