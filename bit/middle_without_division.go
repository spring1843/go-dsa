package bit

// MiddleWithoutDivision solves the problem in O(1) time and O(1) space.
func MiddleWithoutDivision(min, max int) int {
	return (min + max) >> 1
}
