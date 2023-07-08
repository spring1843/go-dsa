package array

// InsertionSort solves the problem in O(n^2) time and O(1) space.
func InsertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		key := list[i]
		j := i - 1

		for j >= 0 && list[j] > key {
			list[j+1] = list[j]
			j--
		}

		list[j+1] = key
	}

	return list
}
