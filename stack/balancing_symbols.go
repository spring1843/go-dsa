package stack

var (
	symbols = map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}
	closers = invertMap(symbols)

	// stack retains the opener symbols e.g. [({.
	stack []rune
)

// IsExpressionBalanced solves the problem in O(n) time and O(n) space.
func IsExpressionBalanced(s string) bool {
	stack = []rune{}
	if s == "" {
		return true
	}
	for _, c := range s {
		if _, ok := symbols[c]; ok {
			push(c)
			continue
		}

		if _, ok := closers[c]; !ok {
			continue
		}

		if len(stack) == 0 {
			return false // More closers than openers
		}

		lastOpener := pop()
		if symbols[lastOpener] != c {
			return false // opener did not match the last closer
		}
	}

	return len(stack) == 0
}

func push(a rune) {
	stack = append(stack, a)
}

func pop() rune {
	tmp := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return tmp
}

func invertMap(m map[rune]rune) map[rune]rune {
	inverted := make(map[rune]rune, len(m))
	for key, value := range m {
		inverted[value] = key
	}
	return inverted
}
