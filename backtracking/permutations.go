package backtracking

// Permutations solves the problem in O(n!) time and O(n) space.
func Permutations(input []int) [][]int {
	permutations := make([][]int, 0)
	permutationsRecursive(input, 0, &permutations)
	return permutations
}

func permutationsRecursive(input []int, start int, permutations *[][]int) {
	end := len(input) - 1
	if start == end {
		*permutations = append(*permutations, append([]int{}, input...))
		return
	}
	for i := start; i <= end; i++ {
		(input)[i], (input)[start] = (input)[start], (input)[i]
		permutationsRecursive(input, start+1, permutations)
		(input)[i], (input)[start] = (input)[start], (input)[i]
	}
}
