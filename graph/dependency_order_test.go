package graph

import (
	"errors"
	"strings"
	"testing"
)

const (
	lineBreak = "\n"
	delimiter = ","
)

/*
TestDependencyOrder tests solution(s) with the following signature and problem description:

	func DependencyOrder(graph []*VertexWithIngress)  []*VertexWithIngress

Given a multiline input, each line starting with  an element and followed by a comma separated
list of dependent elements, output the elements in the order in which the dependencies are met.

In real world this is similar to courses and their prerequisite where a prerequisite must be
taken before a course can be taken. This is also similar to the dependencies of a software,
where there are dependencies of dependencies and a software can only built when all of its
dependencies are built.

For example the following input:
a,b,c,d
b,c
d,e,f

Means that a depends on b, c, d; b depends on c; d depends on e, f.
The output should be f,e,c,d,b,a. Note that the order of the elements in the output is not
important, but the order of the dependencies is important.
*/
func TestDependencyOrder(t *testing.T) {
	tests := []struct {
		input       string
		order       string
		expectedErr error
	}{
		{"", "", nil},
		{"a,b", "b,a", nil},
		{"a,b,c\nc,a", "", ErrNotADAG},
		{"a,b,c,d\nb,c\nd,e,f", "f,e,c,d,b,a", nil},
		{"a,f\nb,f\nc,f\nd,f\ne,f", "f,e,d,c,b,a", nil},
		{"a,f\nb,f\nc,f\nd,f\ne,f\nf,g\ng,h", "h,g,f,e,d,c,b,a", nil},
		{"a,f\nb,f\nc,f\nd,f\ne,f\nf,g\ng,h\nd,h", "h,g,f,e,d,c,b,a", nil},
	}

	for i, test := range tests {
		orderedGraph, err := DependencyOrder(toDependencyGraph(test.input))
		got := ""
		for i := range orderedGraph {
			got += orderedGraph[i].Val.(string) + delimiter
		}
		if len(orderedGraph) > 0 {
			got = strings.TrimSuffix(got, delimiter)
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

		if got != test.order {
			t.Fatalf("Failed test case #%d. Want %s got %s", i, test.order, got)
		}
	}
}

func toDependencyGraph(input string) []*VertexWithIngress {
	graph := []*VertexWithIngress{}
	graphMap := make(map[string]*VertexWithIngress)
	for _, line := range strings.Split(input, lineBreak) {
		firstElement := ""
		for _, element := range strings.Split(line, delimiter) {
			if _, ok := graphMap[element]; !ok {
				vertex := &VertexWithIngress{Val: element, Edges: []*VertexWithIngress{}}
				graphMap[element] = vertex
				graph = append(graph, vertex)
			}

			if firstElement == "" {
				firstElement = element
				continue
			}
			graphMap[firstElement].Edges = append(graphMap[firstElement].Edges, graphMap[element])
		}
	}
	return graph
}
