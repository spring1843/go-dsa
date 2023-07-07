package array

import (
	"reflect"
	"testing"
)

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
		if !reflect.DeepEqual(test.input, test.sorted) {
			t.Fatalf("Failed test case #%d. Want %v got %v", i, test.sorted, test.input)
		}
	}
}
