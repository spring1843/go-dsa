package array

func EqualSubArrays(list []int) [][]int {
	out := make([][]int, 0)
	if len(list) < 2 {
		return out
	}
	sumArr := 0
	for _, el := range list {
		sumArr += el
	}
	if sumArr%2 != 0 {
		return out
	}
	halfSum := sumArr / 2
	runSum := 0
	for i, num := range list {
		runSum += num
		if runSum == halfSum {
			return [][]int{list[:i+1], list[i+1:]}
		}
	}
	return out
}
