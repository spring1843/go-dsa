package stack

import "testing"

/*
TestEvaluatePostfix tests solution(s) with the following signature and problem description:

	func EvaluatePostfixExpression(expression []string) (float64, error)

Given a postfix expression calculate its the value.

The postfix expression is a list of strings where each string is either an operator like
arithmetic symbols or an operand that are numbers.

To evaluate the postfix expression, we start scanning the expression from left to right.
If the current element is an operator we apply the operand to the last two operands we read
we then remove the operator and replace the two operands with the results of the operation.

For example given 1 2 3 + *, return 5 because 1 * (2 + 3) = 5.
*/
func TestEvaluatePostfix(t *testing.T) {
	tests := []struct {
		postfix   []string
		expectErr bool
		outcome   float64
	}{
		{[]string{""}, true, -1},
		{[]string{"+"}, true, -1},
		{[]string{"A", "B", "+"}, true, -1},
		{[]string{"1", "B", "+"}, true, -1},
		{[]string{"1", "2", "+"}, false, 3},
		{[]string{"1", "2", "3", "+", "*"}, false, 5},
		{[]string{"1", "2", "3", "+", "+"}, false, 6},
		{[]string{"1", "3", "2", "-", "+"}, false, 2},
		{[]string{"9", "3", "/"}, false, 3},
		{[]string{"3", "9", "3", "/", "-"}, false, 0},
		{[]string{"1", "2", "3", "4", "5", "*", "+", "+", "*"}, false, 25},
		{[]string{"1", "2", "3", "+", "4", "5", "*", "+", "*"}, false, 25},
	}

	for i, test := range tests {
		got, err := EvaluatePostfixExpression(test.postfix)
		if err != nil && !test.expectErr {
			t.Fatalf("Did not expect an error. Error : %s", err)
		}

		if got != test.outcome {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.outcome, got)
		}
	}
}
