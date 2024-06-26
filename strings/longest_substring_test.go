package strings

import "testing"

/*
TestLongestSubstrings tests solution(s) with the following signature and problem description:

	func LongestSubstringOfTwoChars(input string) string

Given a string like "aabbc" return the longest substring of two unique characters like "aabb".
*/
func TestLongestSubstrings(t *testing.T) {
	tests := []struct {
		input            string
		longestSubstring string
	}{
		{"", ""},
		{"a", ""}, // No occurrence of sequence of two characters
		{"ab", "ab"},
		{"abcc", "bcc"},
		{"aabbc", "aabb"},
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
