package graph

import "slices"

// DependencyOrder solves the problem in O(n) time and O(n) space.
func DependencyOrder(graph []*VertexWithIngress) ([]*VertexWithIngress, error) {
	sorted, err := TopologicalSort(graph)
	if err != nil {
		return nil, err
	}
	slices.Reverse(sorted)
	return sorted, nil
}
