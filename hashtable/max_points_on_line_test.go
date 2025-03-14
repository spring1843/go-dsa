package hashtable

import "testing"

/*
TestMaxPointsOnALine tests solution(s) with the following signature and problem description:

	func MaxPointsOnALine(points [][]int) int

Given multiple coordinates of points like, return the maximum number of points that are on the same.

For example given {[1,1], [2,2], [3,3], [4,5]} return 3 because the points [1,1], [2,2], [3,3] are on
the same line.
*/
func TestMaxPointsOnALine(t *testing.T) {
	tests := []struct {
		points   [][]int
		maxLines int
	}{
		{[][]int{}, 0},
		{[][]int{{2, 5}}, 1},
		{[][]int{{2, 5}, {4, 1}}, 2},
		{[][]int{{2, 5}, {2, 2}, {1, 1}, {7, 7}, {1, 5}}, 3},
		{[][]int{{2, 5}, {2, 2}, {2, 2}, {1, 1}, {7, 7}, {1, 5}}, 4},
		{[][]int{{-2, 1}, {-3, 2}, {-4, 3}, {-5, 4}, {-1, 1}, {-5, 3}}, 4},
		{[][]int{{2, 1}, {3, 2}, {5, 4}, {2, 3}, {3, 4}, {4, 5}}, 3},
	}

	for i, test := range tests {
		got := MaxPointsOnALine(test.points)
		if got != test.maxLines {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.maxLines, got)
		}
	}
}
