package dp

import "testing"

/*
TestIsInterleavingString tests solution(s) with the following signature and problem description:

	IsInterleavingString(a, b, input string) bool

Given three strings a,b, and input, return true if input is made by interleaving a and b such that
it contains all their characters while maintaining their order.

Example:
"a", "bcd", "abcd" would return true.
"a", "bdc", "abd" would return false, because c is missing.
"a", "bdc", "abc" would return false, because d is missing.
*/
func TestIsInterleavingString(t *testing.T) {
	tests := []struct {
		a, b, input          string
		isInterleavingString bool
	}{
		{"a", "b", "ab", true},
		{"abc", "def", "abcdef", true},
		{"ac", "bdef", "abcdef", true},
		{"abe", "cdf", "abcdef", true},
		{"abd", "cef", "abcdef", true},
		{"a", "bdc", "abdc", true},
		{"a", "bdc", "abd", false},
		{"a", "bdc", "abc", false},
		{"bdc", "a", "abc", false},
		{"bcd", "a", "abcd", true},
		{"abcd", "a", "aabcd", true},
		{"aaac", "aaab", "aaaaaabc", true},
	}

	for i, test := range tests {
		if got := IsInterleavingString(test.a, test.b, test.input); test.isInterleavingString != got {
			t.Fatalf("Failed test case #%d. Want %t got %t", i, test.isInterleavingString, got)
		}
	}
}
