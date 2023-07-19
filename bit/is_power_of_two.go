package bit

// IsPowerOfTwo solves the problem in O(1) time and O(1) space.
func IsPowerOfTwo(input int) bool {
	if input <= 0 {
		return false
	}
	return (input & (input - 1)) == 0
}
