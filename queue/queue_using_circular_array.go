package queue

import "errors"

// UsingCircularArray is a queue that is made using a circular array.
type UsingCircularArray struct {
	circular []int
	size     int
	front    int
	rear     int
}

const emptyValue = 0

var ErrQueueAtMaxCapacity = errors.New("queue is at max capacity")

// NewCircularQueue returns a fixed size circular queue.
func NewCircularQueue(size int) *UsingCircularArray {
	circular := make([]int, size)
	for i := range circular {
		circular[i] = emptyValue
	}

	return &UsingCircularArray{
		circular: circular,
		rear:     -1,
		size:     0,
		front:    0,
	}
}

func (queue *UsingCircularArray) enqueue(n int) error {
	if queue.isFull() {
		return ErrQueueAtMaxCapacity
	}

	queue.rear++
	if queue.rear == queue.capacity() {
		queue.rear = 0
	}

	queue.circular[queue.rear] = n
	queue.size++
	return nil
}

func (queue *UsingCircularArray) dequeue() (int, error) {
	if queue.isEmpty() {
		return -1, ErrQueueEmpty
	}
	tmp := queue.circular[queue.front]
	queue.circular[queue.front] = emptyValue
	queue.front++
	if queue.front == queue.capacity() {
		queue.front = 0
	}
	queue.size--
	return tmp, nil
}

func (queue *UsingCircularArray) isFull() bool {
	return queue.size == queue.capacity()
}

func (queue *UsingCircularArray) isEmpty() bool {
	return queue.size == 0
}

func (queue *UsingCircularArray) capacity() int {
	return len(queue.circular)
}
