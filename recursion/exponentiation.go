package recursion

// PowerOf solves the problem in O(log(n)) time and O(1) space.
func PowerOf(x, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n%2 == 0 {
		return PowerOf(x*x, n/2)
	}
	return PowerOf(x, n-1) * x
}
