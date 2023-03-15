package recursion

// IsRegularExpressionMatch returns true if the given input matches the given pattern and false otherwise
// In this regular expression patten
//
//	. indicates any single character
//	* matches zero or more characters of the preceding character
func IsRegularExpressionMatch(input, pattern string) bool {
	if pattern == "" {
		return input == ""
	}
	firstMatch := len(input) > 0 && (input[0] == pattern[0] || pattern[0] == '.')
	if len(pattern) < 2 {
		return firstMatch && IsRegularExpressionMatch(input[1:], pattern[1:])
	}
	return IsRegularExpressionMatch(input, pattern[2:]) || (firstMatch && IsRegularExpressionMatch(input[1:], pattern))
}
