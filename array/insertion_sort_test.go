package array

import (
	"slices"
	"testing"
)

/*
TestInsertionSort tests solution(s) with the following signature and problem description:

	InsertionSort(input []int)

Given an array of unsorted integers, sort the array using the Insertion Sort algorithm.
The algorithm should be in-place, meaning it should not create a new array. The algorithm
works by dividing the input array into sorted and unsorted sections, and moving items from
the unsorted section into the sorted section, one item at a time.
*/
func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input, sorted []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 3, 1}, []int{1, 2, 3}},
		{[]int{4, 2, 3, 1, 5}, []int{1, 2, 3, 4, 5}},
	}
	for i, test := range tests {
		InsertionSort(test.input)
		if !slices.Equal(test.input, test.sorted) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.sorted, test.input)
		}
	}
}
