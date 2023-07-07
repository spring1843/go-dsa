package backtracking

// GenerateParenthesis solves the problem in O(2^n) time and O(n) space.
func GenerateParenthesis(n int) []string {
	return generateParenthesisRecursive([]string{}, "", 0, 0, n)
}

func generateParenthesisRecursive(output []string, cur string, left, right, n int) []string {
	if len(cur) == n*2 {
		output = append(output, cur)
		return output
	}

	if left < n {
		output = generateParenthesisRecursive(output, cur+"(", left+1, right, n)
	}

	if right < left {
		output = generateParenthesisRecursive(output, cur+")", left, right+1, n)
	}

	return output
}
