package backtracking

// GenerateParentheses solves the problem in O(2^n) time and O(n) space.
func GenerateParentheses(n int) []string {
	return generateParenthesesRecursive([]string{}, "", 0, 0, n)
}

func generateParenthesesRecursive(output []string, cur string, left, right, n int) []string {
	if len(cur) == n*2 {
		output = append(output, cur)
		return output
	}

	if left < n {
		output = generateParenthesesRecursive(output, cur+"(", left+1, right, n)
	}

	if right < left {
		output = generateParenthesesRecursive(output, cur+")", left, right+1, n)
	}

	return output
}
