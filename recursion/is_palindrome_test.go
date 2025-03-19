package recursion

import "testing"

/*
TestIsPalindrome tests solution(s) with the following signature and problem description:

	func IsPalindrome(s string) bool

Given a string return true if it's a palindrome and false otherwise.

For example given `abba` return true. Given `abca` return false.
*/
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s            string
		isPalindrome bool
	}{
		{"", true},
		{"1", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"aba", true},
		{"abba", true},
		{"abca", false},
		{"acba", false},
		{"acca", true},
		{"acdca", true},
		{"acedeca", true},
		{"acedec", false},
	}

	for i, test := range tests {
		if got := IsPalindrome(test.s); got != test.isPalindrome {
			t.Fatalf("Failed test case #%d. Want %t got %t", i, test.isPalindrome, got)
		}
	}
}
