package greedy

// JumpGame returns true if we can reach the last index of the given a slice, where each value at
// each position represents the maximum allowed jump length from that position.
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
