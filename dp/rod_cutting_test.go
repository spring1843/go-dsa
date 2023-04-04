package dp

import (
	"reflect"
	"testing"
)

func TestRodCutting(t *testing.T) {
	tests := []struct {
		snacks   []int
		n        int
		solution int
	}{
		{[]int{1}, 1, 1},
		{[]int{1, 2}, 2, 2},
		{[]int{1, 3, 5, 10, 15}, 2, 3},
		{[]int{1, 3, 5, 10, 15}, 3, 5},
		{[]int{1, 5, 8, 9, 10}, 4, 10},
		{[]int{1, 5, 8, 9, 10, 20, 30}, 4, 10},
	}

	for i, test := range tests {
		if got := CutRod(test.snacks, test.n); !reflect.DeepEqual(got, test.solution) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.solution, got)
		}
	}
}
