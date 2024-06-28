package dnc

// QuickSort solves the problem in O(n*Log n) time and O(n) space.
func QuickSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	pivot := list[len(list)/2]

	var less, equal, greater []int
	for i := range len(list) {
		if list[i] == pivot {
			equal = append(equal, list[i])
		}
		if list[i] < pivot {
			less = append(less, list[i])
		}
		if list[i] > pivot {
			greater = append(greater, list[i])
		}
	}

	return append(append(QuickSort(less), equal...), QuickSort(greater)...)
}
