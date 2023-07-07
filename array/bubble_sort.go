package array

// BubbleSort implements the Bubble Sort algorithm.
func BubbleSort(input []int) {
	for i := range input {
		for j := range input {
			if input[i] < input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
}
