package array

import (
	"reflect"
	"testing"
)

func TestRotateKSteps(t *testing.T) {
	tests := []struct {
		list        []int
		steps       int
		rotatedList []int
	}{
		{[]int{}, 0, []int{}},
		{[]int{1, 2, 3}, 1, []int{3, 1, 2}},
		{[]int{1, 2, 3}, 2, []int{2, 3, 1}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
	}

	for i, test := range tests {
		RotateKSteps(test.list, test.steps)
		if !reflect.DeepEqual(test.list, test.rotatedList) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.rotatedList, test.list)
		}
	}
}
