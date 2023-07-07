package graph

// IsDAG solves the problem in O(V+E) time and O(V) space.
func IsDAG(graph []*Vertex) bool {
	for _, vertex := range graph {
		visited := make(map[*Vertex]struct{})
		visited[vertex] = struct{}{}
		if hasCycle(vertex, visited) {
			return false
		}
	}
	return true
}

func hasCycle(vertex *Vertex, visited map[*Vertex]struct{}) bool {
	for _, neighbor := range vertex.Edges {
		if _, ok := visited[neighbor]; !ok {
			visited[neighbor] = struct{}{}
			if hasCycle(neighbor, visited) {
				return true
			}
			delete(visited, neighbor)
		} else {
			return true
		}
	}
	return false
}
