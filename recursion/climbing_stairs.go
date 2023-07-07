package recursion

// ClimbingStairs solves the problem in O(2^n) time and O(n) space.
func ClimbingStairs(n int) int {
	return climbStairsRecursive(0, n)
}

func climbStairsRecursive(stepsTakenSoFar, n int) int {
	if stepsTakenSoFar > n {
		return 0
	}
	if stepsTakenSoFar == n {
		return 1
	}
	return climbStairsRecursive(stepsTakenSoFar+1, n) + climbStairsRecursive(stepsTakenSoFar+2, n)
}
