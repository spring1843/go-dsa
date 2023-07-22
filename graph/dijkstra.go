package graph

import (
	"container/heap"
	"math"
)

type (
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
	verticesHeap []*dijkstraVertex
)

func Dijkstra(graph []*dijkstraVertex, source *dijkstraVertex) {
	for _, vertex := range graph {
		vertex.distance = math.MaxInt
		vertex.discovered = false
	}

	source.distance = 0

	minHeap := new(verticesHeap)
	heap.Push(minHeap, source)

	for minHeap.Len() != 0 {
		u := heap.Pop(minHeap).(*dijkstraVertex)
		u.discovered = true

		for _, edge := range u.edges {
			v := edge.edge
			if v.discovered {
				continue
			}

			if u.distance+edge.weight < v.distance {
				v.distance = u.distance + edge.weight
				heap.Push(minHeap, v)
			}
		}
	}
}

func (p verticesHeap) Len() int            { return len(p) }
func (p verticesHeap) Less(i, j int) bool  { return p[i].distance < p[j].distance }
func (p verticesHeap) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *verticesHeap) Push(x interface{}) { *p = append(*p, x.(*dijkstraVertex)) }
func (p *verticesHeap) Pop() interface{} {
	old := *p
	tmp := old[len(old)-1]
	*p = old[0 : len(old)-1]
	return tmp
}
