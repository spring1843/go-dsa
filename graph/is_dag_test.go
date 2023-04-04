package graph

import "testing"

func TestIsDAG(t *testing.T) {
	tests := []struct {
		graph [][]int
		isDAG bool
	}{
		{[][]int{{2}, {1, 3}, {}}, false},
		{[][]int{{2, 3}, {3}, {4}, {}}, true},
		{[][]int{{2, 3}, {4, 5}, {}, {}, {}}, true},
		{[][]int{{2, 3, 4}, {3}, {4}, {5}, {}}, true},
		{[][]int{{2, 3, 4}, {3, 5}, {5}, {3, 5}, {}}, true},
		{[][]int{{2, 3, 4}, {3, 5}, {5}, {3, 5}, {5, 3}}, false},
		{readmeGraphs["Figure_1_A"], true},
		{readmeGraphs["Figure_1_B"], false},
		{readmeGraphs["Figure_1_C"], true},
	}

	for i, test := range tests {
		if got := IsDAG(toGraph(test.graph)); got != test.isDAG {
			t.Fatalf("Failed case #%d. Want %v got %v", i, test.isDAG, got)
		}
	}
}
