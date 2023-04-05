package graph

import "errors"

// VertexWithIngress is a Vertex with the count of vertices that connect to it.
type VertexWithIngress struct {
	Vertex
}

// ErrNotADAG occurs when a graph has a cycle and hence not a DAG where a DAW was expected.
var ErrNotADAG = errors.New("not a DAG")

// TopologicalSort takes a vertex of a DAG and returns the value of all its
// connected vertices in topological order.
func TopologicalSort(graph []*Vertex) ([]int, error) {
	var (
		output  = []int{}
		ingress = ingress(graph)
		queue   = new(queue)
		i       = 0
	)

	for _, vertex := range graph {
		if ingress[vertex] == 0 {
			queue.Push(vertex)
		}
	}

	for queue.Len() != 0 {
		i++
		vertex := queue.Pop()
		output = append(output, vertex.Val)

		for _, neighbor := range vertex.Edges {
			ingress[neighbor]--
			if ingress[neighbor] == 0 {
				queue.Push(neighbor)
			}
		}
	}

	if i != len(graph) {
		return nil, ErrNotADAG
	}

	return output, nil
}

func ingress(graph []*Vertex) map[*Vertex]int {
	ingress := zeroedMap(graph)
	for _, vertex := range graph {
		for _, neighbor := range vertex.Edges {
			ingress[neighbor]++
		}
	}
	return ingress
}

func zeroedMap(graph []*Vertex) map[*Vertex]int {
	output := make(map[*Vertex]int)
	for _, vertex := range graph {
		output[vertex] = 0
	}
	return output
}
