package stack

const (
	openParenthesis  = "("
	closeParenthesis = ")"

	higherPrecedence = 2
	lowerPrecedence  = 1
)

type operatorStack struct {
	stack []string
}

var operators = map[string]int{
	"*": higherPrecedence,
	"/": higherPrecedence,
	"+": lowerPrecedence,
	"-": lowerPrecedence,
}

// InfixToPostfix solves the problem in O(n) time and O(n) space.
func InfixToPostfix(infix []string) []string {
	output := []string{}
	stack := new(operatorStack)
	for _, element := range infix {
		if _, ok := operators[element]; !ok && element != openParenthesis && element != closeParenthesis {
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

		if element == openParenthesis {
			stack.push(element)
			continue
		}

		output = stack.pushOperator(element, output)
	}

	for len(stack.stack) > 0 {
		output = append(output, stack.pop())
	}

	return output
}

func (s *operatorStack) pushOperator(operator string, output []string) []string {
	for len(s.stack) > 0 && operators[s.stack[len(s.stack)-1]] >= operators[operator] {
		output = append(output, s.pop())
	}
	s.push(operator)
	return output
}

func (s *operatorStack) pop() string {
	tmp := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return tmp
}

func (s *operatorStack) push(operator string) {
	s.stack = append(s.stack, operator)
}
