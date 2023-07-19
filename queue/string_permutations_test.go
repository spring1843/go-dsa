package queue

import (
	"reflect"
	"testing"
)

/*
TestStringPermutations tests solution(s) with the following signature and problem description:

	func StringPermutations(input string) []string

Given a string like "abc", return all possible permutations like "abc,acb,bac,bca,cab,cba" using
a queue.
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
		if got := StringPermutations(test.input); !reflect.DeepEqual(got, test.permutations) {
			t.Fatalf("Failed test #%d, Failed getting list of maximums. Want %#v got %#v", i, test.permutations, got)
		}
	}
}
