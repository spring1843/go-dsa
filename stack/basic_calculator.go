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
	if !hasBalancedParentheses(input) {
		return -1, ErrImbalancedParentheses
	}

	return EvaluatePostfixExpression(InfixToPostfix(toInfix(input)))
}

// hasBalancedParentheses reports whether every closing parenthesis in input
// closes an earlier open one and every open one is closed. Digits and
// operators are ignored.
func hasBalancedParentheses(input string) bool {
	depth := 0
	for _, r := range input {
		switch string(r) {
		case openParenthesis:
			depth++
		case closeParenthesis:
			depth--
			if depth < 0 {
				return false
			}
		}
	}
	return depth == 0
}

func toInfix(input string) []string {
	var (
		curNum string
		infix  []string
	)

	for _, r := range input {
		char := string(r)
		if _, ok := operators[char]; ok {
			if curNum != "" {
				infix = append(infix, curNum)
				curNum = ""
			}
			infix = append(infix, char)
			continue
		}

		if char == openParenthesis || char == closeParenthesis {
			if curNum != "" {
				infix = append(infix, curNum)
				curNum = ""
			}
			infix = append(infix, char)
			continue
		}
		curNum += char
	}

	if curNum != "" {
		infix = append(infix, curNum)
	}

	return infix
}
