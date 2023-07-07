package recursion

// ReverseDigits solves the problem in O(n) time and O(1) space.
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
