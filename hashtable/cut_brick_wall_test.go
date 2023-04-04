package hashtable

import "testing"

func TestCutBrickWall(t *testing.T) {
	tests := []struct {
		cutPosition int
		bricks      [][]int
	}{
		{0, [][]int{}},
		{3, [][]int{{1, 2, 2}, {3, 2}, {2, 3}}},
		{5, [][]int{{5, 4, 1}, {2, 2, 1, 3, 2}, {3, 2, 4, 1}, {3, 4, 3}, {2, 3, 5}, {3, 2, 1, 4}}},
		{1, [][]int{{2, 5, 4, 1}, {1, 2, 2, 1, 3, 2}, {1, 3, 2, 4, 1}, {1, 3, 4, 3}}},
		{1, [][]int{{2, 5, 4, 1}, {1, 2, 2, 1, 3, 2}, {1, 3, 2, 4, 1}, {1, 3, 4, 3}}},
	}

	for i, test := range tests {
		got := CutBrickWall(test.bricks)
		if got != test.cutPosition {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.cutPosition, got)
		}
	}
}
