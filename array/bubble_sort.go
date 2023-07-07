package array

// BubbleSort solves the problem in O(n^2) time and O(1) space.
func BubbleSort(input []int) {
	for i := range input {
		for j := range input {
			if input[i] < input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
}
