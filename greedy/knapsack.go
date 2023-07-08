package greedy

import "sort"

// KnapsackItem is an item with a weight and value that can be put in a knapsack.
type KnapsackItem struct {
	Weight int
	Value  int
}

// Knapsack solves the problem in O(n*Log n)  time and O(1) space.
func Knapsack(items []KnapsackItem, capacity int) int {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Value/items[i].Weight > items[j].Value/items[j].Weight
	})

	var maxValue int
	for _, item := range items {
		if capacity < item.Weight {
			maxValue += (capacity * item.Value) / item.Weight
			break
		}
		maxValue += item.Value
		capacity -= item.Weight
	}

	return maxValue
}
