package bit

// Max returns the maximum of two numbers without using if conditions.
func Max(x, y int) int {
	k := 1 ^ (((x - y) >> 31) & 0x1)
	q := 1 ^ k
	return x*k + y*q
}
