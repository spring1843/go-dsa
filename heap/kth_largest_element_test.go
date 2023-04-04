package heap

import (
	"testing"
)

func TestKthLargestElement(t *testing.T) {
	var tests = []struct {
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
