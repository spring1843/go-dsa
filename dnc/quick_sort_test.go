package dnc

import (
	"slices"
	"testing"
)

/*
TestQuickSort tests solution(s) with the following signature and problem description:

	func QuickSort(list []int) []int

Given a slice of integers like, return a sorted slice of integers using quick sort.

Quick sort works by selecting a 'pivot' element from the array and partitioning
the other elements into two sub-arrays, according to whether they are less than
or greater than the pivot. The sub-arrays are then sorted recursively. The
base case of the recursion is an array of size 0 or 1, which is always sorted.

For example given {3,1,2}, return a {1,2,3}. The quick sort algorithm would first
choose a pivot element, say 2. It would then partition the array into two
sub-arrays: {1} and {3}. It would then recursively sort the two sub-arrays
and combine them with the pivot element to produce the final sorted.
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
		if got := QuickSort(test.list); !slices.Equal(got, test.sorted) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.sorted, got)
		}
	}
}
