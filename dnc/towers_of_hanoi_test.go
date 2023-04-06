package dnc

import (
	"reflect"
	"testing"
)

func TestTowerOfHanoi(t *testing.T) {
	tests := []struct {
		n, start, end int
		moves         [][2]int
	}{
		{0, 1, 1, [][2]int{{1, 1}}},
		{1, 1, 1, [][2]int{{1, 1}}},
		{1, 1, 2, [][2]int{{1, 2}}},
		{2, 1, 3, [][2]int{{1, 2}, {1, 3}, {2, 3}}},
		{3, 1, 3, [][2]int{{1, 3}, {1, 2}, {3, 2}, {1, 3}, {2, 1}, {2, 3}, {1, 3}}},
		{4, 1, 3, [][2]int{{1, 2}, {1, 3}, {2, 3}, {1, 2}, {3, 1}, {3, 2}, {1, 2}, {1, 3}, {2, 3}, {2, 1}, {3, 1}, {2, 3}, {1, 2}, {1, 3}, {2, 3}}},
	}

	for i, test := range tests {
		if got := TowerOfHanoi(test.n, test.start, test.end); !reflect.DeepEqual(got, test.moves) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.moves, got)
		}
	}
}
