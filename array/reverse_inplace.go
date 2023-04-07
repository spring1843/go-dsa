package array

// ReverseInplace reverses parts of the array in place.
func ReverseInplace(nums []int, start, end int) {
	for i := start; i <= start+end/2 && i < end-i+start; i++ {
		nums[i], nums[end-i+start] = nums[end-i+start], nums[i]
	}
}
