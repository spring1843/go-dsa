package strings

import "testing"

/*
TestReadNumberInEnglish tests solution(s) with the following signature and problem description:

	func NumberInEnglish(num int) string

Given n a positive integer smaller than a Trillion, return the number in English words.

For example given 0, return "Zero".
For example given 34, return "Thirty Four".
For example given 10, return "Ten".
For example given 900000000001, return "Nine Hundred Billion One".
*/
func TestReadNumberInEnglish(t *testing.T) {
	tests := []struct {
		number  int
		english string
	}{
		{0, "Zero"},
		{1, "One"},
		{2, "Two"},
		{10, "Ten"},
		{34, "Thirty Four"},
		{123456789, "One Hundred Twenty Three Million Four Hundred Fifty Six Thousand Seven Hundred Eighty Nine"},
		{1000000000, "One Billion"},
		{100000000000, "One Hundred Billion"},
		{900000000001, "Nine Hundred Billion One"},
	}

	for i, test := range tests {
		got := NumberInEnglish(test.number)
		if got != test.english {
			t.Fatalf("Failed test case #%d. Want %q got %q", i, test.english, got)
		}
	}
}
