package queue

import (
	"slices"
	"testing"
)

/*
TestMaxOfKLengthSubArrays tests solution(s) with the following signature and problem description:

	func MaxOfKLengthSubArrays(numbers []int, k int) ([]int, error)

Given a slice of numbers and an integer k return a slice containing the maximum in each k-sized
sub-array (sub-slice) of the input.

For example given {1,2,3,4,5} and k=2, return {2,3,4,5} because:

* Sub-arrays of the input with length 2 are {{1,2},{2,3},{3,4},{4,5}}
* The maximum in each of the sub-arrays is {2,3,4,5}
*/
func TestMaxOfKLengthSubArrays(t *testing.T) {
	tests := []struct {
		k        int
		list     []int
		maximums []int
	}{
		{3, []int{5, 3, 5, 6, 7, 8}, []int{5, 6, 7, 8}},
		{2, []int{5, 3, 5, 6, 7, 8}, []int{5, 5, 6, 7, 8}},
		{3, []int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
	}

	for i, test := range tests {
		got, err := MaxOfKLengthSubArrays(test.list, test.k)
		if err != nil {
			t.Fatalf("Unexpected error occurred. Error : %s", err)
		}
		if !slices.Equal(got, test.maximums) {
			t.Fatalf("Failed test #%d, Failed getting list of maximums. Want %#v got %#v", i, test.maximums, got)
		}
	}
}
