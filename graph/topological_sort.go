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
	Edges []*VertexWithIngress

	// Ingress is the number of vertices that connect to this vertex
	Ingress int
}

// ErrNotADAG occurs when a graph is not a Direct Acyclic Graph - DAG.
var ErrNotADAG = errors.New("graph is not a Direct Acyclic Graph - DAG")

// TopologicalSort takes a vertex of a DAG and returns the value of all its
// connected vertices in topological order.
func TopologicalSort(graph []*VertexWithIngress) ([]int, error) {
	var (
		output = []int{}
		queue  = list.New()
		i      = 0
	)

	for _, vertex := range graph {
		for _, neighbor := range vertex.Edges {
			neighbor.Ingress++
		}
	}

	for _, vertex := range graph {
		if vertex.Ingress == 0 {
			queue.PushBack(vertex)
		}
	}

	for queue.Len() != 0 {
		i++
		vertex := queue.Remove(queue.Front()).(*VertexWithIngress)
		output = append(output, vertex.Val)

		for _, neighbor := range vertex.Edges {
			neighbor.Ingress--
			if neighbor.Ingress == 0 {
				queue.PushBack(neighbor)
			}
		}
	}

	if i != len(graph) {
		return nil, ErrNotADAG
	}
	return output, nil
}
