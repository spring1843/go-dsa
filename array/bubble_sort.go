package array

// BubbleSort solves the problem in O(n^2) time and O(1) space.
func BubbleSort(input []int) {
	n := len(input)
	for i := 0; i < n-1; i++ {
		// flag to check if there were any swaps made
		swapped := false
		// loop can run until n-i-1 since last i elements are already sorted in ith pass
		for j := 0; j < n-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				swapped = true
			}
		}
		if !swapped {
			// if no swaps were made, array is already sorted
			break
		}
	}
}
