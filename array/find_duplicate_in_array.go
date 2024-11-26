package array

// FindDuplicate solves the problem in O(n) time and O(1) space.
func FindDuplicate(list []int) int {
	for _, item := range list {
		itemIndex := abs(item) - 1
		if list[itemIndex] < 0 {
			return item
		}
		list[itemIndex] *= -1
	}
	return -1
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
