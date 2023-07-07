package bit

import "testing"

/*
TestDivision tests solution(s) with the following signature and problem description:

	func Add(x, y int) int

Adds to numbers without using any arithmetic operators.
*/
func TestAdd(t *testing.T) {
	tests := []struct {
		a, b int
	}{
		{0, 1},
		{20, 2},
		{20, 4},
		{20, 3},
	}

	for i, test := range tests {
		got := Add(test.a, test.b)
		want := test.a + test.b
		if got != want {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, want, got)
		}
	}
}
