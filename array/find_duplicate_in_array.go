package array

// FindDuplicate solves the problem in O(n) time and O(1) space.
func FindDuplicate(list []int) int {
	for _, item := range list {
		itemIndex := item - 1
		if list[itemIndex] < 0 {
			return list[itemIndex] * -1
		}
		list[itemIndex] *= -1
	}
	return -1
}
