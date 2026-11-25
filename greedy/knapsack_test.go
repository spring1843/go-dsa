package greedy

import "testing"

/*
TestKnapsack tests solution(s) with the following signature and problem description:

	func Knapsack(items []KnapsackItem, capacity int) int

Given weight capacity of a knapsack like 5, a list of divisible items (such as pieces
of metal) with values and weights like `{Value: 6, Weight: 2}, {Value: 10, Weight: 2},
{Value: 12, Weight: 3}` return the maximum value of items that can be placed in the
knapsack like `22`.
*/
func TestKnapsack(t *testing.T) {
	tests := []struct {
		items     []KnapsackItem
		capacity  int
		maxProfit int
	}{
		{[]KnapsackItem{}, 0, 0},
		{[]KnapsackItem{{Value: 1, Weight: 1}}, 1, 1},
		{[]KnapsackItem{{Value: 1, Weight: 1}, {Value: 10, Weight: 5}}, 5, 10},
		{[]KnapsackItem{{Value: 1, Weight: 1}, {Value: 10, Weight: 5}}, 2, 4},
		{[]KnapsackItem{{Value: 6, Weight: 2}, {Value: 10, Weight: 2}, {Value: 12, Weight: 3}}, 5, 22},
		// Ratios 2.33, 2.5, 2 and 2.5: the two 2.5 items fill the knapsack exactly
		// for 15. Integer division saw four equal ratios and settled for 7+5+2.
		{[]KnapsackItem{{Value: 7, Weight: 3}, {Value: 5, Weight: 2}, {Value: 2, Weight: 1}, {Value: 10, Weight: 4}}, 6, 15},
		{[]KnapsackItem{{Value: 50, Weight: 10}, {Value: 55, Weight: 20}, {Value: 110, Weight: 30}}, 50, 187},
		// Ratios 1.5 and 1.9 both round down to 1 under integer division.
		{[]KnapsackItem{{Value: 15, Weight: 10}, {Value: 19, Weight: 10}}, 10, 19},
		{[]KnapsackItem{{Value: 11, Weight: 10}, {Value: 19, Weight: 10}, {Value: 15, Weight: 10}}, 20, 34},
	}

	for i, test := range tests {
		if got := Knapsack(test.items, test.capacity); got != test.maxProfit {
			t.Fatalf("Failed test case #%d. Want %d got %v", i, test.maxProfit, got)
		}
	}
}
