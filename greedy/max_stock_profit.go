package greedy

// MaxStockProfit solves the problem in O(n) time and O(1) space.
func MaxStockProfit(prices []int) int {
	maxProfit := 0
	if len(prices) <= 1 {
		return maxProfit
	}
	minStockPrice := prices[0]

	for _, price := range prices {
		profit := price - minStockPrice
		if profit > maxProfit {
			maxProfit = profit
		}
		if price < minStockPrice {
			minStockPrice = price
		}
	}
	return maxProfit
}
