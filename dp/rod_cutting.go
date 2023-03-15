package dp

import (
	"math"
)

func CutRod(prices []int, n int) int {
	revenues := make([]int, n+1)
	revenues[0] = 0
	for i := 1; i <= n; i++ {
		maximumRevenue := math.MinInt
		for j := 1; j <= i; j++ {
			maximumRevenue = max(maximumRevenue, prices[j-1]+revenues[i-j])
		}
		revenues[i] = maximumRevenue
	}
	return revenues[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
