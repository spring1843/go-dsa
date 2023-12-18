package stack

var (
	symbols = map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}
)

// IsExpressionBalanced solves the problem in O(n) time and O(n) space.
func IsExpressionBalanced(s string) bool {
	// stack retains the opener symbols e.g. [({.
	stack := []rune{} // Localized stack

    for _, c := range s {
		if _, ok := symbols[c]; ok {
            stack = push(stack, c)
            continue
        }

        if len(stack) == 0 {
            return false // More closers than openers
        }

        var lastOpener rune
        stack, lastOpener = pop(stack)
        if symbols[lastOpener] != c {
            return false // Opener did not match the last closer
        }
    }

    return len(stack) == 0
}

func push(stack []rune, a rune) []rune {
	return append(stack, a)
}

func pop(stack []rune) ([]rune, rune) {
    tmp := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    return stack, tmp
}
