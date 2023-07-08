package array

import (
	"reflect"
	"testing"
)

/*
TestEqualSumSubArrays tests solution(s) with the following signature and problem description:

	func EqualSubArrays(list []int) [][]int

Given an list of integers A, return two sub-arrays with equal sums without changing the
order of the elements in the list.
*/
func TestEqualSumSubArrays(t *testing.T) {
	tests := []struct {
		list      []int
		subArrays [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{1}, [][]int{}},
		{[]int{1, 2, 2, 3}, [][]int{}},
		{[]int{1, 2, 3}, [][]int{{1, 2}, {3}}},
		{[]int{2, 3, 5}, [][]int{{2, 3}, {5}}},
		{[]int{1, 7, 3, 5}, [][]int{{1, 7}, {3, 5}}},
		{[]int{-4, 1, 1, 1, 1}, [][]int{}},
		{[]int{4, 1, 1, 1, 1}, [][]int{{4}, {1, 1, 1, 1}}},
		{[]int{1, 1, 1, 1, 4}, [][]int{{1, 1, 1, 1}, {4}}},
	}

	for i, test := range tests {
		if got := EqualSubArrays(test.list); !reflect.DeepEqual(got, test.subArrays) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.subArrays, got)
		}
	}
}
