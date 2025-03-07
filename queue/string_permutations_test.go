package queue

import (
	"slices"
	"testing"
)

/*
TestStringPermutations tests solution(s) with the following signature and problem description:

	func StringPermutations(input string) []string

Given a string return all possible permutations that can be made by rearranging the letters in the string
using a queue.

For example given "abc" return "abc,acb,bac,bca,cab,cba"
*/
func TestStringPermutations(t *testing.T) {
	tests := []struct {
		input        string
		permutations []string
	}{
		{"", []string{""}},
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"ba", []string{"ba", "ab"}},
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"123", []string{"123", "132", "213", "231", "312", "321"}},
	}

	for i, test := range tests {
		if got := StringPermutations(test.input); !slices.Equal(got, test.permutations) {
			t.Fatalf("Failed test #%d, Failed getting list of maximums. Want %#v got %#v", i, test.permutations, got)
		}
	}
}
