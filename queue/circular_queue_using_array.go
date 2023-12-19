package queue

import "errors"

// CircularQueue is a queue that is made using a circular array.
type CircularQueue struct {
	data  []int
	size  int
	front int
	rear  int
}

const emptyValue = 0

// ErrQueueAtMaxCapacity occurs when the queue is at max capacity.
var ErrQueueAtMaxCapacity = errors.New("queue is at max capacity")

// NewCircularQueue returns a fixed size circular queue.
func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{
		data:  make([]int, size),
		rear:  -1,
		size:  0,
		front: 0,
	}
}

// enqueue solves the problem in O(1) time and O(1) space.
func (queue *CircularQueue) enqueue(n int) error {
	if queue.isFull() {
		return ErrQueueAtMaxCapacity
	}

	queue.rear++
	if queue.rear == queue.capacity() {
		queue.rear = 0
	}

	queue.data[queue.rear] = n
	queue.size++
	return nil
}

// dequeue solves the problem in O(1) time and O(1) space.
func (queue *CircularQueue) dequeue() (int, error) {
	tmp := queue.data[queue.front]
	queue.data[queue.front] = emptyValue
	queue.front++
	if queue.front == queue.capacity() {
		queue.front = 0
	}
	queue.size--
	return tmp, nil
}

func (queue *CircularQueue) isFull() bool {
	return queue.size == queue.capacity()
}

func (queue *CircularQueue) capacity() int {
	return len(queue.data)
}

func (queue *CircularQueue) isEmpty() bool {
	return queue.front == (queue.rear+1)%queue.capacity()
}
