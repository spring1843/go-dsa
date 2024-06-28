package array

// EqualSubArrays solves the problem in O(n^2) time and O(1) space.
func EqualSubArrays(list []int) [][]int {
	output := make([][]int, 0)
	if len(list) < 2 {
		return output
	}

	splitPoint := findSplitPoint(list)
	if splitPoint == -1 || splitPoint == len(list) {
		return output
	}

	output = append(output, list[0:splitPoint])
	output = append(output, list[splitPoint:])

	return output
}

func findSplitPoint(list []int) int {
	lSum := 0
	for i := range len(list) {
		lSum += list[i]

		rSum := 0
		for j := i + 1; j < len(list); j++ {
			rSum += list[j]
		}

		if lSum == rSum {
			return i + 1
		}
	}
	return -1
}
