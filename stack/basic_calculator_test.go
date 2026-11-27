package stack

import "testing"

/*
TestBasicCalculator tests solution(s) with the following signature and problem description:

func BasicCalculator(input string) (float64, error)

Given an expression containing integers, parentheses and the four basic arithmetic operations
{*,/,+,-}, evaluate it. * and / have the same priority, which is higher than the priority of
+ and -, and + and - have the same priority. Operators with the same priority are evaluated
from left to right.

For example given 2+3*4 return 14 because * has a higher priority than +, so 2+(3*4) = 14.
Given 8/2*4 return 16 because / and * have the same priority and / is on the left, so
(8/2)*4 = 16. Given 8-3-2 return 3 because (8-3)-2 = 3.
*/
func TestBasicCalculator(t *testing.T) {
	tests := []struct {
		expression string
		expectErr  bool
		outcome    float64
	}{
		{"", true, -1},
		{"1++", true, -1},
		{"1+2", false, 3},
		{"1*(2+3)", false, 5},
		{"1+2+3", false, 6},
		{"1+(3-2)", false, 2},
		{"9/3", false, 3},
		{"3-9/3", false, 0},
		{"1*(2+3+(4*5))", false, 25},
		{"1*(2+3)+(4*5)", false, 25},
		{"5.10/2", false, 2.55},
		{"2*3+4", false, 10},
		{"2+3*4", false, 14},
		{"8-3-2", false, 3},
		{"8/4/2", false, 1},
		{"1*2+3+4*5", false, 25},
		{"8/2*4", false, 16},
	}

	for i, test := range tests {
		got, err := BasicCalculator(test.expression)
		if err != nil && !test.expectErr {
			t.Fatalf("Failed test case #%d. Did not expect an error. Error : %s", i, err)
		}

		if got != test.outcome {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.outcome, got)
		}
	}
}
