package bit

// Add adds to numbers without using any arithmetic operators
func Add(x, y int) int {
	for y != 0 {
		sum := x ^ y
		carry := (x & y) << 1
		x = sum
		y = carry
	}
	return x
}
