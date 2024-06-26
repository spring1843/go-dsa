package dp

import "testing"

/*
TestWordDistance tests solution(s) with the following signature and problem description:

	func WordDistance(input1, input2 string) int

Given a string like abc, and another string like abcde return how many character
modifications (insert, delete, edit) have to be done on the first string to become
identical to the second string. In this case, the answer is 2.
*/
func TestWordDistance(t *testing.T) {
	tests := []struct {
		input1   string
		input2   string
		distance int
	}{
		{"", "", 0},
		{"", "abcde", 5},
		{"a", "a", 0},
		{"a", "b", 1},
		{"ab", "ac", 1},
		{"ab", "dc", 2},
		{"ab", "dcd", 3},
		{"abc", "abcde", 2},
		{"abcdef", "abcde", 1},
		{"gabcdef", "abcde", 2},
	}

	for i, test := range tests {
		if got := WordDistance(test.input1, test.input2); got != test.distance {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.distance, got)
		}
	}
}
