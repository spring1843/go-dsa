package heap

import "testing"

/*
TestKthLargestElement tests solution(s) with the following signature and problem description:

	func KthLargestElement(elements []int, k int) int

Given a slice of integers and an integer k, return the kth largest element.

For example given {3,5,6,3,1,2,9} and k=3 return 5, because largest integers are 9, 6, 5 and
the 3rd largest is 5.
*/
func TestKthLargestElement(t *testing.T) {
	tests := []struct {
		list       []int
		k          int
		kthLargest int
	}{
		{[]int{1}, 1, 1},
		{[]int{1, 2, 3, 4, 5, 6}, 3, 4},
		{[]int{6, 1, 2, 3, 4, 5}, 1, 6},
		{[]int{6, 1, 2, 3, 4, 5}, 6, 1},
	}

	for i, test := range tests {
		got := KthLargestElement(test.list, test.k)
		if got != test.kthLargest {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.kthLargest, got)
		}
	}
}

func TestMinHeap(t *testing.T) {
	minHeap := new(minimumHeap)
	minHeap.Push(1)
	got := minHeap.Pop().(int)
	if got != 1 {
		t.Fatalf("expected %d got %d", got, 1)
	}
}
