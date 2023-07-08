package queue

import (
	"testing"
)

/*
TestGenerateBinaryNumbers tests solution(s) with the following signature and problem description:

	func GenerateBinaryNumbers(n int) []string

Given a number n (n>=0) count from 0 to n in binary.
*/
func TestGenerateBinaryNumbers(t *testing.T) {
	tests := []struct {
		n        int
		lastLine string
	}{
		{0, "0"},
		{1, "1"},
		{2, "10"},
		{3, "11"},
		{4, "100"},
		{5, "101"},
		{6, "110"},
		{7, "111"},
		{8, "1000"},
		{9, "1001"},
		{10, "1010"},
	}

	for i, test := range tests {
		binaryNumbers := GenerateBinaryNumbers(test.n)
		got := binaryNumbers[len(binaryNumbers)-1]
		if got != test.lastLine {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.lastLine, got)
		}
	}
}
