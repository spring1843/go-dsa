package dnc

// BinarySearch solves the problem in O(Log n) time and O(1) space.
func BinarySearch(list []int, search int) int {
	return binarySearchRecursive(list, 0, len(list), search)
}

func binarySearchRecursive(list []int, low, high, search int) int {
	if low == high {
		return -1
	}

	mid := ((high - low) / 2) + low
	if list[mid] == search {
		return mid
	}

	if list[mid] > search {
		return binarySearchRecursive(list, low, mid, search)
	}

	return binarySearchRecursive(list, mid+1, high, search)
}
