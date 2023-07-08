package bit

import "testing"

/*
TestDivision tests solution(s) with the following signature and problem description:

	func Divide(x, y int) int

Divide x by y, two integers without using the built-in `/` or `*` operators.
*/
func TestDivision(t *testing.T) {
	tests := []struct {
		a, b int
	}{
		{0, 1},
		{20, 2},
		{20, 4},
		{20, 3},
	}

	for i, test := range tests {
		got := Divide(test.a, test.b)
		want := test.a / test.b
		if got != want {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, want, got)
		}
	}
}
