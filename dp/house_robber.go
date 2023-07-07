package dp

// MaxHouseRobber solves the problem in O(n) time and O(1) space.
func MaxHouseRobber(wealth []int) int {
	if len(wealth) == 0 {
		return 0
	}

	var n1, n2, tmp int
	n2 = wealth[0]

	for i := 1; i < len(wealth); i++ {
		tmp = max(n1+wealth[i], n2)
		n1 = n2
		n2 = tmp
	}
	return n2
}
