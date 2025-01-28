package strings

import "testing"

/*
TestFindDuplicate tests solution(s) with the following signature and problem description:

	func LookAndTell(depth int) []string

Given a positive integer n, return the output of the Look and Tell algorithm until the nth depth.

The Look and Tell algorithm starts by outputting 1 at first level, then at each subsequent level
it reads the previous line by counting the number of times a digit is repeated and then writes
the count and the digit.

For example given 4, return:

1
11
21
1211

Which reads:
one
one one
two ones
one two one one.

The third output (two ones) reads the level before it which is 11. Two ones means repeat
1 two times i.e. 11.
*/
func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		depth    int
		lastLine string
	}{
		{0, "-1"},
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
		{6, "312211"},
		{7, "13112221"},
		{8, "1113213211"},
		{9, "31131211131221"},
		{10, "13211311123113112211"},
	}

	for i, test := range tests {
		tell := LookAndTell(test.depth)
		got := tell[len(tell)-1]
		if got != test.lastLine {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.lastLine, got)
		}
	}
}
