package heap

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		numbers    []int
		k          int
		maxSliding []int
	}{
		{[]int{}, 2, []int{}},
		{[]int{1, 4, 5, -2, 4, 6}, 1, []int{1, 4, 5, -2, 4, 6}},
		{[]int{1, 4, 5, -2, 4, 6}, 2, []int{4, 5, 5, 4, 6}},
		{[]int{1, 4, 5, -2, 4, 6}, 3, []int{5, 5, 5, 6}},
		{[]int{1, 4, 5, -2, 4, 6}, 4, []int{5, 5, 6}},
		{[]int{1, 4, 5, -2, 4, 6}, 6, []int{6}},
	}

	for i, test := range tests {
		got := MaxSlidingWindow(test.numbers, test.k)
		if !reflect.DeepEqual(got, test.maxSliding) {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.maxSliding, got)
		}
	}
}
