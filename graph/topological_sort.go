package graph

import (
	"container/list"
	"errors"
)

// VertexWithIngress is a Vertex with the count of vertices that connect to it.
type VertexWithIngress struct {
	// Val is the value of the vertex
	Val int

	// The edges that this Vertex is connected to
	Edges   []*VertexWithIngress
	ingress int
}

// ErrNotADAG occurs when a graph has a cycle and hence not a DAG where a DAW was expected.
var ErrNotADAG = errors.New("not a DAG")

// TopologicalSort takes a vertex of a DAG and returns the value of all its
// connected vertices in topological order.
func TopologicalSort(graph []*VertexWithIngress) ([]int, error) {
	var (
		output = []int{}
		// queue  = new(vertexWithIngressQueue)
		queue = list.New()
		i     = 0
	)

	setIngress(graph)
	for _, vertex := range graph {
		if vertex.ingress == 0 {
			queue.PushBack(vertex)
		}
	}

	for queue.Len() != 0 {
		i++
		vertex := queue.Front().Value.(*VertexWithIngress)
		queue.Remove(queue.Front())

		output = append(output, vertex.Val)

		for _, neighbor := range vertex.Edges {
			neighbor.ingress--
			if neighbor.ingress == 0 {
				queue.PushBack(neighbor)
			}
		}
	}

	if i != len(graph) {
		return nil, ErrNotADAG
	}

	return output, nil
}

func setIngress(graph []*VertexWithIngress) {
	for _, vertex := range graph {
		for _, neighbor := range vertex.Edges {
			neighbor.ingress++
		}
	}
}
