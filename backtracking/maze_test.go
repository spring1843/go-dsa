package backtracking

import (
	"testing"
)

/*
TestMaze tests solution(s) with the following signature and problem description:

	func Maze(walls [][2]int, start, finish [2]int,m,n int) string

Given the coordinates of walls in a m x n maze in tuples of {row, col} format, and start
and finish coordinates in the same format, find a path from start to finish. return
a string of directions like `lrud` (left, right, up, down) to get a robot from
start to finish.

The robot can only move in the four left, right, down and up directions and not
through walls. The robot does not know where the finish line is so it has to
explore every possible cell in the order of directions given.
*/
func TestMaze(t *testing.T) {
	tests := []struct {
		m      int
		n      int
		walls  [][2]int
		start  [2]int
		finish [2]int
		moves  string
	}{
		{1, 1, [][2]int{}, [2]int{0, 0}, [2]int{0, 1}, ""},
		{5, 5, [][2]int{}, [2]int{0, 0}, [2]int{0, 1}, "r"},
		{10, 10, [][2]int{}, [2]int{0, 0}, [2]int{0, 1}, "r"},
		{5, 5, [][2]int{}, [2]int{0, 0}, [2]int{4, 4}, "rrrrdlllldrrrrdlllldrrrr"},
		{5, 5, [][2]int{{1, 1}, {1, 2}, {1, 3}, {2, 3}, {3, 3}, {3, 4}}, [2]int{0, 0}, [2]int{2, 4}, "rrrrdd"},
		{5, 5, [][2]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 3}, {2, 1}, {2, 2}, {2, 3}, {3, 1}, {4, 1}}, [2]int{4, 0}, [2]int{1, 2}, "uuurr"},
		{5, 5, [][2]int{{1, 0}, {1, 1}, {1, 2}, {1, 3}, {3, 1}, {3, 2}, {3, 3}, {3, 4}}, [2]int{0, 0}, [2]int{4, 4}, "rrrrddllllddrrrr"},
		{5, 5, [][2]int{{1, 0}, {1, 1}, {1, 4}, {1, 3}, {3, 1}, {3, 2}, {3, 3}, {3, 4}}, [2]int{0, 0}, [2]int{4, 4}, "rrddllddrrrr"},
	}

	for i, test := range tests {
		if got := Maze(test.m, test.n, test.walls, test.start, test.finish); test.moves != got {
			t.Fatalf("Failed test case #%d. Want %s got %s", i, test.moves, got)
		}
	}
}
