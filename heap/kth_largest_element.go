package heap

import (
	"container/heap"
)

type minimumHeap []int

// KthLargestElement finds the kth largest element in an list.
func KthLargestElement(elements []int, k int) int {
	minHeap := minimumHeap(elements)
	heap.Init(&minHeap)
	for minHeap.Len() > k {
		heap.Pop(&minHeap)
	}
	return heap.Pop(&minHeap).(int)
}

func (m minimumHeap) Len() int            { return len(m) }
func (m minimumHeap) Less(i, j int) bool  { return m[i] < m[j] }
func (m minimumHeap) Swap(i, j int)       { m[i], m[j] = m[j], m[i] }
func (m *minimumHeap) Push(x interface{}) { *m = append(*m, x.(int)) }

func (m *minimumHeap) Pop() interface{} {
	old := *m
	tmp := old[len(old)-1]
	*m = old[0 : len(old)-1]
	return tmp
}
