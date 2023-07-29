package recursion

// Multiply solves the problem in O(b) time and O(1) space.
func Multiply(a, b int) int {
	if b == 0 {
		return 0
	}

	if b < 0 {
		return -Multiply(a, -b)
	}

	return a + Multiply(a, b-1)
}
