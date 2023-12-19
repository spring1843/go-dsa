package array

// FindDuplicate solves the problem in O(n) time and O(1) space.
func FindDuplicate(list []int) int {
	if len(list) == 0 {
		return -1
	}

	for _, item := range list {
		itemIndex := item - 1
		// Check for Invalid index if any, return -1 indicating an error.
		if itemIndex < 0 || itemIndex >= len(list) {
			return -1
		}
		if list[itemIndex] < 0 {
			return list[itemIndex] * -1
		}
		list[itemIndex] *= -1
	}
	return -1
}
