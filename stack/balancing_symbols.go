package stack

var (
	symbols = map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}

	// stack retains the opener symbols e.g. [({
	stack []rune
)

// IsExpressionBalanced returns true if a given expression contains symbols like [](){} is balanced or not
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

		if len(stack) == 0 {
			return false // More closers than openers
		}

		lastOpener := pop()
		if symbols[lastOpener] != c {
			return false // opener did not match the last closer
		}
	}

	if len(stack) != 0 {
		return false // Some openers remain unclosed
	}

	return true
}

func push(a rune) {
	stack = append(stack, a)
}

func pop() rune {
	tmp := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return tmp
}
