package bit

// Max solves the problem in O(1) time and O(1) space.
func Max(x, y int) int {
	k := 1 ^ (((x - y) >> 31) & 0x1)
	q := 1 ^ k
	return x*k + y*q
}
