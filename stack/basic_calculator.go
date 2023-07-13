package stack

import (
	"errors"
)

type evaluation struct {
	stack []float64
}

// ErrImbalancedParentheses occurs when the expression has imbalanced parentheses.
var ErrImbalancedParentheses = errors.New("expression has imbalanced parentheses")

// BasicCalculator solves the problem in O(n) time and O(n) space.
func BasicCalculator(input string) (float64, error) {
	if IsExpressionBalanced(input) {
		return -1, ErrImbalancedParentheses
	}

	return EvaluatePostfixExpression(InfixToPostfix(toInfix(input)))
}

func toInfix(input string) []string {
	var (
		curNum string
		infix  []string
	)

	for _, r := range input {
		char := string(r)
		if _, ok := operands[char]; ok {
			if len(curNum) > 0 {
				infix = append(infix, curNum)
				curNum = ""
			}
			infix = append(infix, char)
			continue
		}

		if char == openParenthesis || char == closeParenthesis {
			if len(curNum) > 0 {
				infix = append(infix, curNum)
				curNum = ""
			}
			infix = append(infix, char)
			continue
		}
		curNum += char
	}

	if len(curNum) > 0 {
		infix = append(infix, curNum)
	}

	return infix
}
