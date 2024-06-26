package recursion

import "testing"

/*
TestMultiplication tests solution(s) with the following signature and problem description:

	func Multiply(a, b int) int

Given two integers, return their product using recursion and without using the
multiplication operator.
*/
func TestMultiplication(t *testing.T) {
	tests := []struct {
		a, b, product int
	}{
		{0, 0, 0},
		{0, 1, 0},
		{0, 1, 0},
		{1, 1, 1},
		{1, 2, 2},
		{2, 2, 4},
		{3, 3, 9},
		{3, -3, -9},
		{-3, -3, 9},
	}

	for i, test := range tests {
		if got := Multiply(test.a, test.b); got != test.product {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.product, got)
		}
	}
}
