package array

import (
	"reflect"
	"testing"
)

func TestZeroSumTriplets(t *testing.T) {
	tests := []struct {
		list     []int
		triplets [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{1}, [][]int{}},
		{[]int{1, 2}, [][]int{}},
		{[]int{1, 2, 3}, [][]int{}},
		{[]int{1, -4, 3}, [][]int{{-4, 1, 3}}},
		{[]int{1, 2, -4, 6, 3}, [][]int{{-4, 1, 3}}},
		{[]int{-1, -2, -8, 6, 2, 3}, [][]int{{-8, 2, 6}, {-2, -1, 3}}},
		{[]int{1, -2, -4, 5, -2, 4, 1, 3}, [][]int{{-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}}},
	}

	for i, test := range tests {
		if got := ZeroSumTriplets(test.list); !reflect.DeepEqual(got, test.triplets) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.triplets, got)
		}
	}
}
