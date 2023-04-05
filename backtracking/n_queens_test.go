package backtracking

import (
	"reflect"
	"testing"
)

func TestNQueens(t *testing.T) {
	tests := []struct {
		n         int
		solutions []Chessboard
	}{
		{0, []Chessboard{}},
		{1, []Chessboard{{{queen}}}},
		{2, []Chessboard{}},
		{3, []Chessboard{}},
		{4, []Chessboard{{{empty, queen, empty, empty}, {empty, empty, empty, queen}, {queen, empty, empty, empty}, {empty, empty, queen, empty}}, {{empty, empty, queen, empty}, {queen, empty, empty, empty}, {empty, empty, empty, queen}, {empty, queen, empty, empty}}}},
	}

	for i, test := range tests {
		if got := NQueens(test.n); !reflect.DeepEqual(test.solutions, got) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.solutions, got)
		}
	}
}
