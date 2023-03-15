package array

// FindDuplicate Finds the duplicate in a list of integers (1,n)
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
