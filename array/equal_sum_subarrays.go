package array

// EqualSubArrays solves the problem in O(n^2) time and O(1) space.
func EqualSubArrays(list []int) [][]int {
	out := make([][]int, 0)
	if len(list) < 2 {
		return out
	}

	totalSum := 0
	for _, el := range list {
		totalSum += el
	}

	if totalSum%2 != 0 {
		return out
	}

	halfSum := totalSum / 2
	runSum := 0

	for i := 0; i < len(list)-1; i++ {
		runSum += list[i]

		if runSum == halfSum {

			subArr1 := make([]int, i+1)
			copy(subArr1, list[:i+1])
			subArr2 := make([]int, len(list)-i-1)
			copy(subArr2, list[i+1:])

			if len(subArr2) > 0 {
				out = append(out, subArr1, subArr2)
			}
		}
	}

	return out
}
