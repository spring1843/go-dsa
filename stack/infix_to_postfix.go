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

// InfixToPostfix converts an infix expression to a postfix one supporting the 4 basic arithmetic operations and parentheses
func InfixToPostfix(infix []string) ([]string, error) {
	output := []string{}
	stack := new(operators)
	for _, element := range infix {
		if _, ok := operands[element]; !ok && element != openParenthesis && element != closeParenthesis {
			output = append(output, element)
			continue
		}

		if element == closeParenthesis {
			popped, _ := stack.pop()
			for popped != openParenthesis {
				output = append(output, popped)
				popped, _ = stack.pop()
			}
			continue
		}

		if len(stack.stack) > 0 && operands[element] > operands[stack.stack[len(stack.stack)-1]] {
			stack.push(element)
			continue
		}

		for len(stack.stack) > 0 && operands[element] > operands[stack.stack[len(stack.stack)-1]] {
			popped, _ := stack.pop()
			output = append(output, popped)
		}

		stack.push(element)
	}

	for len(stack.stack) > 0 {
		operand, err := stack.pop()
		if err != nil {
			return nil, err
		}
		output = append(output, operand)
	}

	return output, nil
}

func (operators *operators) pop() (string, error) {
	if len(operators.stack) == 0 {
		return "", ErrEmptyStack
	}
	tmp := operators.stack[len(operators.stack)-1]
	operators.stack = operators.stack[:len(operators.stack)-1]
	return tmp, nil
}

func (operators *operators) push(operator string) {
	operators.stack = append(operators.stack, operator)
}
