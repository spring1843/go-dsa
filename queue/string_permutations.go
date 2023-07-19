package queue

type state struct {
	Permutation string
	Remaining   string
}

func StringPermutations(input string) []string {
	output := []string{}
	queue := []state{{"", input}}

	for len(queue) > 0 {
		currentState := queue[0]
		queue = queue[1:]

		if len(currentState.Permutation) == len(input) {
			output = append(output, currentState.Permutation)
			continue
		}
		for i := 0; i < len(currentState.Remaining); i++ {
			nextPermutation := currentState.Permutation + string(currentState.Remaining[i])
			nextRemaining := currentState.Remaining[:i] + currentState.Remaining[i+1:]
			queue = append(queue, state{nextPermutation, nextRemaining})
		}
	}
	return output
}
