package graph

import (
	"testing"
)

/*
TestNumberOfIslands tests solution(s) with the following signature and problem description:

	func NumberOfIslands(grid [][]int) int {

Returns the number of islands or connected lands in a grid where 0 represents water
and one represents land.
*/
func TestNumberOfIslands(t *testing.T) {
	tests := []struct {
		grid        [][]int
		islandCount int
	}{
		{[][]int{}, 0},
		{[][]int{{1, 0, 0}, {0, 0, 0}}, 1},
		{[][]int{{1, 0, 1}, {0, 0, 0}}, 2},
		{[][]int{{1, 0, 1}, {0, 0, 0}, {0, 1, 0}}, 3},
		{[][]int{{1, 0, 1}, {0, 0, 0}, {0, 1, 1}}, 3},
		{[][]int{{1, 0, 1}, {0, 0, 1}, {0, 1, 1}}, 2},
	}

	for i, test := range tests {
		if got := NumberOfIslands(test.grid); got != test.islandCount {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.islandCount, got)
		}
	}
}
