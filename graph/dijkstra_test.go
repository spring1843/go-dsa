package graph

import (
	"testing"
)

type dijkstraTestCase struct {
	graph     []*dijkstraVertex
	distances []int
}

/*
TestDijkstra tests solution(s) with the following signature and problem description:

	dijkstraVertex struct {
		val        int
		distance   int
		discovered bool
		edges      []*weightedEdge
	}
	weightedEdge struct {
		edge   *dijkstraVertex
		weight int
	}
	func Dijkstra(graph []*dijkstraVertex, source *dijkstraVertex)

Implement Dijkstra's algorithm. Given a single sourced edge weighted graph, find the
shortest path from the source of the graph to every node and set it to the distance
field of the node.

			     5
		     (4)  ↖ (1)
			2 ─────→ 4
	    (1) ↑ (2) ↗ (1)
		    1 ──→ 3

For the above graph, the shortest paths for each node are {0, 1, 3, 5, 8}.
*/
func TestDijkstra(t *testing.T) {
	tests := []*dijkstraTestCase{
		{
			graph:     newGraph(2),
			distances: []int{0, 10},
		},
		{
			graph:     newGraph(5),
			distances: []int{0, 1, 3, 5, 8},
		},
		{
			graph:     newGraph(4),
			distances: []int{0, 5, 3, 10},
		},
		{
			graph:     newGraph(3),
			distances: []int{0, 3, 5},
		},
		{
			graph:     newGraph(6),
			distances: []int{0, 1, 3, 4, 6, 10},
		},
	}
	addEdges(tests)

	for i, test := range tests {
		Dijkstra(test.graph, test.graph[0])
		for j, v := range test.graph {
			if v.distance != test.distances[j] {
				t.Errorf("Test Case %d, Vertex %d distance: got %d, expected %d", i, v.val, v.distance, test.distances[j])
			}
		}
	}
}

func newGraph(edges int) []*dijkstraVertex {
	graph := make([]*dijkstraVertex, edges)
	for i := 0; i < edges; i++ {
		graph[i] = &dijkstraVertex{val: i}
	}
	return graph
}

func addEdges(tests []*dijkstraTestCase) {
	tests[0].graph[0].edges = []*weightedEdge{{edge: tests[0].graph[1], weight: 10}}

	tests[1].graph[0].edges = []*weightedEdge{{edge: tests[1].graph[1], weight: 1}, {edge: tests[1].graph[2], weight: 3}}
	tests[1].graph[1].edges = []*weightedEdge{{edge: tests[1].graph[2], weight: 2}, {edge: tests[1].graph[3], weight: 4}}
	tests[1].graph[2].edges = []*weightedEdge{{edge: tests[1].graph[3], weight: 3}, {edge: tests[1].graph[4], weight: 5}}
	tests[1].graph[3].edges = []*weightedEdge{{edge: tests[1].graph[4], weight: 5}}

	tests[2].graph[0].edges = []*weightedEdge{{edge: tests[2].graph[1], weight: 5}, {edge: tests[2].graph[2], weight: 3}}
	tests[2].graph[1].edges = []*weightedEdge{{edge: tests[2].graph[2], weight: 2}, {edge: tests[2].graph[3], weight: 6}}
	tests[2].graph[2].edges = []*weightedEdge{{edge: tests[2].graph[3], weight: 7}}

	tests[3].graph[0].edges = []*weightedEdge{{edge: tests[3].graph[1], weight: 3}, {edge: tests[3].graph[2], weight: 5}}
	tests[3].graph[1].edges = []*weightedEdge{{edge: tests[3].graph[2], weight: 2}}

	tests[4].graph[0].edges = []*weightedEdge{{edge: tests[4].graph[1], weight: 1}, {edge: tests[4].graph[2], weight: 4}}
	tests[4].graph[1].edges = []*weightedEdge{{edge: tests[4].graph[2], weight: 2}, {edge: tests[4].graph[3], weight: 5}}
	tests[4].graph[2].edges = []*weightedEdge{{edge: tests[4].graph[3], weight: 1}, {edge: tests[4].graph[4], weight: 3}}
	tests[4].graph[3].edges = []*weightedEdge{{edge: tests[4].graph[4], weight: 2}, {edge: tests[4].graph[1], weight: 1}}
	tests[4].graph[4].edges = []*weightedEdge{{edge: tests[4].graph[5], weight: 4}}
}
