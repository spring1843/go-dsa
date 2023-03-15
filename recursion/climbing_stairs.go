package recursion

// ClimbingStairs returns in how many ways you can climb a stairs of n steps
// if we can only climb 1 or 2 steps only at a time
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
