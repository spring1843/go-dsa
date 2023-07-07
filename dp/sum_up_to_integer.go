package dp

// SumUpToInteger solves the problem in O(n*sum) time and O(sum) space.
func SumUpToInteger(numbers []int, sum int) bool {
	dp := make([]bool, sum+1)
	dp[0] = true

	for i := 1; i <= sum; i++ {
		if numbers[0] == i {
			dp[i] = true
		}
	}

	for i := 1; i < len(numbers); i++ {
		for j := sum; j >= 0; j-- {
			if !dp[j] && j >= numbers[i] {
				dp[j] = dp[j-numbers[i]]
			}
		}
	}

	return dp[sum]
}
