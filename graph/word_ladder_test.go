package graph

import "testing"

/*
TestWordLadder tests solution(s) with the following signature and problem description:

	func WordLadder(start, end string, dic []string) int {

Returns the minimum number of transformations from start to end in a dictionary
where words are all equal in length, and a transformation can only happen if the difference
between two words is only in one letter. Zero is returned if no such transformations can occur.
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
