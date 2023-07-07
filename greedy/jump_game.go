package greedy

// JumpGame solves the problem in O(n) time and O(1) space.
func JumpGame(jumps []int) bool {
	if len(jumps) == 0 {
		return true
	}
	maxReachableDistance := 0
	for i := 0; i < len(jumps); i++ {
		if i > maxReachableDistance {
			break
		}
		maxReachableDistance = max(maxReachableDistance, i+jumps[i])
	}
	return maxReachableDistance >= len(jumps)-1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
