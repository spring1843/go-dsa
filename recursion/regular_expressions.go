package recursion

// IsRegularExpressionMatch solves the problem in O(n) time and O(n) space.
func IsRegularExpressionMatch(input, pattern string) bool {
	if len(pattern) == 0 {
		return len(input) == 0
	}
	firstMatch := len(input) > 0 && (pattern[0] == input[0] || pattern[0] == '.')
	if len(pattern) >= 2 && pattern[1] == '*' {
		return IsRegularExpressionMatch(input, pattern[2:]) || (firstMatch && IsRegularExpressionMatch(input[1:], pattern))
	}
	return firstMatch && IsRegularExpressionMatch(input[1:], pattern[1:])
}
