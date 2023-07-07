package array

// ProductOfAllOtherElements solves the problem in O(n) time and O(1) space.
func ProductOfAllOtherElements(list []int) []int {
	if len(list) == 0 {
		return list
	}

	right := make([]int, len(list))

	product := 1
	for i := len(list) - 1; i >= 0; i-- {
		product *= list[i]
		right[i] = product
	}

	product = list[0]
	for i := 0; i < len(list); i++ {
		if i == 0 {
			list[i] = right[1]
			continue
		}
		value := list[i]

		if i != len(list)-1 {
			list[i] = right[i+1] * product
		} else {
			list[i] = product
		}

		product *= value
	}

	return list
}
