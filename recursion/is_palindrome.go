package recursion

// IsPalindrome solves the problem in O(n) time and O(n) space.
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
