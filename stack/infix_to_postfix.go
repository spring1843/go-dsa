package stack

const (
	openParenthesis  = "("
	closeParenthesis = ")"
)

var operands = map[string]int{
	"*": 1,
	"/": 2,
	"+": 3,
	"-": 4,
}

type operators struct {
	stack []string
}

// InfixToPostfix solves the problem in O(n) time and O(n) space.
func InfixToPostfix(infix []string) []string {
	output := []string{}
	stack := new(operators)
	for _, element := range infix {
		if _, ok := operands[element]; !ok && element != openParenthesis && element != closeParenthesis {
			output = append(output, element)
			continue
		}

		if element == closeParenthesis {
			popped := stack.pop()
			for popped != openParenthesis {
				output = append(output, popped)
				popped = stack.pop()
			}
			continue
		}

		if len(stack.stack) > 0 && operands[element] > operands[stack.stack[len(stack.stack)-1]] {
			stack.push(element)
			continue
		}

		stack.push(element)
	}

	for len(stack.stack) > 0 {
		output = append(output, stack.pop())
	}

	return output
}

func (operators *operators) pop() string {
	tmp := operators.stack[len(operators.stack)-1]
	operators.stack = operators.stack[:len(operators.stack)-1]
	return tmp
}

func (operators *operators) push(operator string) {
	operators.stack = append(operators.stack, operator)
}
