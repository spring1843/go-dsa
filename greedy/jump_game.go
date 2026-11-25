package greedy

// JumpGame solves the problem in O(n) time and O(1) space.
func JumpGame(jumps []int) bool {
	if len(jumps) == 0 {
		return true
	}
	maxReachableDistance := 0
	for i, jump := range jumps {
		if i > maxReachableDistance {
			break
		}
		maxReachableDistance = max(maxReachableDistance, i+jump)
	}
	return maxReachableDistance >= len(jumps)-1
}
