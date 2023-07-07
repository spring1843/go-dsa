package dnc

// MergeSort solves the problem in O(n log n) time and O(n) space.
func MergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	left, right := divide(list)
	return merge(left, right)
}

func divide(list []int) ([]int, []int) {
	mid := len(list) / 2
	left := MergeSort(list[:mid])
	right := MergeSort(list[mid:])
	return left, right
}

func merge(left, right []int) []int {
	output := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			output = append(output, left[i])
			i++
		} else {
			output = append(output, right[j])
			j++
		}
	}

	if i >= len(left) {
		output = append(output, right[j:]...)
	} else {
		output = append(output, left[i:]...)
	}
	return output
}
