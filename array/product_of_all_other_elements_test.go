package array

import (
	"reflect"
	"testing"
)

func TestProductOfAllOtherElements(t *testing.T) {
	tests := []struct {
		list     []int
		products []int
	}{
		{[]int{}, []int{}},
		{[]int{2, 3}, []int{3, 2}},
		{[]int{1, 2, 3}, []int{6, 3, 2}},
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
	}

	for i, test := range tests {
		if got := ProductOfAllOtherElements(test.list); !reflect.DeepEqual(got, test.products) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.products, got)
		}
	}
}
