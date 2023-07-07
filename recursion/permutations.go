package recursion

// Permutations solves the problem in O(n!) time and O(n) space.
func Permutations(numbers []int) [][]int {
	var permutations [][]int
	permutationsRecursive(numbers, 0, &permutations)
	return permutations
}

func permutationsRecursive(numbers []int, start int, permutations *[][]int) {
	if start >= len(numbers) {
		*permutations = append(*permutations, append([]int{}, numbers...))
		return
	}
	for i := start; i < len(numbers); i++ {
		numbers[start], numbers[i] = numbers[i], numbers[start]
		permutationsRecursive(numbers, start+1, permutations)
		numbers[start], numbers[i] = numbers[i], numbers[start]
	}
}
