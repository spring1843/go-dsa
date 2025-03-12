package queue

import "math"

const emptyValue = math.MinInt64

type CircularQueue struct {
	data []int
	size int
	head int
	tail int
}

func NewCircularQueue(size int) *CircularQueue {
	data := make([]int, size)
	for i := range data {
		data[i] = emptyValue
	}
	return &CircularQueue{
		data: data,
		size: size,
		head: -1,
		tail: 0,
	}
}

func (queue *CircularQueue) enqueue(n int) {
	queue.head = queue.nextCircularSlicePosition(queue.head)

	if queue.data[queue.head] != emptyValue {
		queue.tail = queue.nextCircularSlicePosition(queue.tail)
	}

	queue.data[queue.head] = n
}

func (queue *CircularQueue) dequeue() (int, error) {
	if queue.data[queue.tail] == emptyValue {
		return 0, ErrQueueEmpty
	}

	output := queue.data[queue.tail]
	queue.data[queue.tail] = emptyValue
	queue.tail = queue.nextCircularSlicePosition(queue.tail)
	return output, nil
}

func (queue *CircularQueue) nextCircularSlicePosition(value int) int {
	value++
	if value == queue.size {
		return 0
	}
	return value
}
