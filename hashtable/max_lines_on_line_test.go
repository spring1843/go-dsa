package hashtable

import "testing"

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
