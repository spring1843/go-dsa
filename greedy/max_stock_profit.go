package greedy

// MaxStockProfit returns the highest amount of profit that can be made by buying and selling
// a stock, given a list of prices at different times
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
