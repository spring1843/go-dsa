package bit

import "testing"

/*
TestIsPowerOfTwo tests solution(s) with the following signature and problem description:

	func IsPowerOfTwo(input int) bool

Using bit manipulation, return true if a given number like 2 and false otherwise.
*/
func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		input        int
		isPowerOfTwo bool
	}{
		{0, false},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{20, false},
		{256, true},
	}

	for i, test := range tests {
		if got := IsPowerOfTwo(test.input); got != test.isPowerOfTwo {
			t.Fatalf("Failed test case #%d. Want %t got %t", i, test.isPowerOfTwo, got)
		}
	}
}
