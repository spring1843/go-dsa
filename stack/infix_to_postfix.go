package stack

const (
	openParenthesis  = "("
	closeParenthesis = ")"
)

var operators = map[string]int{
	"*": 1,
	"/": 2,
	"+": 3,
	"-": 4,
}

type operatorStack struct {
	stack []string
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

		if len(stack.stack) > 0 && operators[element] > operators[stack.stack[len(stack.stack)-1]] {
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

func (s *operatorStack) pop() string {
	tmp := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return tmp
}

func (s *operatorStack) push(operator string) {
	s.stack = append(s.stack, operator)
}
