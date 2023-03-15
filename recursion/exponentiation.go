package recursion

// PowerOf raises x to the power of n
func PowerOf(x, n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return x
	}
	if n%2 == 0 {
		return PowerOf(x*x, n/2)
	}
	return PowerOf(x, n-1) * x
}
