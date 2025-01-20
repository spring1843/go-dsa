package heap

import "container/heap"

type slidingWindow []int

// MaxSlidingWindow solves the problem in O(n*Log k) time and O(k) space.
func MaxSlidingWindow(numbers []int, k int) []int {
	output := []int{}
	if len(numbers) <= 1 || len(numbers) < k {
		return output
	}

	pq := make(slidingWindow, k)
	i := 0
	for i < k {
		pq[i] = numbers[i]
		i++
	}

	heap.Init(&pq)
	output = append(output, pq[0])

	for i < len(numbers) {
		for j := range k {
			if pq[j] == numbers[i-k] {
				pq[j] = pq[len(pq)-1]
				pq = pq[:len(pq)-1]
				break
			}
		}
		heap.Push(&pq, numbers[i])
		output = append(output, pq[0])
		i++
	}

	return output
}

func (m slidingWindow) Len() int           { return len(m) }
func (m slidingWindow) Less(i, j int) bool { return m[i] > m[j] }
func (m slidingWindow) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *slidingWindow) Push(x any)        { *m = append(*m, x.(int)) }

func (m *slidingWindow) Pop() any {
	old := *m
	tmp := old[len(old)-1]
	*m = old[0 : len(old)-1]
	return tmp
}
