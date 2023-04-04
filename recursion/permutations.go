package recursion

// Permutations returns all possible permutations of the given integers.
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
