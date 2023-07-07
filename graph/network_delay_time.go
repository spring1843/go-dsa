package graph

import (
	"container/heap"
	"math"
)

type edgeMinHeap [][]int

const (
	vertexSource, vertexDestination, vertexDelay = 0, 1, 2
	edgeDestination, edgeCost                    = 0, 1
)

// NetworkDelayTime solves the problem in O(n log n) time and O(n) space.
func NetworkDelayTime(n, k int, edges [][3]int) int {
	var (
		verticesMap, edgesHeap = verticesAndEdges(edges, k)
		max                    = math.MinInt64
		cost                   = make(map[int]int)
	)
	cost[k] = 0
	for len(edgesHeap) > 0 {
		edge := heap.Pop(&edgesHeap).([]int)
		if _, ok := cost[edge[edgeDestination]]; ok {
			continue
		}
		cost[edge[edgeDestination]] = edge[edgeCost]
		if max < edge[edgeCost] {
			max = edge[edgeCost]
		}
		if len(cost) == n {
			break
		}
		for _, vertex := range verticesMap[edge[edgeDestination]] {
			i := 0
			for i < len(edgesHeap) && edgesHeap[i][0] != vertex[vertexSource] {
				i++
			}
			if i == len(edgesHeap) {
				heap.Push(&edgesHeap, []int{vertex[vertexSource], vertex[vertexDestination] + cost[edge[edgeDestination]]})
			}
			if edgesHeap[i][edgeCost] < vertex[vertexDestination]+cost[edge[edgeDestination]] {
				edgesHeap[i][edgeCost] = vertex[vertexDestination] + cost[edge[edgeDestination]]
				heap.Fix(&edgesHeap, i)
			}
		}
	}
	if len(cost) < n {
		return -1
	}
	return max
}

// verticesAndEdges takes edges, and returns
// a map of vertex ID as keys and the outgoing edges in the [vertexDestination, cost] form
// a minimum heap of edges in the [vertexDestination, cost].
func verticesAndEdges(edges [][3]int, k int) (map[int][][2]int, edgeMinHeap) {
	verticesMap := make(map[int][][2]int)
	for _, vertex := range edges {
		verticesMap[vertex[vertexSource]] = append(verticesMap[vertex[vertexSource]], [2]int{vertex[vertexDestination], vertex[vertexDelay]})
	}
	edgeHeap := make(edgeMinHeap, 0, len(edges))
	for _, edge := range verticesMap[k] {
		edgeHeap = append(edgeHeap, []int{edge[edgeDestination], edge[edgeCost]})
	}
	heap.Init(&edgeHeap)
	return verticesMap, edgeHeap
}

func (e edgeMinHeap) Len() int            { return len(e) }
func (e edgeMinHeap) Less(i, j int) bool  { return e[i][1] <= e[j][1] }
func (e edgeMinHeap) Swap(i, j int)       { e[i], e[j] = e[j], e[i] }
func (e *edgeMinHeap) Push(x interface{}) { *e = append(*e, x.([]int)) }
func (e *edgeMinHeap) Pop() interface{} {
	x := (*e)[len(*e)-1]
	*e = (*e)[0 : len(*e)-1]
	return x
}
