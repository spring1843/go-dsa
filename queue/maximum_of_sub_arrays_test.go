package queue

import (
	"reflect"
	"testing"
)

func TestMaxOfKLengthSubArrays(t *testing.T) {
	var tests = []struct {
		k        int
		list     []int
		maximums []int
	}{
		{3, []int{5, 3, 5, 6, 7, 8}, []int{5, 6, 7, 8}},
		{2, []int{5, 3, 5, 6, 7, 8}, []int{5, 5, 6, 7, 8}},
		{3, []int{1, 2, 3, 4, 5}, []int{3, 4, 5}},
	}

	for i, test := range tests {
		got, err := MaxOfKLengthSubArrays(test.list, test.k)
		if err != nil {
			t.Fatalf("Unexpected error occurred. Error : %s", err)
		}
		if !reflect.DeepEqual(got, test.maximums) {
			t.Fatalf("Failed test #%d, Failed getting list of maximums. Want %#v got %#v", i, test.maximums, got)
		}
	}
}
