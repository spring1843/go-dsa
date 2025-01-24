package strings

import "testing"

/*
TestLongestDictionaryWordContainingKey tests solution(s) with the following signature and problem description:

	func LongestDictionaryWordContainingKey(key string, dic []string) string

Given a key as string, and a slice of strings containing a dictionary of words, return the longest
word that contains all letters of the key.

For example given "cat" and {"rectify", "race", "archeology", "racoon"}, it should return "archeology",
because "archeology" is the longest word in the given set that contains "c","a",and "t".
*/
func TestLongestDictionaryWordContainingKey(t *testing.T) {
	tests := []struct {
		input                           string
		dictionary                      []string
		longestWordContainingCharacters string
	}{
		{"a", []string{}, ""},
		{"", []string{"abc"}, "abc"},
		{"a", []string{"a", "b", "c"}, "a"},
		{"a", []string{"a", "ba", "c", "cc"}, "ba"},
		{"a", []string{"a", "baa", "c"}, "baa"},
		{"abc", []string{"abc", "abcdefghijklmn", "abcdefghijkl", "abcdef", "abcijkl"}, "abcdefghijklmn"},
		{"car", []string{"rectify", "race", "archeology", "racoon"}, "archeology"},
	}

	for i, test := range tests {
		got := LongestDictionaryWordContainingKey(test.input, test.dictionary)
		if got != test.longestWordContainingCharacters {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.longestWordContainingCharacters, got)
		}
	}
}
