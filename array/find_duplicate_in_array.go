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

// Note: This solution uses Marking Technique
// 	Pros: No extra space required.
// 	Cons: Modifies the original list, Only works if the item value is between 0 and len(list)-1
// A more general solution would be to use a set. But the space complexity will become O(n).
