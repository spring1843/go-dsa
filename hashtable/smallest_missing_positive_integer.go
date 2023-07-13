package hashtable

// SmallestMissingPositiveInteger solves the problem in O(n) time and O(n) space.
func SmallestMissingPositiveInteger(input []int) int {
	if len(input) == 0 {
		return 1
	}
	exist := make(map[int]struct{})
	largest := input[0]
	for _, number := range input {
		if number > largest {
			largest = number
		}
		exist[number] = struct{}{}
	}
	if largest < 0 {
		return 1
	}

	for i := 1; i < largest; i++ {
		if _, ok := exist[i]; !ok {
			return i
		}
	}
	return largest + 1
}
