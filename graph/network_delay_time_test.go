package graph

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		edges              [][]int
		n, k, shortestPath int
	}{
		{[][]int{{0, 1, 1}}, 2, 1, -1},
		{[][]int{{0, 1, 1}, {1, 2, 2}, {2, 3, 2}, {3, 1, 3}}, 4, 0, 5},
		{[][]int{{0, 1, 1}, {1, 2, 2}, {2, 3, 2}, {3, 1, 3}}, 4, 1, -1},
		{[][]int{{0, 1, 2}, {0, 2, 3}, {2, 4, 2}, {1, 3, 2}, {3, 4, 3}, {4, 5, 2}}, 5, 0, 5},
	}

	for i, test := range tests {
		if got := NetworkDelayTime(test.n, test.k, test.edges); got != test.shortestPath {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.shortestPath, got)
		}
	}
}
