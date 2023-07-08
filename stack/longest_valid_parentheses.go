package stack

var intStack []int

// LongestValidParentheses solves the problem in O(n) time and O(n) space.
func LongestValidParentheses(s string) int {
	intStack = make([]int, 0)
	intStackPush(-1)

	longest := 0
	for i, char := range s {
		if char == '(' {
			intStackPush(i)
			continue
		}
		intStackPop()
		if len(intStack) == 0 {
			intStackPush(i)
			continue
		}
		longest = max(longest, i-intStackPeek())
	}
	return longest
}

func intStackPush(a int) {
	intStack = append(intStack, a)
}

func intStackPop() int {
	tmp := intStack[len(intStack)-1]
	intStack = intStack[:len(intStack)-1]
	return tmp
}

func intStackPeek() int {
	return intStack[len(intStack)-1]
}
