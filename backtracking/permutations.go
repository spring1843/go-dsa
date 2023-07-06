package backtracking

// Permutations intakes a list of numbers and returns all possible permutations of their orders
// For example for {1,2} it would return {1,2}, {2,1}.
func Permutations(num []int) [][]int {
	permutations := make([][]int, 0)
	permutationsRecursive(&num, 0, len(num)-1, &permutations)
	return permutations
}

func permutationsRecursive(nums *[]int, start, end int, retSlice *[][]int) {
	if start == end {
		copyNums := []int{}
		copyNums = append(copyNums, *nums...)
		*retSlice = append(*retSlice, copyNums)
		return
	}
	for i := start; i <= end; i++ {
		(*nums)[i], (*nums)[start] = (*nums)[start], (*nums)[i]
		permutationsRecursive(nums, start+1, end, retSlice)
		(*nums)[i], (*nums)[start] = (*nums)[start], (*nums)[i]
	}
}
