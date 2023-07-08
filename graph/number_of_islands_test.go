package graph

import (
	"testing"
)

/*
TestNumberOfIslands tests solution(s) with the following signature and problem description:

	func NumberOfIslands(grid [][]int) int

Given a grid in which 0 represents water and 1 represents a piece of land, return the
number of islands. An Island is one or more piece of land that could be connected to other
pieces of land on left, right, top or bottom.
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
