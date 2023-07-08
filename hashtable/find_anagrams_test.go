package hashtable

import (
	"reflect"
	"testing"
)

/*
TestFindAnagrams tests solution(s) with the following signature and problem description:

	func FindAnagrams(words []string) [][]string

Given a dictionary, return lists of words that are anagrams of each other.
*/
func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		words    []string
		anagrams [][]string
	}{
		{[]string{"foo", "bar", "baz"}, [][]string{}},
		{[]string{"foo", "cat", "tac", "act"}, [][]string{{"cat", "tac", "act"}}},
		{[]string{"car", "cap", "arc", "tac"}, [][]string{{"car", "arc"}}},
	}

	for i, test := range tests {
		got := FindAnagrams(test.words)
		if !reflect.DeepEqual(got, test.anagrams) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.anagrams, got)
		}
	}
}
