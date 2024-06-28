package graph

import (
	"slices"
	"testing"
)

/*
TestIterativeTraversal tests solution(s) with the following signature and problem description:

	Vertex struct {
			Val int
			Edges []*Vertex
	}
	func IterativeTraversal(graph []*Vertex) ([]int, []int)

Implement BFS and DFS on a graph without using recursion, and return the value of each vertex
as it is visited.
*/
func TestIterativeTraversal(t *testing.T) {
	tests := []struct {
		graph    [][]int
		bfs, dfs []int
	}{
		{[][]int{{2, 3}, {3}, {4}, {}}, []int{1, 2, 3, 4}, []int{1, 3, 4, 2}},
		{[][]int{{2, 3}, {4, 5}, {}, {}, {}}, []int{1, 2, 3, 4, 5}, []int{1, 3, 2, 5, 4}},
		{[][]int{{2, 3, 4}, {3}, {4}, {5}, {}}, []int{1, 2, 3, 4, 5}, []int{1, 4, 5, 3, 2}},
		{[][]int{{2, 3, 4}, {3, 5}, {5}, {3, 5}, {}}, []int{1, 2, 3, 4, 5}, []int{1, 4, 5, 3, 2}},
		{readmeGraphs["Figure_1_A"], []int{1, 2, 3, 4, 5}, []int{1, 4, 3, 5, 2}},
		// {readmeGraphs["Figure_1_B"], []int{1, 2, 3, 4, 5}, []int{1, 4, 5, 3, 2}}, TODO as visited to avoid cycles
		{readmeGraphs["Figure_1_C"], []int{1, 2, 4, 5}, []int{1, 2, 4, 5}},
	}

	for i, test := range tests {
		bfs, dfs := IterativeTraversal(toGraph(test.graph))
		if !slices.Equal(bfs, test.bfs) {
			t.Fatalf("Failed BFS test case #%d. Want %#v got %#v", i, test.bfs, bfs)
		}
		if !slices.Equal(dfs, test.dfs) {
			t.Fatalf("Failed DFS test case #%d. Want %#v got %#v", i, test.dfs, dfs)
		}
	}
}

// toGraph converts an adjacency formatted graph into a connected collection of graph vertices.
func toGraph(graph [][]int) []*Vertex {
	graphMap := make(map[int]*Vertex)

	for i := range graph {
		graphMap[i+1] = newVertex(i + 1)
	}

	for i, v := range graph {
		graphMap[i+1].Edges = make([]*Vertex, len(v))
		for j, e := range v {
			graphMap[i+1].Edges[j] = graphMap[e]
		}
	}

	output := make([]*Vertex, len(graphMap))
	for i := range len(graph) {
		output[i] = graphMap[i+1]
	}
	return output
}
