package hashtable

import (
	"reflect"
	"testing"
)

func TestSumUpToK(t *testing.T) {
	var tests = []struct {
		k                int
		numbers, indices []int
	}{
		{6, []int{3, 3}, []int{0, 1}},
		{3, []int{1, 2, 3, 4}, []int{0, 1}},
		{5, []int{1, 2, 3, 6}, []int{1, 2}},
		{7, []int{1, 2, 3, 4}, []int{2, 3}},
	}

	for i, test := range tests {
		got := SumUpToK(test.numbers, test.k)
		if !reflect.DeepEqual(got, test.indices) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.indices, got)
		}
	}
}
