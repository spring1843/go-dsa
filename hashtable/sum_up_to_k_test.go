package hashtable

import (
	"slices"
	"testing"
)

/*
TestSumUpToK tests solution(s) with the following signature and problem description:

	func SumUpToK(numbers []int, k int) []int

Given a list, output the indices of the first two elements that sum up to K.
*/
func TestSumUpToK(t *testing.T) {
	tests := []struct {
		k                int
		numbers, indices []int
	}{
		{6, []int{3, 3}, []int{0, 1}},
		{3, []int{1, 2, 3, 4}, []int{0, 1}},
		{5, []int{1, 2, 3, 6}, []int{1, 2}},
		{7, []int{1, 2, 3, 4}, []int{2, 3}},
	}

	for i, test := range tests {
		got := SumUpToK(test.numbers, test.k)
		if !slices.Equal(got, test.indices) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.indices, got)
		}
	}
}
