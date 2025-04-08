package bit

import "testing"

/*
TestMax tests solution(s) with the following signature and problem description:

	func Max(x, y int) int

Given two integers, return the larger of the two without using any comparison
operations like {if, switch,…}.

For example given 20 and 2 return 20.
*/
func TestMax(t *testing.T) {
	tests := []struct {
		a, b, max int
	}{
		{0, 1, 1},
		{20, 2, 20},
		{2, 20, 20},
		{2, -2, 2},
		{-3, -2, -2},
	}

	for i, test := range tests {
		got := Max(test.a, test.b)
		if got != test.max {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.max, got)
		}
	}
}
