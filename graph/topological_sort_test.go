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

func TestTopologicalSort(t *testing.T) {
	var tests = []struct {
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
		{readmeGraphs["Figure_1_A"], []int{1, 2, 4, 3, 5}, nil},
		{readmeGraphs["Figure_1_B"], nil, ErrNotADAG},
		{readmeGraphs["Figure_1_C"], []int{1, 3, 2, 4, 5}, nil},
	}

	for i, test := range tests {
		got, err := TopologicalSort(toGraph(test.graph))
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
