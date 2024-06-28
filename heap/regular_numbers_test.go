package heap

import (
	"slices"
	"testing"
)

/*
TestRegularNumbers tests solution(s) with the following signature and problem description:

	func RegularNumbers(n int) []int

Regular numbers are numbers whose only prime divisors are 2, 3, and 5.
For example 24 is a regular number because it can be factored into 2^3 * 3^1 * 5^0.
Given a positive integer n, write a function that returns the first n regular numbers.
*/
func TestRegularNumbers(t *testing.T) {
	tests := []struct {
		n              int
		regularNumbers []int
	}{
		{0, []int{}},
		{1, []int{1}},
		{2, []int{1, 2}},
		{3, []int{1, 2, 3}},
		{4, []int{1, 2, 3, 4}},
		{5, []int{1, 2, 3, 4, 5}},
		{10, []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12}},
		{15, []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, 16, 18, 20, 24}},
	}

	for i, test := range tests {
		if got := RegularNumbers(test.n); !slices.Equal(got, test.regularNumbers) {
			t.Fatalf("Failed test case #%d. Want %v got %d", i, test.regularNumbers, got)
		}
	}
}
