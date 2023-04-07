package dnc

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		sortedNumbers    []int
		search, position int
	}{
		{[]int{}, 3, -1},
		{[]int{1}, 1, 0},
		{[]int{1}, 2, -1},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 5, 10}, 4, -1},
		{[]int{1, 2, 3, 8, 10}, 8, 3},
		{[]int{1, 2, 3, 8, 10, 11}, 8, 3},
		{[]int{1, 2, 3, 4, 5, 9}, 4, 3},
	}

	for i, test := range tests {
		if got := BinarySearch(test.sortedNumbers, test.search); got != test.position {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.position, got)
		}
	}
}
