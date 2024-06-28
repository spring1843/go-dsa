package queue

import "container/list"

type state struct {
	permutation string
	remaining   string
}

// StringPermutations solves the problem in O(n!) time and O(n!) space.
func StringPermutations(input string) []string {
	output := []string{}
	queue := list.New()
	queue.PushBack(state{"", input})

	for queue.Len() > 0 {
		currentState := queue.Remove(queue.Front()).(state)

		if len(currentState.permutation) == len(input) {
			output = append(output, currentState.permutation)
			continue
		}
		for i := range len(currentState.remaining) {
			nextPermutation := currentState.permutation + string(currentState.remaining[i])
			nextRemaining := currentState.remaining[:i] + currentState.remaining[i+1:]
			queue.PushBack(state{nextPermutation, nextRemaining})
		}
	}
	return output
}
