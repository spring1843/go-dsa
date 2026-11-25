package dnc

// BinarySearch solves the problem in O(Log n) time and O(1) space.
func BinarySearch(list []int, search int) int {
	return binarySearchRecursive(list, 0, len(list), search)
}

func binarySearchRecursive(list []int, min, max, search int) int {
	if min == max {
		return -1
	}

	mid := ((max - min) / 2) + min
	if list[mid] == search {
		return mid
	}

	if list[mid] > search {
		return binarySearchRecursive(list, min, mid, search)
	}

	return binarySearchRecursive(list, mid+1, max, search)
}
