package array

// BubbleSort solves the problem in O(n^2) time and O(1) space.
func BubbleSort(input []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(input); i++ {
			if input[i] < input[i-1] {
				input[i], input[i-1] = input[i-1], input[i]
				swapped = true
			}
		}
	}
}
