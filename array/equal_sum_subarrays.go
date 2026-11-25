package array

// EqualSubArrays solves the problem in O(n) time and O(1) space.
func EqualSubArrays(list []int) [][]int {
	if len(list) < 2 {
		return [][]int{}
	}

	splitPoint := findSplitPoint(list)
	if splitPoint == -1 || splitPoint == len(list) {
		return [][]int{}
	}

	return [][]int{list[:splitPoint], list[splitPoint:]}
}

func findSplitPoint(list []int) int {
	lSum, rSum := 0, 0

	for _, n := range list {
		rSum += n
	}

	for i, n := range list {
		lSum += n
		rSum -= n

		if lSum == rSum {
			return i + 1
		}
	}
	return -1
}
