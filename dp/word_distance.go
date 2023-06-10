package dp

// WordDistance returns how many character modifications (insert, delete, edit) can
// be done on the first input string so that it becomes equal to the second string
func WordDistance(input1, input2 string) int {
	dp := createDP(input2, input1)
	for i := 1; i <= len(input1); i++ {
		for j := 1; j <= len(input2); j++ {
			if input1[i-1] == input2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j], // delete
					min(dp[i][j-1], // insert
						dp[i-1][j-1])) // edit
			}
		}
	}
	return dp[len(input1)][len(input2)]
}

func createDP(input2 string, input1 string) [][]int {
	dp := make([][]int, len(input1)+1)

	for i := range dp {
		dp[i] = make([]int, len(input2)+1)
	}

	for i := 0; i <= len(input1); i++ {
		dp[i][0] = i
	}

	for j := 0; j <= len(input2); j++ {
		dp[0][j] = j
	}

	return dp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
