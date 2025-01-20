package bit

// AdditionWithoutArithmeticOperators solves the problem in O(1) time and O(1) space.
func AdditionWithoutArithmeticOperators(x, y int) int {
	for y != 0 {
		sum := x ^ y
		carry := (x & y) << 1
		x = sum
		y = carry
	}
	return x
}
