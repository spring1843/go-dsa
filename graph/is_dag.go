package graph

type node struct {
	val int
}

// IsDAG returns true if the given graph is a Directed Acyclic Graph (DAG)
// A directed graph is acyclic if it contains no cycles
func IsDAG(graph []*Vertex) bool {
	for i, vertex := range graph {
		visited := make(map[*Vertex]struct{})
		visited[vertex] = struct{}{}
		if hasCycle(i, vertex, visited) {
			return false
		}
	}
	return true
}

func hasCycle(i int, vertex *Vertex, visited map[*Vertex]struct{}) bool {
	for _, neighbour := range vertex.Edges {
		if _, ok := visited[neighbour]; !ok {
			visited[neighbour] = struct{}{}
			if hasCycle(vertex.Val-1, neighbour, visited) {
				return true
			}
			delete(visited, neighbour)
		} else {
			return true
		}
	}
	return false
}
