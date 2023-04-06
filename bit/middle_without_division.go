package bit

// MiddleWithoutDivision finds the middle integer between two integers.
func MiddleWithoutDivision(min, max int) int {
	return (min + max) >> 1
}
