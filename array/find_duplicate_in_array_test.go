package array

import (
	"testing"
)

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		list      []int
		duplicate int
	}{
		{[]int{}, -1},
		{[]int{1, 2, 2}, 2},
		{[]int{1, 2, 3}, -1},
		{[]int{1, 1, 2, 3}, 1},
		{[]int{1, 2, 2, 3}, 2},
	}

	for i, test := range tests {
		if got := FindDuplicate(test.list); got != test.duplicate {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.duplicate, got)
		}
	}
}
