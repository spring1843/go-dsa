package bit

// MiddleWithoutDivision solves the problem in O(1) time and O(1) space.
func MiddleWithoutDivision(low, high int) int {
	return (low + high) >> 1
}
