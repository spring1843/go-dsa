package dnc

import (
	"reflect"
	"testing"
)

/*
TestTowerOfHanoi tests solution(s) with the following signature and problem description:

	func TowerOfHanoi(n, start, end int) [][2]int

Given n, number of disks, and start and end tower, return the moves it takes to move all
disks from start to end tower. The disks are stacked on top of each other with the
lightest being on top and heaviest being in the bottom. A heavier disk cannot be placed
on a lighter disk. You can move one disk at a time.
*/
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
