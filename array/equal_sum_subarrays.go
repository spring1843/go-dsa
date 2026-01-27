package array

// EqualSubArrays solves the problem in O(n) time and O(1) space.
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
