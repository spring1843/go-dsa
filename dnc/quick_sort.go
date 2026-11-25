package dnc

import "slices"

// QuickSort solves the problem in O(n*Log n) time and O(n) space.
func QuickSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	pivot := list[len(list)/2]

	var less, equal, greater []int
	for _, value := range list {
		if value == pivot {
			equal = append(equal, value)
		}
		if value < pivot {
			less = append(less, value)
		}
		if value > pivot {
			greater = append(greater, value)
		}
	}

	return slices.Concat(QuickSort(less), equal, QuickSort(greater))
}
