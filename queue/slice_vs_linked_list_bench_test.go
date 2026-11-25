package queue

import (
	"container/list"
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
		if _, err := q.dequeue(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSliceQueue(b *testing.B) {
	q := sliceQueue{}
	for n := range b.N {
		q.enqueue(n)
	}
	for range b.N {
		if _, err := q.dequeue(); err != nil {
			b.Fatal(err)
		}
	}
}

func (q *sliceQueue) enqueue(val int) {
	*q = append(*q, val)
}

func (q *sliceQueue) dequeue() (int, error) {
	if len(*q) == 0 {
		return 0, ErrQueueEmpty
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
		return 0, ErrQueueEmpty
	}
	return q.items.Remove(q.items.Front()).(int), nil
}
