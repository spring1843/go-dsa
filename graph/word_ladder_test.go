package graph

import "testing"

/*
TestWordLadder tests solution(s) with the following signature and problem description:

	func WordLadder(start, end string, dic []string) int

Given a start word like `pop` and an end word like `car`, a dictionary of same length words
like  `{"top","cop","cap","car"}` return the minimum number of transformations like 4 to get
from start to end where each transformation between two words can happen when they are
different by only one letter.
*/
func TestWordLadder(t *testing.T) {
	tests := []struct {
		start, end         string
		dic                []string
		minTransformations int
	}{
		{"foo", "bar", []string{}, 0},
		{"foo", "bar", []string{"baz"}, 0},
		{"a", "c", []string{"a", "b", "c"}, 1},
		{"pop", "cop", []string{"tap", "top", "cop"}, 1},
		{"car", "tor", []string{"cap", "tap", "top", "tar", "tor"}, 3},
		{"pop", "car", []string{"top", "cop", "cap", "car"}, 4},
		{"pot", "cop", []string{"rot", "rat", "fat", "fax", "tax", "tap", "top", "cop"}, 8},
	}

	for i, test := range tests {
		if got := WordLadder(test.start, test.end, test.dic); got != test.minTransformations {
			t.Fatalf("Failed test case #%d. Want %#v got %d", i, test.minTransformations, got)
		}
	}
}
