package graph

import "testing"

/*
TestIsDAG tests solution(s) with the following signature and problem description:

	Vertex struct {
			// Val is the value of the vertex
			Val int

			// The edges that this Vertex is connected to
			Edges []*Vertex
	}
	func IsDAG(graph []*Vertex) bool

Given a directed graph determine if it's a DAG or not. A directed acyclic graph (DAG)
is a directed graph that has no cycles.
*/
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
