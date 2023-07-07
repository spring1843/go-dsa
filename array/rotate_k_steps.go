package array

// RotateKSteps solves problem in O(n) time and O(1) space.
func RotateKSteps(list []int, k int) {
	ReverseInPlace(list, 0, len(list)-k-1)
	ReverseInPlace(list, 0, len(list)-1)
	ReverseInPlace(list, 0, k-1)
}
