package heap

import (
	"reflect"
	"testing"
)

func TestKClosestPointToOrigin(t *testing.T) {
	var tests = []struct {
		points     [][]int
		k          int
		kthClosest [][]int
	}{
		{[][]int{{}}, 1, [][]int{{}}},
		{[][]int{{10, 10}}, 1, [][]int{{10, 10}}},
		{[][]int{{10, 10}, {-10, -10}, {11, 11}, {12, 12}}, 1, [][]int{{10, 10}}},
		{[][]int{{10, 10}, {-10, -10}, {11, 11}, {12, 12}}, 2, [][]int{{10, 10}, {-10, -10}}},
		{[][]int{{35, 11}, {14, -12}, {50, 100}, {2, 19}, {-17, -18}, {1, 1}, {27, 22}}, 1, [][]int{{1, 1}}},
	}

	for i, test := range tests {
		got := KClosestPointToOrigin(test.points, test.k)
		if !reflect.DeepEqual(got, test.kthClosest) {
			t.Fatalf("Failed test case #%d. Want %d got %d", i, test.kthClosest, got)
		}
	}
}
