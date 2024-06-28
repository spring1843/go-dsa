package heap

import (
	"container/heap"
	"math"
)

type (
	point struct {
		coordinates []int
		distance    float64
	}
	pointsHeap []*point
)

// KClosestPointToOrigin solves the problem in O(n*Log k) time and O(k) space.
func KClosestPointToOrigin(points [][]int, k int) [][]int {
	if len(points) <= 1 {
		return points
	}
	h := new(pointsHeap)
	for _, point := range points {
		heap.Push(h, newPoint(point))
	}

	output := make([][]int, k)
	for i := range k {
		output[i] = heap.Pop(h).(*point).coordinates
	}
	return output
}

func newPoint(a []int) *point {
	return &point{
		coordinates: a,
		distance:    math.Sqrt(float64(a[0]*a[0] + a[1]*a[1])),
	}
}

func (p pointsHeap) Len() int           { return len(p) }
func (p pointsHeap) Less(i, j int) bool { return p[i].distance < p[j].distance }
func (p pointsHeap) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p *pointsHeap) Push(x any)        { *p = append(*p, x.(*point)) }

func (p *pointsHeap) Pop() any {
	old := *p
	tmp := old[len(old)-1]
	*p = old[0 : len(old)-1]
	return tmp
}
