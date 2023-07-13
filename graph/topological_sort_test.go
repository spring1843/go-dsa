package graph

import (
	"errors"
	"reflect"
	"testing"
)

var readmeGraphs = map[string][][]int{
	"Figure_1_A": {{2, 3, 4}, {3}, {5}, {3}, {}},
	"Figure_1_B": {{2}, {3}, {4}, {1}},
	"Figure_1_C": {{2}, {4}, {2}, {5}, {}},
}

/*
TestTopologicalSort tests solution(s) with the following signature and problem description:

	func TopologicalSort(graph []*VertexWithIngress) ([]int, error)

In a DAG the topological order returns elements such that if there's a path from v(i) to
v(j), then v(i) appears before v(j) in the ordering. Given a graph like (B) of _Figure 1_
output the graph in topological order like {1,2,3,4}.
*/
func TestTopologicalSort(t *testing.T) {
	tests := []struct {
		graph           [][]int
		topologicalSort []int
		expectedErr     error
	}{
		{[][]int{}, []int{}, nil},
		{[][]int{{}}, []int{1}, nil},
		{[][]int{{}, {}}, []int{1, 2}, nil},
		{[][]int{{}, {}, {4}, {}}, []int{1, 2, 3, 4}, nil},
		{[][]int{{}, {1}, {}, {3}}, []int{2, 4, 1, 3}, nil},
		{[][]int{{}, {1}, {}, {3}}, []int{2, 4, 1, 3}, nil},
		{[][]int{{2, 5}, {3, 4}, {}, {5}, {}}, []int{1, 2, 3, 4, 5}, nil},
		{[][]int{{2}, {3}, {1}}, []int{}, ErrNotADAG},
		{readmeGraphs["Figure_1_A"], []int{1, 2, 4, 3, 5}, nil},
		{readmeGraphs["Figure_1_B"], []int{}, ErrNotADAG},
		{readmeGraphs["Figure_1_C"], []int{1, 3, 2, 4, 5}, nil},
	}

	for i, test := range tests {
		orderedGraph, err := TopologicalSort(toGraphWithIngress(test.graph))
		got := make([]int, 0, len(orderedGraph))
		for i := range orderedGraph {
			got = append(got, orderedGraph[i].Val.(int))
		}

		if err != nil {
			if test.expectedErr == nil {
				t.Fatalf("Failed test case #%d. Unexpected error. Error :%s", i, err)
			}
			if !errors.Is(err, test.expectedErr) {
				t.Fatalf("Failed test case #%d. Unexpected error. Want %s, Got Error :%s", i, test.expectedErr, err)
			}
		}
		if err == nil && test.expectedErr != nil {
			t.Fatalf("Failed test case #%d. Expected error %s did not occur", i, test.expectedErr)
		}

		if !reflect.DeepEqual(got, test.topologicalSort) {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, test.topologicalSort, got)
		}
	}
}

func toGraphWithIngress(graph [][]int) []*VertexWithIngress {
	graphMap := make(map[int]*VertexWithIngress)

	for i := range graph {
		graphMap[i+1] = &VertexWithIngress{Val: i + 1, Edges: []*VertexWithIngress{}}
	}

	for i, v := range graph {
		graphMap[i+1].Edges = make([]*VertexWithIngress, len(v))
		for j, e := range v {
			graphMap[i+1].Edges[j] = graphMap[e]
		}
	}

	output := make([]*VertexWithIngress, len(graphMap))
	for i := 0; i < len(graph); i++ {
		output[i] = graphMap[i+1]
	}
	return output
}
