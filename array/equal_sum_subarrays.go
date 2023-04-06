package array

// EqualSubArrays given a list of integers returns two sub arrays that have
// equal sums without changing the order of the items in the list.
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
	for i := 0; i < len(list); i++ {
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
