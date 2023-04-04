package string

import (
	"testing"
)

func TestReadNumberInEnglish(t *testing.T) {
	tests := []struct {
		number  int
		english string
	}{
		{0, "Zero"},
		{1, "One"},
		{2, "Two"},
		{34, "Thirty Four"},
		{123456789, "One Hundred Twenty Three Million Four Hundred Fifty Six Thousand Seven Hundred Eighty Nine"},
		{1000000000, "One Billion"},
		{100000000000, "One Hundred Billion"},
	}

	for i, test := range tests {
		got := NumberInEnglish(test.number)
		if got != test.english {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.english, got)
		}
	}
}
