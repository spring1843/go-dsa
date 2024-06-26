package stack

import "testing"

/*
TestBasicCalculator tests solution(s) with the following signature and problem description:
func BasicCalculator(input string) (float64, error)

Given an expression containing integers, parentheses and the four basic arithmetic operations
like 1*2+3+4*5 calculate the expression into a numerical value like 25.
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
