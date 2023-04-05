package stack

import (
	"fmt"
	"strconv"
)

// EvaluatePostfixExpression evaluates an expression into a float64
func EvaluatePostfixExpression(expression []string) (float64, error) {
	stack := new(evaluation)
	for _, element := range expression {
		if _, ok := operands[element]; !ok {
			if err := stack.pushString(element); err != nil {
				return -1, err
			}
			continue
		}

		op2, err := stack.pop()
		if err != nil {
			return -1, err
		}

		op1, err := stack.pop()
		if err != nil {
			return -1, err
		}

		var result float64
		switch element {
		case "+":
			result = op1 + op2
		case "-":
			result = op1 - op2
		case "*":
			result = op1 * op2
		case "/":
			result = op1 / op2
		}
		stack.push(result)
	}
	return stack.stack[len(stack.stack)-1], nil
}

func (evaluation *evaluation) pop() (float64, error) {
	if len(evaluation.stack) == 0 {
		return -1, ErrEmptyStack
	}
	tmp := evaluation.stack[len(evaluation.stack)-1]
	evaluation.stack = evaluation.stack[:len(evaluation.stack)-1]
	return tmp, nil
}

func (evaluation *evaluation) pushString(operand string) error {
	f, err := strconv.ParseFloat(operand, 64)
	if err != nil {
		return fmt.Errorf("failed converting from float to string, %s", err)
	}
	evaluation.push(f)
	return nil
}

func (evaluation *evaluation) push(f float64) {
	evaluation.stack = append(evaluation.stack, f)
}
