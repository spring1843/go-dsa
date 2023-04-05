package greedy

import "testing"

func TestMaxStockProfit(t *testing.T) {
	tests := []struct {
		prices    []int
		maxProfit int
	}{
		{[]int{}, 0},
		{[]int{1}, 0},
		{[]int{1, 3}, 2},
		{[]int{1, 3, 1}, 2},
		{[]int{1, 3, 1, 2, 3, 8, 2, 12}, 11},
		{[]int{50, 10, 4, 100, 1, 101, 5, 10}, 100},
	}

	for i, test := range tests {
		if got := MaxStockProfit(test.prices); got != test.maxProfit {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.maxProfit, got)
		}
	}
}
