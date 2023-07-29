package recursion

func Multiply(a, b int) int {
	if b == 0 {
		return 0
	}

	if b < 0 {
		return -Multiply(a, -b)
	}

	return a + Multiply(a, b-1)
}
