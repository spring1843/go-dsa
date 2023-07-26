package heap

import "container/heap"

// RegularNumbers solves the problem in O(n*log n) time and O(n) space.
func RegularNumbers(n int) []int {
	output := []int{}
	minHeap := &minHeap{}
	heap.Push(minHeap, 1)

	last, count := 0, 0

	for count < n {
		current := heap.Pop(minHeap).(int)
		if current > last {
			output = append(output, current)
			last = current
			count++
			heap.Push(minHeap, current*2)
			heap.Push(minHeap, current*3)
			heap.Push(minHeap, current*5)
		}
	}

	return output
}
