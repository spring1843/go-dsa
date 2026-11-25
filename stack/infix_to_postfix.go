package stack

const (
	openParenthesis  = "("
	closeParenthesis = ")"
)

// operators maps each operator to its precedence: * and / bind tighter than
// + and -, and operators of equal precedence are applied left to right.
var operators = map[string]int{
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
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

// pushOperator pops the operators that bind at least as tightly as operator
// onto output, then pushes operator. A higher precedence applies first, and an
// equal one is to the left, so it applies first too. An open parenthesis has
// no precedence and stops the popping.
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
