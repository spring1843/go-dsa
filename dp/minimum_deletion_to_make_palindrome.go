package dp

// MinimumDeletionsToMakePalindrome returns how many deletions can be done in the input string
// to make it a palindrome
func MinimumDeletionsToMakePalindrome(input string) int {
	if len(input) <= 1 {
		return 0
	}
	return len(input) - findLength(input)
}

func findLength(input string) int {
	dp := make([][]int, len(input))

	for i := range input {
		dp[i] = make([]int, len(input))
		dp[i][i] = 1
	}

	for start := len(input) - 1; start >= 0; start-- {
		for end := start + 1; end < len(input); end++ {
			if input[start] == input[end] {
				dp[start][end] = 2 + dp[start+1][end-1]
			} else {
				dp[start][end] = max(dp[start+1][end], dp[start][end-1])
			}
		}
	}

	return dp[0][len(input)-1]
}
