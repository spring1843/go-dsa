package dnc

import (
	"math"
	"testing"
)

/*
TestSquareRoot tests solution(s) with the following signature and problem description:

	func SquareRoot(number, precision int) float64

Given a number and precision, return the square root of the number using the binary
search algorithm.
*/
func TestSquareRoot(t *testing.T) {
	tests := []struct {
		number    int
		precision int
		solution  float64
	}{
		{0, 0, 1},
		{1, 0, 1},
		{1, 1, 1},
		{4, 1, 2},
		{4, 2, 1.9999999999999998},
		{4, 3, 2},
		{5, 3, 2},
		{9, 3, 3},
	}

	floatAlmostEquals := func(a, b float64) bool {
		return math.Abs(a-b) <= 1e9 // Equality threshold
	}

	for i, test := range tests {
		if got := SquareRoot(test.number, test.precision); !floatAlmostEquals(got, test.solution) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.solution, got)
		}
	}
}
