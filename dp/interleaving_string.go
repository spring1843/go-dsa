package dp

// IsInterleavingString solves the problem in O(m*n) time and O(m*n) space.
func IsInterleavingString(a, b, input string) bool {
	m, n := len(a), len(b)
	if len(input) != m+n {
		return false
	}

	dp := dpMap(m, n, a, b, input)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == input[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}
			if b[j-1] == input[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

func dpMap(m, n int, a, b, input string) [][]bool {
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= m; i++ {
		if a[i-1] == input[i-1] {
			dp[i][0] = dp[i-1][0]
		}
	}

	for j := 1; j <= n; j++ {
		if b[j-1] == input[j-1] {
			dp[0][j] = dp[0][j-1]
		}
	}
	return dp
}
