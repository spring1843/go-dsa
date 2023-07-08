package dnc

import (
	"reflect"
	"testing"
)

/*
TestQuickSort tests solution(s) with the following signature and problem description:

	func QuickSort(list []int) []int

Given a list of integers like {3,1,2}, return a sorted set like {1,2,3} using Quick Sort.
*/
func TestQuickSort(t *testing.T) {
	tests := []struct {
		list   []int
		sorted []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 3, 2}, []int{1, 2, 3}},
		{[]int{-1, 3, 2, 0, 4}, []int{-1, 0, 2, 3, 4}},
	}

	for i, test := range tests {
		if got := QuickSort(test.list); !reflect.DeepEqual(got, test.sorted) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.sorted, got)
		}
	}
}
