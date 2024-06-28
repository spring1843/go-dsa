package heap

import (
	"container/heap"

	"github.com/spring1843/go-dsa/linkedlist"
)

type (
	priorityQueue []*linkedlist.Node
)

// MergeSortedLists solves the problem in O(n*Log k) time and O(k) space.
func MergeSortedLists(lists []*linkedlist.Node) *linkedlist.Node {
	pq := new(priorityQueue)

	for _, item := range lists {
		if item == nil {
			continue
		}
		heap.Push(pq, item)
	}

	var head, cur *linkedlist.Node
	for pq.Len() > 0 {
		node := heap.Pop(pq).(*linkedlist.Node)
		if cur == nil {
			cur = node
			head = cur
			continue
		}
		cur.Next = node
		cur = cur.Next

		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
	}

	return head
}

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func (pq *priorityQueue) Push(x any) {
	*pq = append(*pq, x.(*linkedlist.Node))
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
