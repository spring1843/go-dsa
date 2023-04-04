package backtracking

import (
	"reflect"
	"testing"
)

func TestNQueens(t *testing.T) {
	tests := []struct {
		n         int
		solutions []chessboard
	}{
		{0, []chessboard{}},
		{1, []chessboard{{{queen}}}},
		{2, []chessboard{}},
		{3, []chessboard{}},
		{4, []chessboard{{{empty, queen, empty, empty}, {empty, empty, empty, queen}, {queen, empty, empty, empty}, {empty, empty, queen, empty}}, {{empty, empty, queen, empty}, {queen, empty, empty, empty}, {empty, empty, empty, queen}, {empty, queen, empty, empty}}}},
	}

	for i, test := range tests {
		if got := NQueens(test.n); !reflect.DeepEqual(test.solutions, got) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.solutions, got)
		}
	}
}
