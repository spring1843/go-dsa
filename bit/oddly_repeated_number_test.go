package bit

import (
	"testing"
)

/*
TestOddlyRepeatedNumber tests solution(s) with the following signature and problem description:

	func OddlyRepeatedNumber(list []int) int

Given an array of integers that are all repeated an even number of times except one,
find the oddly repeated element.
*/
func TestOddlyRepeatedNumber(t *testing.T) {
	tests := []struct {
		list          []int
		oddlyRepeated int
	}{
		{[]int{}, -1},
		{[]int{1, 2, 2}, 1},
		{[]int{1, 2, 1, 2, 3}, 3},
		{[]int{1, 2, 2, 3, 3}, 1},
	}

	for i, test := range tests {
		got := OddlyRepeatedNumber(test.list)
		if got != test.oddlyRepeated {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.oddlyRepeated, got)
		}
	}
}
