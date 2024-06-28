package backtracking

import (
	"slices"
	"sort"
	"testing"
)

/*
TestPhoneLetterCombinations tests solution(s) with the following signature and problem description:

	func PhoneLetterCombinations(digits string) []string

Intakes a string of digits from 2 to 9 inclusive and returns all possible combinations of letters
that could be generated from those. For example for input 2 it should return abc.
*/
func TestPhoneLetterCombinations(t *testing.T) {
	tests := []struct {
		digits       string
		combinations []string
	}{
		{"", []string{}},
		{"1", []string{}},
		{"2", []string{"a", "b", "c"}},
		{"9", []string{"w", "x", "y", "z"}},
		{"12", []string{}},
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"232", []string{"ada", "adb", "adc", "aea", "aeb", "aec", "afa", "afb", "afc", "bda", "bdb", "bdc", "bea", "beb", "bec", "bfa", "bfb", "bfc", "cda", "cdb", "cdc", "cea", "ceb", "cec", "cfa", "cfb", "cfc"}},
	}

	for i, test := range tests {
		got := PhoneLetterCombinations(test.digits)
		if len(got) > 0 {
			sort.Strings(got)
		}
		if !slices.Equal(test.combinations, got) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.combinations, got)
		}
	}
}
