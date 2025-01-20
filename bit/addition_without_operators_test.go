package bit

import "testing"

/*
TestAdditionWithoutArithmeticOperators tests solution(s) with the following signature and problem description:

	func Add(x, y int) int

Add x by y, two integers without using any arithmetic operators such as {+,-,/,*,++,--,+=,…}.
*/
func TestAdditionWithoutArithmeticOperators(t *testing.T) {
	tests := []struct {
		a, b int
	}{
		{0, 1},
		{20, 2},
		{20, 4},
		{20, 3},
	}

	for i, test := range tests {
		got := AdditionWithoutArithmeticOperators(test.a, test.b)
		want := test.a + test.b
		if got != want {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, want, got)
		}
	}
}
