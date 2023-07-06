package backtracking

// Permutations intakes a list of numbers and returns all possible permutations of their orders
// For example for {1,2} it would return {1,2}, {2,1}.
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
