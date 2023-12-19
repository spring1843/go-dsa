package recursion

// Multiply solves the problem in O(b) time and O(1) space.
func Multiply(a, b int) int {
	result := 0
	negative := b < 0
	if negative {
		b = -b
	}

	for b > 0 {
		result += a
		b--
	}

	if negative {
		return -result
	}
	return result
}
