package heap

import (
	"math"
	"slices"
	"testing"
)

/*
TestSort tests solution(s) with the following signature and problem description:

	func HeapSort(list []int) []int

Given a slice of integers, return a sorted slice using Heap Sort.

For example given {3,1,2}, return {1,2,3}.

Heap sort works by inserting all items to a heap and then popping them one by. As heap returns the minimum
element with each pop the outcome is a sorted list of items..
*/
func TestSort(t *testing.T) {
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
		if got := Sort(test.list); !slices.Equal(got, test.sorted) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.sorted, got)
		}
	}
}

func TestMinHeapImplementation(t *testing.T) {
	heap := NewMinHeap()
	if got := heap.Pop(); got != math.MinInt {
		t.Fatalf("Failed test case. Want %v got %v", math.MinInt, got)
	}
	heap.Push(1)
	if got := heap.Pop(); got != 1 {
		t.Fatalf("Failed test case. Want %v got %v", 1, got)
	}
}
