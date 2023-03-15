package recursion

// ReverseDigits returns an integer with digits in the reverse order of the input integer
func ReverseDigits(n int) int {
	return reverseRecursiveLy(n, 0, digitsCount(n))
}

func reverseRecursiveLy(n, i, l int) int {
	if n < 10 {
		return n
	}
	return (n%10)*PowerOf(10, l-i-1) + reverseRecursiveLy(n/10, i+1, l)
}

func digitsCount(n int) int {
	if n < 10 {
		return 1
	}
	return 1 + digitsCount(n/10)
}
