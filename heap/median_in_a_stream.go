package heap

import "container/heap"

type (
	minHeap      []int
	maxHeap      []int
	medianKeeper struct {
		max *maxHeap
		min *minHeap
	}
)

func newMedianKeeper() medianKeeper {
	return medianKeeper{&maxHeap{}, &minHeap{}}
}

// addNumber solves the problem in O(Log n) time and O(n) space.
func (m *medianKeeper) addNumber(num int) {
	if m.len()%2 == 0 {
		if m.len() == 0 {
			heap.Push(m.max, num)
			return
		}
		if num > (*m.min)[0] {
			heap.Push(m.max, heap.Pop(m.min))
			heap.Push(m.min, num)
			return
		}
		heap.Push(m.max, num)
		return
	}
	if num < (*m.max)[0] {
		heap.Push(m.min, heap.Pop(m.max))
		heap.Push(m.max, num)
		return
	}
	heap.Push(m.min, num)
}

func (m *medianKeeper) len() int {
	return m.min.Len() + m.max.Len()
}

func (m *medianKeeper) median() float64 {
	if m.len()&1 == 1 {
		return float64((*m.max)[0])
	}
	return float64((*m.max)[0]+(*m.min)[0]) / 2.0
}

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *minHeap) Push(x any)        { *m = append(*m, x.(int)) }

func (m *minHeap) Pop() any {
	ret := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return ret
}

func (m maxHeap) Len() int           { return len(m) }
func (m maxHeap) Less(i, j int) bool { return m[i] > m[j] }
func (m maxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *maxHeap) Push(x any)        { *m = append(*m, x.(int)) }

func (m *maxHeap) Pop() any {
	ret := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return ret
}
