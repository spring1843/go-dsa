package queue

import (
	"container/list"
	"errors"
	"testing"
)

type (
	linkedListQueue struct {
		items *list.List
	}

	sliceQueue []int
)

func BenchmarkLinkedListQueue(b *testing.B) {
	q := newLinkedListQueue()
	for n := range b.N {
		q.enqueue(n)
	}
	for range b.N {
		q.dequeue()
	}
}

func BenchmarkSliceQueue(b *testing.B) {
	q := sliceQueue{}
	for n := range b.N {
		q.enqueue(n)
	}
	for range b.N {
		q.dequeue()
	}
}

func (q *sliceQueue) enqueue(val int) {
	*q = append(*q, val)
}

func (q *sliceQueue) dequeue() (int, error) {
	if len(*q) == 0 {
		return 0, errors.New("queue is empty")
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value, nil
}

func newLinkedListQueue() *linkedListQueue {
	return &linkedListQueue{items: list.New()}
}

func (q *linkedListQueue) enqueue(val int) {
	q.items.PushBack(val)
}

func (q *linkedListQueue) dequeue() (int, error) {
	if q.items.Len() == 0 {
		return 0, errors.New("queue is empty")
	}
	return q.items.Remove(q.items.Front()).(int), nil
}
