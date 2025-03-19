package recursion

import "testing"

/*
TestExpressionOperators tests solution(s) with the following signature and problem description:

	func ExpressionOperators(list []int, target int) string

Given a slice of numbers representing operands in an equation, and a target integer representing
the result of the equation, return a string representing operators that can be inserted between
the operands to form the equation and yield the target result.
Only + and - operators are allowed and the are assumed to have the same priority

For example given {1,5,3} and 3 return {+,-} because 1+5-3 = 3.
*/
func TestExpressionOperators(t *testing.T) {
	tests := []struct {
		operands  []int
		target    int
		operators string
	}{
		{[]int{}, 1, ""},
		{[]int{1}, 1, ""},
		{[]int{1, 2}, 3, "+"},
		{[]int{1, 2, 3}, 0, "+-"},
		{[]int{1, 5, 3}, 3, "+-"},
		{[]int{1, 2, 3, 4, 5, 6}, 5, "+-+-+"},
	}

	for i, test := range tests {
		if got := ExpressionOperators(test.operands, test.target); got != test.operators {
			t.Fatalf("Failed test case #%d. Want %s got %s", i, test.operators, got)
		}
	}
}
