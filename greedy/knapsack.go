package greedy

import "sort"

// KnapsackItem is an item with a weight and value that can be put in a knapsack.
type KnapsackItem struct {
	Weight int
	Value  int
}

// Knapsack returns the maximum value that can be put in a knapsack of a given capacity.
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
