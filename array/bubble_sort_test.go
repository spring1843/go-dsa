package array

import (
	"slices"
	"testing"
)

/*
TestBubbleSort tests solution(s) with the following signature and problem description:

	BubbleSort(input []int)

Given an array of unsorted integers, sort the array using the Bubble Sort algorithm.
The algorithm should be in-place, meaning it should not create a new array and it should
work by swapping elements in the array until it is sorted.
*/
func TestBubbleSort(t *testing.T) {
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
		BubbleSort(test.input)
		if !slices.Equal(test.input, test.sorted) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.sorted, test.input)
		}
	}
}
