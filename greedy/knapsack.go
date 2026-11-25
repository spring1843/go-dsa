package greedy

import (
	"cmp"
	"slices"
)

// KnapsackItem is an item with a weight and value that can be put in a knapsack.
type KnapsackItem struct {
	Weight int
	Value  int
}

// Knapsack solves the problem in O(n*Log n) time and O(1) space.
func Knapsack(items []KnapsackItem, capacity int) int {
	// Order by Value/Weight, highest first. The ratios are compared by cross
	// multiplication: integer division would round 1.5 and 1.9 both down to 1
	// and treat items with different ratios as equal.
	slices.SortFunc(items, func(a, b KnapsackItem) int {
		return cmp.Compare(b.Value*a.Weight, a.Value*b.Weight)
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
