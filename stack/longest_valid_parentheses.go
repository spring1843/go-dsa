package stack

import "container/list"

// LongestValidParentheses solves the problem in O(n) time and O(n) space.
func LongestValidParentheses(input string) int {
	stack := list.New()

	longest := 0

	tmp := 0
	for _, char := range input {
		if char == '(' {
			stack.PushBack(char)
			continue
		}

		if stack.Len() == 0 {
			tmp = 0
			continue
		}

		stack.Remove(stack.Front())
		tmp += 2
		if tmp > longest {
			longest = tmp
		}
	}

	return longest
}
