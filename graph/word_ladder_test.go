package graph

import "testing"

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
