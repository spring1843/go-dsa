package greedy

// MaxNumber solves the problem in O(n^2) time and O(n) space.
func MaxNumber(number1, number2 []int, n int) []int {
	output := []int{}
	if n > len(number1)+len(number2) {
		return output
	}

	for pickFromNumber1 := min(len(number1), n); pickFromNumber1 >= 0; pickFromNumber1-- {
		pickFromNumber2 := n - pickFromNumber1
		if pickFromNumber2 <= len(number2) {
			merged := merge(maximumNumberInSequence(number1, pickFromNumber1), maximumNumberInSequence(number2, pickFromNumber2))
			if len(output) == 0 || isGreater(merged, output) {
				output = merged
			}
		}
	}
	return output
}

func merge(a, b []int) []int {
	output := make([]int, len(a)+len(b))
	var i, j int
	for k := range len(a) + len(b) {
		if isGreater(a[i:], b[j:]) {
			output[k] = a[i]
			i++
			continue
		}
		output[k] = b[j]
		j++
	}
	return output
}

func isGreater(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return len(a) > len(b)
}

func maximumNumberInSequence(input []int, k int) []int {
	stack := make([]int, 0, k)
	for i, num := range input {
		for len(stack) > 0 && len(input)-i > k-len(stack) && num > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) < k {
			stack = append(stack, num)
		}
	}
	return stack
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
