package graph

import "container/list"

// RemoveInvalidParentheses solves the problem in O(n^2) time and O(n) space.
func RemoveInvalidParentheses(input string) []string {
	bfs := func(input string) []string {
		output := make([]string, 0)
		if input == "" {
			return output
		}

		seen := make(map[string]struct{})

		queue := list.New()
		queue.PushBack(input)
		seen[input] = struct{}{}

		for queue.Len() != 0 {
			input = queue.Remove(queue.Front()).(string)

			if isValidParentheses(input) {
				if len(output) > 0 && len(output[0]) != len(input) {
					continue
				}
				output = append(output, input)
				continue
			}

			for i := range len(input) {
				if !(string(input[i]) == "(" || string(input[i]) == ")") {
					continue
				}

				tmp := input[0:i] + input[i+1:]
				if _, ok := seen[tmp]; !ok {
					queue.PushBack(tmp)
					seen[tmp] = struct{}{}
				}
			}
		}
		return output
	}
	return bfs(input)
}

func isValidParentheses(input string) bool {
	count := 0
	for i := range len(input) {
		if input[i] == '(' {
			count++
		} else if input[i] == ')' {
			count--
		}

		if count < 0 {
			return false
		}
	}
	return count == 0
}
