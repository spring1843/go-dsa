package greedy

// ActivitySelector selects activities that don't overlap.
func ActivitySelector(start, finish []int) []int {
	if len(start) == 0 {
		return []int{}
	}
	activities := []int{1}
	k := 0
	for m := 1; m < len(start); m++ {
		if start[m] >= finish[k] {
			activities = append(activities, m+1)
			k = m
		}
	}
	return activities
}
