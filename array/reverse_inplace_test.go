package array

import (
	"reflect"
	"testing"
)

func TestReverseInplace(t *testing.T) {
	tests := []struct {
		list     []int
		start    int
		end      int
		reversed []int
	}{
		{[]int{}, 0, 0, []int{}},
		{[]int{1, 2, 3}, 1, 2, []int{1, 3, 2}},
		{[]int{1, 2, 3}, 2, 1, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 1, 2, []int{1, 3, 2, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 0, 4, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6}, 0, 4, []int{5, 4, 3, 2, 1, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, 0, 3, []int{4, 3, 2, 1, 5, 6}},
	}

	for i, test := range tests {
		ReverseInplace(test.list, test.start, test.end)
		if !reflect.DeepEqual(test.list, test.reversed) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.reversed, test.list)
		}
	}
}
