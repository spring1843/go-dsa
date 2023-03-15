package array

// RotateKSteps rotates a given array k steps
func RotateKSteps(nums []int, k int) {
	ReverseInplace(nums, 0, len(nums)-k-1)
	ReverseInplace(nums, 0, len(nums)-1)
	ReverseInplace(nums, 0, k-1)
}
