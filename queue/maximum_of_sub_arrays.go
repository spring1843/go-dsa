package queue

import (
	"container/list"
	"errors"
)

// doubleEndedQueue is a queue that allows dequeue from both sides
type doubleEndedQueue struct {
	queue []int
}

// ErrQueueEmpty occurs when popping an empty stack
var ErrQueueEmpty = errors.New("Empty queue")

// MaxOfKLengthSubArrays takes a list of integers and returns
// returns a list of the maximum of each subarray of length k.
func MaxOfKLengthSubArrays(numbers []int, k int) ([]int, error) {
	output := []int{}
	queue := list.New()
	var i int
	for i = 0; i < k; i++ {
		for queue.Len() != 0 && numbers[i] >= numbers[queue.Back().Value.(int)] {
			queue.Remove(queue.Back())
		}

		queue.PushBack(i)
	}

	for i < len(numbers) {
		output = append(output, numbers[queue.Front().Value.(int)])
		for queue.Len() != 0 && queue.Front().Value.(int) <= i-k {
			queue.Remove(queue.Front())
		}

		for queue.Len() != 0 && numbers[i] >= numbers[queue.Back().Value.(int)] {
			queue.Remove(queue.Back())
		}

		queue.PushBack(i)
		i++
	}

	output = append(output, numbers[queue.Front().Value.(int)])
	return output, nil
}
