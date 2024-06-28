package graph

// DependencyOrder solves the problem in O(n) time and O(n) space.
func DependencyOrder(graph []*VertexWithIngress) ([]*VertexWithIngress, error) {
	sorted, err := TopologicalSort(graph)
	if err != nil {
		return nil, err
	}
	return reverse(sorted), nil
}

func reverse(graph []*VertexWithIngress) []*VertexWithIngress {
	for i := range len(graph) / 2 {
		graph[i], graph[len(graph)-1-i] = graph[len(graph)-1-i], graph[i]
	}
	return graph
}
