package array

// FindFirstMissingPositive returns the smallest missing positive integer.
// The algorithm runs in O(n) time and O(1) space by placing each value at its correct index.
func FindFirstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			// Swap nums[i] with nums[nums[i] - 1]
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

