package hashtable

// MissingNumber solves the problem in O(n) time and O(n) space.
func MissingNumber(numbers []int) int {
	if len(numbers) < 1 {
		return -1
	}
	maxNumber := numbers[0]
	minNumber := numbers[0]
	seen := make(map[int]struct{})
	for _, number := range numbers {
		maxNumber = max(maxNumber, number)
		minNumber = min(minNumber, number)
		seen[number] = struct{}{}
	}

	for i := minNumber; i <= maxNumber; i++ {
		if _, ok := seen[i]; !ok {
			return i
		}
	}
	return -1
}
