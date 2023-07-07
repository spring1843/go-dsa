package recursion

// IsPalindrome solves the problem in O(n) time and O(n) space.
func IsPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}

	if len(s) == 2 {
		return s[0] == s[1]
	}

	if s[0] != s[len(s)-1] {
		return false
	}

	return IsPalindrome(s[1 : len(s)-1])
}
