package strings

import (
	"testing"
)

/*
TestFindDuplicate tests solution(s) with the following signature and problem description:

	func LookAndTell(depth int) []string

Outputs an array such that each element reads the previous element the first element being "1".
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
