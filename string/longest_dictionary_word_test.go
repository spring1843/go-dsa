package string

import (
	"testing"
)

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
